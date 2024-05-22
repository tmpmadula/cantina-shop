package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	ClientID:     "822314717402-o2g7eeau9h5uc6patrcofi2o1n58j5tn.apps.googleusercontent.com",
	ClientSecret: "niRBJRSWSpDeN48aN0GGz3wG",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}
