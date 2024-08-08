package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

const (
	clientID     = "131846"
	clientSecret = "397dde79546b1d38fce1a0a9ecff87ad8919f5bf"
	redirectURI  = "http://localhost:8080/callback"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		authUrl := fmt.Sprintf("https://www.strava.com/oauth/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=read,activity:read_all&approval_prompt=auto",
			clientID, redirectURI)

		c.Redirect(http.StatusFound, authUrl)
	})

	router.GET("/callback", func(c *gin.Context) {
		code := c.Query("code")
		if code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no cede in request"})
		}

		token, err := exchangeToken(code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to exchange token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"access_token": token})
	})

	router.Run(":8080")
}

func exchangeToken(code string) (string, error) {
	tokenURL := "https://www.strava.com/oauth/token"

	// Menggunakan url.Values untuk membentuk data POST
	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")

	res, err := http.PostForm(tokenURL, data)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return "", nil
	}

	token, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("access_token not found in respone")
	}
	return token, nil
}
