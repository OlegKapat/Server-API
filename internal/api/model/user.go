package model

type User struct {
	ID int64 `json:"id"`
	Email string `json:"email"`
	EncryptedPassword string `json:"encrypted_password"`
}