package main

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		Name   string
		Email  string
		Result error
	}{
		{"Empty", "", ErrInvalidEmail},
		{"Invalid", "this-is-invalid-email", ErrInvalidEmail},
		{"ValidNoExistent", "this@oh-existent.gg", ErrDomainNotFound},
		{"ValidAndExists", "that@gmail.com", nil},
	}
	for _, item := range tests {
		t.Run(item.Name, func(t *testing.T) {
			assert.Equal(t, item.Result, validateMail(item.Email))
		})
	}
}

func TestValidatePassword(t *testing.T) {
	// Password need to be
	// 9 character or more
	// at least one lowercase
	// at least one uppercase
	// at least one symbol

	tests := []struct {
		Name     string
		Password string
		Result   error
	}{
		{"Empty", "", ErrInvalidPassword},
		{"Short", "xys", ErrInvalidPassword},
		{"Valid", "!Shhfahh6", nil},
	}
	for _, item := range tests {
		t.Run(item.Name, func(t *testing.T) {
			assert.Equal(t, item.Result, validatePassword(item.Password))
		})
	}
}

func TestExtractHref(t *testing.T) {
	twURL, _ := url.Parse("https://twitter.com")
	tests := []struct {
		Name     string
		Fragment string
		Result   *url.URL
		Err      error
	}{
		{"Error", "", nil, ErrNoAnchorFound},
		{"NoAnchor", "<html><head></head><body></body></html>", nil, ErrNoAnchorFound},
		{"SomeAnchor", "<a href=\"https://twitter.com\"></a>", twURL, nil},
	}
	for _, item := range tests {
		t.Run(item.Name, func(t *testing.T) {
			u, err := getURLFromFristAnchor(item.Fragment)
			if assert.Equal(t, item.Err, err) {
				assert.Equal(t, item.Result, u)
			}
		})
	}
}
