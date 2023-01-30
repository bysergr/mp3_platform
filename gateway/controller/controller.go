package controller

import (
	"net/http"

	"github.com/bysergr/mp3_platform/gateway/auth"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	body, code, err := auth.Login(c)
	if err != nil {
		c.String(http.StatusConflict, err.Error())
		return
	}

	if code != http.StatusOK {
		c.String(code, body)
	}

	maxAge := 86400
	c.SetCookie("jwt", body, maxAge, "/", "", true, true)
	c.String(http.StatusOK, "Login Sucessfully")
}

func Upload(c *gin.Context) {
	body, code, err := auth.Validate(c)
	if err != nil {
		c.String(http.StatusConflict, err.Error())
		return
	}

	if code != http.StatusOK {
		c.String(http.StatusUnauthorized, body)
		return
	}

	c.F
}

func Download(c *gin.Context) {
}

func Register() {
}
