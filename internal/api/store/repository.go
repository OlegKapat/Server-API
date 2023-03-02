package store

import "Go-http-rest-api/internal/api/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}