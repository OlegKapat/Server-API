package teststore_test

import (
	"Go-http-rest-api/internal/api/model"
	"Go-http-rest-api/internal/api/store"
	"Go-http-rest-api/internal/api/store/teststore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u.ID)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "mynew@email.com"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err,store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email

	s.User().Create(u)
	u, err = s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
