package data

import (
	"context"

	"github.com/zeindevs/sendit/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SSHKey struct {
	Key string
}

type Settings struct {
	SSHKeys   map[int]SSHKey
	Subdomain string
}

type User struct {
	Settings Settings
}

var (
	userColl string = "users"
)

func FindUserByID(hex string) (*User, error) {
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return nil, err
	}
	var (
		filter = bson.M{"_id": id}
		user   User
	)
	err = db.Collection(userColl).FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUser(query bson.M) (*User, error) {
	var user User
	err := db.Collection(userColl).FindOne(context.TODO(), query).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func RegisterUserSubdomain(hex string, domain string, sshKey string) error {
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return err
	}
	_ = id
	return nil
}
