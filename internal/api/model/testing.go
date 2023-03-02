package model

import "testing"

func TestUser(t *testing.T) *User{
  return &User{
	Email: "mynew@email.com",
	Password: "123456",

  }
}