package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type ReviewList struct {
	ID             uint `gorm:"primary_key"`
	UserID         int
	GameID         int
	ReviewID       int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
}

func init() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DropTableIfExists(&ReviewList{})
	db.AutoMigrate(&ReviewList{})

	db.Create(&ReviewList{
		UserID:    1,
		GameID:    1,
		ReviewID:  1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}

func InsertReviewList(token string, gameid int, reviewid int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))

	db.Create(&ReviewList{
		UserID:    userid,
		GameID:    gameid,
		ReviewID:  reviewid,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return true, nil
}

func CheckIsAdded(token string, gameid int, reviewid int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid:= getIdFromJWTToken(token)

	var reviewList ReviewList
	db.Where("user_id = ? and game_id = ? and review_id = ?", userid, gameid, reviewid).First(&reviewList)

	if(reviewList.ID != 0){
		return true, nil
	}

	return false, nil
}
