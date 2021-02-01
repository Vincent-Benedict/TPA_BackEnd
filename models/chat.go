package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type Chat struct {
	ID          uint `gorm:"primary_key"`
	SenderId    int
	RecipientId int
	Message     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

func init() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DropTableIfExists(&Chat{})
	db.AutoMigrate(&Chat{})
}

func InsertChat(chat string, userid int, recipientid int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&Chat{
		SenderId:    userid,
		RecipientId: recipientid,
		Message:     chat,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	return true, nil
}

func GetLastChat(userid int, recipientid int)(Chat, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var chat Chat
	db.Where("(sender_id = ? and recipient_id = ?) or (sender_id = ? and recipient_id = ?)", userid, recipientid, recipientid, userid).Last(&chat)

	return chat, nil
}

func GetAllChat(userid int, recipientid int)([]Chat, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var chat []Chat
	db.Where("(sender_id = ? and recipient_id = ?) or (sender_id = ? and recipient_id = ?)", userid, recipientid, recipientid, userid).Find(&chat)

	return chat, nil
}
