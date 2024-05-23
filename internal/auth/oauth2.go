// internal/auth/oauth2.go
package auth

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/tmpmadula/cantina-shop/internal/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8000/callback",
	ClientID:     "822314717402-o2g7eeau9h5uc6patrcofi2o1n58j5tn.apps.googleusercontent.com",
	ClientSecret: "niRBJRSWSpDeN48aN0GGz3wG",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

// @Summary Login with Google
// @Description Login with Google
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 302 {string} string "Redirects to Google login page"
// @Router /google/login [get]
func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL("randomState")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// @Summary Google callback
// @Description Google callback
// @Tags auth
// @Accept  json
// @Produce  json
// @Param state query string true "State parameter"
// @Param code query string true "Code parameter"
// @Success 302 {string} string "Redirects to home page"
// @Router /callback [get]
func HandleGoogleCallback(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := r.FormValue("state")
		if state != "randomState" {
			http.Error(w, "State parameter doesn't match", http.StatusBadRequest)
			return
		}

		code := r.FormValue("code")
		token, err := googleOauthConfig.Exchange(context.Background(), code)
		if err != nil {
			http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
		if err != nil {
			http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()

		var googleUser struct {
			Id    string `json:"id"`
			Email string `json:"email"`
			Name  string `json:"name"`
		}
		if err := json.NewDecoder(response.Body).Decode(&googleUser); err != nil {
			http.Error(w, "Failed to parse user info: "+err.Error(), http.StatusInternalServerError)
			return
		}

		var user models.User
		err = db.QueryRow("SELECT id, name, email, role FROM users WHERE email = $1", googleUser.Email).
			Scan(&user.ID, &user.Name, &user.Email, &user.Role)
		if err == sql.ErrNoRows {
			err = db.QueryRow("INSERT INTO users (name, email, role) VALUES ($1, $2, $3) RETURNING id",
				googleUser.Name, googleUser.Email, "user").Scan(&user.ID)
			if err != nil {
				http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
				return
			}
		} else if err != nil {
			http.Error(w, "Failed to get user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		//	tokenString, err := auth.GenerateJWT(user.Email)
		if err != nil {
			http.Error(w, "Failed to generate token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/?token=", http.StatusTemporaryRedirect)
	}
}
