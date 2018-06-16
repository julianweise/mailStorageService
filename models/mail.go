package models

import "errors"

type Mail struct {
	Received 		string 	`json:"received,omitempty"`
	ReceivedFrom 	string 	`json:"received_from,omitempty"`
	ReceivedBy		string	`json:"received_by,omitempty"`
	MailFrom		string	`json:"mail_from,omitempty"`
	RCPTTo			string	`json:"rcpt_to,omitempty"`
	Data			string	`json:"data,omitempty"`
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