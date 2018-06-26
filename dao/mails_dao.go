package dao

import (
	"gopkg.in/mgo.v2"
	"mailStorageService/models"
	"gopkg.in/mgo.v2/bson"
)

type MailsDAO struct {
	Server		string
	Database	string
}

var db *mgo.Database

const (
	COLLECTION = "mails"
)

func (m *MailsDAO) Connect() (err error) {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		return
	}
	db = session.DB(m.Database)
	return
}

func (m *MailsDAO) Select(queryAttributes bson.M) (mails []models.Mail, err error) {
	err = db.C(COLLECTION).Find(queryAttributtes).All(&mails)
}

func (m *MailsDAO) SelectAll() (mails []models.Mail, err error) {
	err = db.C(COLLECTION).Find(bson.M{}).All(&mails)
	return
}

func (m *MailsDAO) Insert(mail models.Mail) (err error) {
	err = db.C(COLLECTION).Insert(&mail)
	return
}

func (m *MailsDAO) Delete(mail models.Mail) (err error) {
	err = db.C(COLLECTION).Remove(&mail)
	return
}

func (m *MailsDAO) Update(mail models.Mail) (err error) {
	err = db.C(COLLECTION).Update(mail.Id, &mail)
	return
}