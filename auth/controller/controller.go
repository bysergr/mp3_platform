package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bysergr/mp3_platform/auth/database"
	u "github.com/bysergr/mp3_platform/auth/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	db  = database.DB
	JWT = os.Getenv("JWT_SECRET")
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		c.String(http.StatusUnauthorized, "Miss Credentials")
		return
	}

	query := fmt.Sprintf("SELECT password FROM user WHERE email=%s", username)
	result, err := db.Query(query)
	if err != nil {
		c.String(http.StatusConflict, err.Error())
		return
	}

	defer result.Close()

	var passwordDB string
	for result.Next() {
		err = result.Scan(&passwordDB)

		if err != nil {
			c.String(http.StatusConflict, err.Error())
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(password)); err != nil {
		c.String(http.StatusUnauthorized, "Invalid Credentials")
		return

	} else {
		token, err := u.CreateJWT(username, JWT)
		if err != nil {
			c.String(http.StatusForbidden, err.Error())
			return
		}

		c.String(http.StatusOK, token)
	}
}

func Register(c *gin.Context) {
	var user User

	if err := c.BindJSON(&user); err != nil {
		c.String(http.StatusNoContent, "Miss Values")
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 15)
	if err != nil {
		c.String(http.StatusConflict, err.Error())
	}

	user.Password = string(password)

	query := fmt.Sprintf("INSERT INTO user (email, password) VALUES ('%s', '%s');", user.Email, user.Password)
	result, err := db.Query(query)
	if err != nil {
		c.String(http.StatusConflict, err.Error())
		return
	}

	defer result.Close()
	if result.Err() != nil {
		c.String(http.StatusConflict, err.Error())
		return
	}

	c.JSON(http.StatusCreated, user)
}

func Validate(c *gin.Context) {
	encoded_jwt := c.Request.Header.Get("Authorization")
	if encoded_jwt == "" {
		c.String(http.StatusUnauthorized, "Missig Credentials")
		return
	}

	decoded, err := jwt.ParseWithClaims(encoded_jwt, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT), nil
	})
	if err != nil {
		c.String(http.StatusUnauthorized, "Not Authorized")
		return
	}

	c.String(http.StatusOK, decoded.Raw)
}
