package store_test

import (
	"Go-http-rest-api/internal/api/model"
	"Go-http-rest-api/internal/api/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")
	u, err := s.User().Create(&model.User{
		Email: "mynew@email.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "mynew@email.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)
	s.User().Create(&model.User{
		Email: "mynew@email.com",
	})
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
