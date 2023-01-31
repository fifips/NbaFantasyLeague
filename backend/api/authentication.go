package api

import (
	"backend/common"
	db "backend/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"net/http"
	"os"
	"strconv"
	"time"
)

var secretKey string

func init() {
	secretKey = os.Getenv("jwt_secret")
}

func createAccessToken(user db.User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(user.Id),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(common.AccessTokenExpiration)),
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func createRefreshToken(user db.User) (string, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	session := db.Session{
		Id:     tokenId,
		UserId: user.Id,
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        session.Id.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(common.RefreshTokenExpiration)),
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	if err := db.CreateOrUpdateSession(session); err != nil {
		return "", err
	}

	return token, nil
}

// createAndSetTokenPair creates user's access and refresh tokens and sets them in the response as cookies
func createAndSetTokenPair(c *gin.Context, user db.User) error {
	accessToken, err := createAccessToken(user)
	if err != nil {
		return err
	}

	refreshToken, err := createRefreshToken(user)
	if err != nil {
		return err
	}

	c.SetCookie("access_token", accessToken, int(common.AccessTokenExpiration.Seconds()), "/", "localhost", true, true)
	c.SetCookie("refresh_token", refreshToken, int(common.RefreshTokenExpiration.Seconds()), "/", "localhost", true, true)
	return nil
}

func removeTokenPair(c *gin.Context, user db.User) error {
	if err := db.DeleteSessionByUserId(user.Id); err != nil {
		return err
	}
	c.SetCookie("access_token", "", int(-time.Hour.Seconds()), "/", "", true, true)
	c.SetCookie("refresh_token", "", int(-time.Hour.Seconds()), "/", "", true, true)
	return nil
}

func getUserFromAccessToken(token string) (*db.User, error) {
	var claims jwt.RegisteredClaims
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	issuerId, err := strconv.Atoi(claims.Issuer)
	if err != nil {
		return nil, err
	}

	user, err := db.GetUserById(issuerId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func getUserFromRefreshToken(token string) (*db.User, error) {
	var claims jwt.RegisteredClaims
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	sessionId, err := uuid.Parse(claims.ID)
	if err != nil {
		return nil, err
	}

	session, err := db.GetSessionById(sessionId)
	if err != nil {
		return nil, err
	}

	user, err := db.GetUserById(session.UserId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func Authenticate(c *gin.Context) {
	accessToken, err := c.Cookie("access_token")
	if err != nil && err == (http.ErrNoCookie) {
		refreshToken, err := c.Cookie("refresh_token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message1": err.Error()})
			return
		}

		user, err := getUserFromRefreshToken(refreshToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message2": err.Error()})
			return
		}

		if err := createAndSetTokenPair(c, *user); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message3": err.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusOK, user)
		return
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message4": err.Error()})
		return
	}

	user, err := getUserFromAccessToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message5": err.Error()})
		return
	}

	c.Set("user", *user)
	c.Next()
}
