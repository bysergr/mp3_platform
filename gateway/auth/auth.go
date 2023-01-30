package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	authAddress = os.Getenv("AUTH_SVC_ADDRESS")
	client      = &http.Client{}
)

func Login(c *gin.Context) (string, int, error) {
	username, password, ok := c.Request.BasicAuth()
	if !ok {
	}

	address := fmt.Sprintf("http://%s/login", authAddress)
	req, err := http.NewRequest("POST", address, nil)
	if err != nil {
		return "", 0, err
	}

	req.SetBasicAuth(username, password)

	res, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", 0, err
	}

	bodyString := string(body)

	return bodyString, res.StatusCode, nil
}

func Validate(c *gin.Context) (string, int, error) {
	cookie, err := c.Request.Cookie("jwt")
	if err != nil {
		return "", 0, err
	}

	err = cookie.Valid()
	if err != nil {
		return "", 0, err
	}

	if cookie.Value == "" {
		return "missing credential", 401, nil
	}

	address := fmt.Sprintf("http://%s/validate", authAddress)
	req, err := http.NewRequest("POST", address, nil)
	if err != nil {
		return "", 0, err
	}

	req.Header.Add("Authorization", cookie.Value)

	res, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", 0, err
	}

	bodyString := string(body)

	return bodyString, res.StatusCode, nil
}
