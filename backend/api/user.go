package api

import (
	db "NbaFantasyLeague/database"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

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

	if err := db.CreateUser(newUser); err != nil {
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
	u, exists := c.Get("user")
	if exists != true {
		c.Status(http.StatusUnauthorized)
		return
	}
	user, ok := u.(db.User)
	if ok != true {
		c.Status(http.StatusInternalServerError)
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
	if ok != true {
		c.Status(http.StatusUnauthorized)
		return
	}
	c.Status(http.StatusOK)
}
