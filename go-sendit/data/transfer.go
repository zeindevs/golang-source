package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const transferColl = "transfers"

type Transfer struct {
	Link            string             `bson:"link"`
	IsAirdrop       bool               `bson:"isAirdrop"`
	Initiator       string             `bson:"initiator"`
	InitiatorIP     string             `bson:"initiatorIP"`
	IsVerified      bool               `bson:"isVerified"`
	UserID          primitive.ObjectID `bson:"userID"`
	Filename        string             `bson:"filename"`
	ToEmail         string             `bson:"toEmail"`
	Message         string             `bson:"message"`
	From            string             `bson:"from"`
	Hook            string             `bson:"hook"`
	BytesTransfered int                `bson:"bytesTransfered"`
	Expired         bool               `bson:"expired"`
	CreatedAt       int64              `bson:"createdAt"`
	CompletedAt     int64              `bson:"CompletedAt"`
}

func FinalizeTransfer(t Transfer) error {
	return nil
}

func FindTransferByLink(link string) (*Transfer, error) {
	return &Transfer{}, nil
}
