package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type FriendRequest struct {
	ID            uint `gorm:"primary_key"`
	UserId        int
	FriendId      int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}

type FriendRequestIgnored struct {
	ID            uint `gorm:"primary_key"`
	UserId        int
	FriendId      int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}

type ChatNotification struct{
	ID            uint `gorm:"primary_key"`
	UserId        int
	FriendId      int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}

type GiftNotification struct{
	ID            uint `gorm:"primary_key"`
	UserId        int
	FriendId      int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}

type CommentNotification struct{
	ID            uint `gorm:"primary_key"`
	UserId        int
	FriendId      int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}

type ItemNotification struct{
	ID            uint `gorm:"primary_key"`
	UserId        int
	ItemId      int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}

func init() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DropTableIfExists(&FriendRequest{}, &ChatNotification{}, &GiftNotification{}, &CommentNotification{}, &FriendRequestIgnored{}, &ItemNotification{})
	db.AutoMigrate(&FriendRequest{}, &ChatNotification{}, &GiftNotification{}, &CommentNotification{}, &FriendRequestIgnored{}, &ItemNotification{})
}


// FRIEND REQUEST IGNORED
func GetFriendRequestIgnoredById(jwt string)([]FriendRequestIgnored, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid, _:= GetIdFromJwtToken(jwt)

	var friendRequestIgnored []FriendRequestIgnored
	db.Where("friend_id = ?", userid).Find(&friendRequestIgnored)

	return friendRequestIgnored, nil
}

func InsertFriendRequestIgnored(userid int, friendid int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&FriendRequestIgnored{
		UserId:    userid,
		FriendId:  friendid,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return true, nil
}

// FRIEND REQUEST
func GetFriendRequestById(jwt string)([]FriendRequest, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid, _:= GetIdFromJwtToken(jwt)

	var friendRequest []FriendRequest
	db.Where("friend_id = ?", userid).Find(&friendRequest)

	return friendRequest, nil
}

func GetFriendRequestByIdSent(jwt string)([]FriendRequest, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid, _:= GetIdFromJwtToken(jwt)

	var friendRequest []FriendRequest
	db.Where("user_id = ?", userid).Find(&friendRequest)

	return friendRequest, nil
}

func InsertFriendRequest(userid int, friendid int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&FriendRequest{
		UserId:    userid,
		FriendId:  friendid,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return true, nil
}

func DeleteFriendRequest(userid int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Where("user_id = ? or friend_id = ?", userid, userid).Delete(&FriendRequest{})
	return true, nil
}

func DeleteFriendRequest2(userid int, friendid int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Where("user_id = ? and friend_id = ?", userid, friendid).Delete(&FriendRequest{})
	return true, nil
}




// CHAT NOTIFICATION
func GetChatNotificationByToken(jwt string)([]ChatNotification, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid, _:= GetIdFromJwtToken(jwt)

	var chatNotification []ChatNotification
	db.Where("user_id = ?", userid).Find(&chatNotification)

	return chatNotification, nil
}

func InsertChatNotification(userid int, friendid int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&ChatNotification{
		UserId:    userid,
		FriendId:  friendid,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return true, nil
}

func DeleteChatNotification(userid int, friendid int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Where("user_id = ? and friend_id = ?", userid, friendid).Delete(&ChatNotification{})
	return true, nil
}


// GIFT NOTIFICATION
func GetGiftNotificationByToken(jwt string)([]GiftNotification, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid, _:= GetIdFromJwtToken(jwt)

	var giftNotification []GiftNotification
	db.Where("user_id = ?", userid).Find(&giftNotification)

	return giftNotification, nil
}

func InsertGiftNotification(userid int, friendid int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&GiftNotification{
		UserId:    userid,
		FriendId:  friendid,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return true, nil
}

func DeleteGiftNotification(userid int, friendid int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Where("user_id = ? and friend_id = ?", userid, friendid).Delete(&GiftNotification{})
	return true, nil
}


// COMMENT NOTIFICATION
func GetCommentNotificationByToken(jwt string)([]CommentNotification, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid, _:= GetIdFromJwtToken(jwt)

	var commentNotification []CommentNotification
	db.Where("user_id = ?", userid).Find(&commentNotification)

	return commentNotification, nil
}

func InsertCommentNotification(userid int, friendid int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&CommentNotification{
		UserId:    userid,
		FriendId:  friendid,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return true, nil
}

func DeleteCommentNotification(userid int, friendid int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Where("user_id = ? and friend_id = ?", userid, friendid).Delete(&CommentNotification{})
	return true, nil
}

// ITEM NOTIFICATION
func GetItemNotificationByToken(jwt string)([]ItemNotification, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid, _:= GetIdFromJwtToken(jwt)

	var itemNotification []ItemNotification
	db.Where("user_id = ?", userid).Find(&itemNotification)

	return itemNotification, nil
}

func InsertItemNotification(userid int, itemid int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&ItemNotification{
		UserId:    userid,
		ItemId:    itemid,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return true, nil
}

func DeleteItemNotification(id int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Where("id = ?", id).Delete(&ItemNotification{})
	return true, nil
}