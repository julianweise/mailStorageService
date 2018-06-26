package models

import (
	"errors"
)

type Mail struct {
	Id				string			`bson:"_id" json:"id"`
	Received 		string			`bson:"received" json:"received,omitempty"`
	ReceivedFrom 	string 			`bson:"received_from" json:"received_from,omitempty"`
	ReceivedBy		string			`bson:"received_by" json:"received_by,omitempty"`
	MailFrom		string			`bson:"mail_from" json:"mail_from,omitempty"`
	RCPTTo			[]string		`bson:"rcpt_to" json:"rcpt_to,omitempty"`
	Data			string			`bson:"data" json:"data,omitempty"`
}

func (mail* Mail) IsValid() error {
	if mail.Received == "" {
		return errors.New("received should not be empty")
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
	if len(mail.RCPTTo) < 1 {
		return errors.New("rcpt_to should not be empty")
	}
	if mail.Data == "" {
		return errors.New("data should not be empty")
	}
	return nil
}

/*
func NewMail(input []byte) (mail Mail, err error) {
	type StringMail struct {
		Received 		string			`json:"received,omitempty"`
		ReceivedFrom 	string 			`json:"received_from,omitempty"`
		ReceivedBy		string			`json:"received_by,omitempty"`
		MailFrom		string			`json:"mail_from,omitempty"`
		RCPTTo			[]string		`json:"rcpt_to,omitempty"`
		Data			string			`json:"data,omitempty"`

	}

	var rawMail StringMail
	mail = Mail{}

	err = json.Unmarshal(input, &rawMail)
	if err != nil {
		return mail, err
	}

	log.Printf("received: '%s'", rawMail)

	mail.Received = time.Now()
	mail.ReceivedBy = rawMail.ReceivedBy
	mail.ReceivedFrom = rawMail.ReceivedFrom
	mail.MailFrom = rawMail.MailFrom
	mail.RCPTTo = rawMail.RCPTTo
	mail.Data = rawMail.Data

	return
}
*/