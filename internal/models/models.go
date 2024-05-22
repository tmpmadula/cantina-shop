package models

type User struct {
	ID    int64  `json:"id" pg:"id,pk"`
	Name  string `json:"name" pg:"name"`
	Email string `json:"email" pg:"email,unique"`
}
