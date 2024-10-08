package auth

import "testing"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}
	if hash == "" {
		t.Error("expected hash to be not empty")
	}
	if hash == "password" {
		t.Error("expected hash to be different form password")
	}
}

func TestComparePasswords(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}
	if !ComparePasswords(hash, "password") {
		t.Errorf("expected password to match hash")
	}
	if ComparePasswords(hash, "notpassowrd") {
		t.Errorf("expected password to not match hash")
	}
}
