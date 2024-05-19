package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("a@gmail.com", "hunter2024")
	assert.Nil(t, err)
	assert.NotNil(t, user.Password)
}

func TestUserPassword(t *testing.T) {
	pw := "hunter2024"
	user, err := NewUser("a@gmail.com", pw)
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword(pw))
	assert.False(t, user.ValidatePassword("hunter2023"))
}
