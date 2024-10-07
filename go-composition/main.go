package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

type Database interface {
	GetLastPaymentDate(userId string) string
	GetSessionToken(userId string) string
	GetHashedPassword(userId string) string
	LoginUser(userId string, token string)
	GetUserFeatures(userId string) []string
	GetUserTypes(userId string) string
}

var database Database

const (
	Paid  string = "Paid"
	Trial string = "Trial"
	Guest string = "Guest"
)

type User struct {
	ID       string
	Features *[]string
}

type PaidUser struct {
	User
	Auth
	Payment
}

type TrialUser struct {
	User
	Auth
}

type GuestUser struct {
	User
}

type Auth struct {
	SessionToken string
}

type Payment struct {
	SessionToken string
}

func (u *Payment) GetLastPaymentDate(userId string) time.Time {
	dateStr := database.GetLastPaymentDate(userId)
	date, _ := time.Parse("2006-01-02 15:04", dateStr)
	return date
}

func (u *Payment) IsAccountPayed(userId string) bool {
	paymentDate := u.GetLastPaymentDate(userId)
	return time.Now().Sub(paymentDate) < 30*24*time.Hour
}

func (u *Auth) IsLoggedIn(userId string) bool {
	dbToken := database.GetSessionToken(userId)
	return dbToken == u.SessionToken
}

func (u *Auth) LoginUser(userId string, hashedPassword string) {
	if hashedPassword != database.GetHashedPassword(userId) {
		return
	}
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return
	}
	sessionToken := base64.URLEncoding.EncodeToString(token)
	database.LoginUser(userId, sessionToken)
	u.SessionToken = sessionToken
}

func (u *PaidUser) HasFeature(feature string) bool {
	if !u.IsLoggedIn(u.ID) {
		return false
	}
	if !u.IsAccountPayed(u.ID) {
		return false
	}
	if u.Features == nil {
		*u.Features = database.GetUserFeatures(u.ID)
	}
	for _, f := range *u.Features {
		if f == feature {
			return true
		}
	}
	return false
}

func (u *TrialUser) HasFeature(feature string) bool {
	if !u.IsLoggedIn(u.ID) {
		return false
	}
	if u.Features == nil {
		*u.Features = database.GetUserFeatures(u.ID)
	}
	for _, f := range *u.Features {
		if f == feature {
			return true
		}
	}
	return false
}

func (u *GuestUser) HasFeature(feature string) bool {
	if u.Features == nil {
		*u.Features = database.GetUserFeatures(u.ID)
	}
	for _, f := range *u.Features {
		if f == feature {
			return true
		}
	}
	return false
}

type Features interface {
	HasFeature(Feature string) bool
}

func featureCheckHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	var userType string
	if userId == "guest" {
		userType = "Guest"
	} else {
		userType = database.GetUserTypes(userId)
	}
	sessionToken := r.Header.Get("Session-Token")
	var user Features
	switch userType {
	case Paid:
		user = &PaidUser{
			User: User{
				ID: userId,
			},
			Auth: Auth{
				SessionToken: sessionToken,
			},
			Payment: Payment{},
		}
	case Trial:
		user = &TrialUser{
			User: User{
				ID: userId,
			},
			Auth: Auth{
				SessionToken: sessionToken,
			},
		}
	case Guest:
		user = &GuestUser{
			User: User{},
		}
	default:
		http.Error(w, "Unknown user type", http.StatusBadRequest)
		return
	}

	enabled := user.HasFeature(r.URL.Query().Get("feature"))
	fmt.Fprintln(w, enabled)
}

func main() {

}
