package models

type Mail struct {
	Received 		string 	`json:"received,omitempty"`
	ReceivedFrom 	string 	`json:"received_from,omitempty"`
	ReceivedBy		string	`json:"received_by,omitempty"`
	MailFrom		string	`json:"mail_from,omitempty"`
	RCPTTo			string	`json:"rcpt_to,omitempty"`
	Data			string	`json:"data,omitempty"`
}