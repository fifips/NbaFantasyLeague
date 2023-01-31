package api

import (
	. "backend/common"
	db "backend/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/smtp"
	"os"
	"time"
)

// getUserFromContext returns User from gin context and
// throws error if the User is not found or not of correct type
func getUserFromContext(c *gin.Context) (db.User, error) {
	var user db.User

	u, exists := c.Get("user")
	if !exists {
		return user, fmt.Errorf("user not found in context")
	}
	user, ok := u.(db.User)
	if ok == false {
		return user, fmt.Errorf("user in context is not of correct type")
	}

	return user, nil
}

func register(c *gin.Context) {
	var newUser db.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	password, err := bcrypt.GenerateFromPassword(newUser.Password, 12)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	newUser.Password = password

	if err := db.CreateUser(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := createTokenPair(c, newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	code, err := uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	aC := db.ActivationCode{
		Code:    code,
		UserId:  newUser.Id,
		Expires: time.Now().Add(ActivationCodeExpiration),
	}

	if err := db.CreateOrUpdateActivationCode(aC); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := sendActivationEmail(newUser.Email, aC.Code.String()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func login(c *gin.Context) {
	var newUser db.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	user, err := db.GetUserByEmail(newUser.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, newUser.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	if err := createTokenPair(c, *user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func logout(c *gin.Context) {
	user, err := getUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message1": err.Error()})
		return
	}
	if err := removeTokenPair(c, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func isLoggedIn(c *gin.Context) {
	_, ok := c.Get("user")
	if ok == false {
		c.Status(http.StatusUnauthorized)
		return
	}
	c.Status(http.StatusOK)
}

// sendActivationEmail generates and sends an activation email to given recipient
// with activation code(string) consisting of 4 random characters needed for activating user account
// The code is valid for 15 minutes.
func sendActivationEmail(recipient string, activationCode string) error {
	activationLink := fmt.Sprintf(ActivationUrl, activationCode)
	subject := "Verify your account!"
	msg := []byte(
		fmt.Sprintf("To:%s\r\n", recipient) +
			fmt.Sprintf("Subject:%s\r\n", subject) +
			"Content-Type:text/html\r\n" +
			`
		<!DOCTYPE html>
		<html>
			<head>
			<title>Verify your account</title>
			</head>
			<body>
			<h1>Verify your account using this link:</h1>
				<div style="width: 100vw; display: grid; place-items: center;">
				  <a style="text-align: center; background-color:lightgray;
							font-weight: bold; font-size: 25px; width:250px; margin:0;"
                     target="_blank" href="` + activationLink + `" >
                    Activate my account
				  </a>
                  <p style="margin:8px;font-weight: bold;font-size: 13px;">
                    (it's valid for 15 minutes)
                  </p>
				</div>
			</body>
		</html>`)

	auth := smtp.PlainAuth("", "domieniapowia@gmail.com", os.Getenv("EMAIL_PASS"), "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, "domieniapowia@gmail.com", []string{recipient}, msg)
	if err != nil {
		return err
	}
	return nil
}

// IsActiveUser is a gin middleware for checking if user's account is active. If user is not active,
// aborts the request with status unauthorized (401)
func IsActiveUser(c *gin.Context) {
	user, err := getUserFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	if !user.IsActive {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Next()
}

// activateUser activates User's account if given activation code matches the one in the database
func activateUser(c *gin.Context) {
	user, err := getUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	if user.IsActive {
		c.JSON(http.StatusConflict, gin.H{"message": CustomError{Message: "Account is already active."}.Error()})
		return
	}

	activationCodeFromRequest := c.Query("activation_code")

	activationCodeFromDb, err := db.GetActivationCodeByUserId(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if activationCodeFromRequest != activationCodeFromDb.Code.String() ||
		activationCodeFromDb.Expires.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user.IsActive = true
	if err := db.ActivateUserById(user.Id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := db.DeleteActivationCodeByUserId(user.Id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": nil})
}
