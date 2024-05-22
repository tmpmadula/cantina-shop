package models

type User struct {
	ID       int64  `json:"id" pg:"id,pk"`
	Name     string `json:"name" pg:"name"`
	Email    string `json:"email" pg:"email,unique"`
	Password string `json:"-" pg:"password"`
	Verified bool   `json:"verified" pg:"verified"`
}

type Login struct {
	ID        int64  `json:"id" pg:"id,pk"`
	UserID    int64  `json:"user_id" pg:"user_id"`
	Method    string `json:"method" pg:"method"`
	Timestamp string `json:"timestamp" pg:"timestamp"`
}
