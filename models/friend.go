package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type Friend struct {
	ID            uint `gorm:"primary_key"`
	UserId        int
	FriendId      int
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

	db.DropTableIfExists(&Friend{})
	db.AutoMigrate(&Friend{})

	SeedFriend()
}

func SeedFriend(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	db.Create(&Friend{
		UserId:    1,
		FriendId:  2,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserId:    1,
		FriendId:  3,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserId:    1,
		FriendId:  4,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserId:    2,
		FriendId:  1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserId:    3,
		FriendId:  1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserId:    4,
		FriendId:  1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}

func GetFriendsByUserId(userId int)([]Friend, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var friend []Friend
	db.Where("user_id = ?", userId).Find(&friend)

	return friend, nil
}

func IsFriend(userId int, friendId int)(string, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var friend Friend
	db.Where("user_id = ? and friend_id = ?", userId, friendId).First(&friend)

	if friend.ID != 0 {
		return "true" , nil
	}
		return "false" , nil

}

func InsertFriend(userId int, friendId int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&Friend{
		UserId:    userId,
		FriendId:  friendId,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserId:    friendId,
		FriendId:  userId,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return true, nil
}


func GetFriendsByUsername(userid int, username string)([]Friend, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var friend []Friend
	db.Joins("join users on friends.friend_id = users.id and users.username like ? and friends.user_id = ?", "%"+username+"%", userid).Find(&friend)

	return friend, nil
}