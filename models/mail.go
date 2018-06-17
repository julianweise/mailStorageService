package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type Mail struct {
	Id				bson.ObjectId	`bson:"_id" json:"id"`
	Received 		string 			`bson:"received" json:"received,omitempty"`
	ReceivedFrom 	string 			`bson:"received_from" json:"received_from,omitempty"`
	ReceivedBy		string			`bson:"received_by" json:"received_by,omitempty"`
	MailFrom		string			`bson:"mail_from" json:"mail_from,omitempty"`
	RCPTTo			string			`bson:"rcpt_to" json:"rcpt_to,omitempty"`
	Data			string			`bson:"data" json:"data,omitempty"`
}

func (mail* Mail) IsValid() error {
	if mail.Received == "" {
		return errors.New("receive should not be empty")
	}
	if mail.ReceivedFrom == "" {
		return errors.New("received_from should not be empty")
	}
	if mail.ReceivedBy == "" {
		return errors.New("received_by should not be empty")
	}
	if mail.MailFrom == "" {
		return errors.New("mail_from should not be empty")
	}
	if mail.RCPTTo == "" {
		return errors.New("rcpt_to should not be empty")
	}
	if mail.Data == "" {
		return errors.New("data should not be empty")
	}
	return nil
}