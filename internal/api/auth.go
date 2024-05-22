package api

import (
	"context"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/tmpmadula/cantina-shop/config"
	"github.com/tmpmadula/cantina-shop/internal/db"
	"github.com/tmpmadula/cantina-shop/internal/models"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/oauth2/v2"
)

var oauthStateString = "random"

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares a password with a hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// RegisterHandler handles user registration
func RegisterHandler(c *gin.Context) {
	var request struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	_, err := db.GetUserByEmail(request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// Hash password
	hashedPassword, err := HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user
	user := &models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
		Verified: false,
	}
	if err := db.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Send verification email (dummy implementation)
	// In production, use an email service to send the verification link
	verificationToken := "dummy-verification-token"
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "verification_token": verificationToken})
}

// VerifyEmailHandler handles email verification
func VerifyEmailHandler(c *gin.Context) {
	// Dummy implementation: In production, check the verification token
	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}

// LoginHandler handles email/password login
func LoginHandler(c *gin.Context) {
	var request struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user from the database
	user, err := db.GetUserByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check password
	if !CheckPasswordHash(request.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check if user is verified
	if !user.Verified {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not verified"})
		return
	}

	// Create a session for the user
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}

// GoogleLoginHandler handles Google OAuth2 login
func GoogleLoginHandler(c *gin.Context) {
	url := config.GoogleOauthConfig.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallbackHandler handles the callback from Google OAuth2
func GoogleCallbackHandler(c *gin.Context) {
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
			Name:     userinfo.Name,
			Email:    userinfo.Email,
			Password: "", // Password is not needed for OAuth2 users
			Verified: true,
		}
		if err := db.CreateUser(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
	}

	// Create a session for the user
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	// Save the login method
	db.SaveLogin(user.ID, "google")

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}
