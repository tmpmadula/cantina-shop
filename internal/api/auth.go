package api

import (
	"context"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/tmpmadula/cantina-shop/config"
	"github.com/tmpmadula/cantina-shop/internal/db"
	"github.com/tmpmadula/cantina-shop/internal/models"
	"golang.org/x/oauth2"
	"google.golang.org/api/oauth2/v2"
)

var oauthStateString = "random"

func LoginHandler(c *gin.Context) {
	url := config.GoogleOauthConfig.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func CallbackHandler(c *gin.Context) {
	state := c.Query("state")
	if state != oauthStateString {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state parameter"})
		return
	}

	code := c.Query("code")
	token, err := config.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Code exchange failed"})
		return
	}

	client := config.GoogleOauthConfig.Client(context.Background(), token)
	service, err := oauth2.New(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create oauth2 client"})
		return
	}

	userinfo, err := service.Userinfo.Get().Do()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get userinfo"})
		return
	}

	// Check if user exists in the database
	user, err := db.GetUserByEmail(userinfo.Email)
	if err != nil {
		// If user doesn't exist, create a new one
		user = &models.User{
			Name:  userinfo.Name,
			Email: userinfo.Email,
		}
		err = db.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
	}

	// Create a session for the user
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}
