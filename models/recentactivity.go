package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type RecentActivity struct {
	ID          uint `gorm:"primary_key"`
	UserId      int
	Activity   string
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

	db.DropTableIfExists(&RecentActivity{})
	db.AutoMigrate(&RecentActivity{})
}

func GetAllRecentActivity(id int)([]RecentActivity, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var recentActivity []RecentActivity
	db.Where("user_id = ?", id).Find(&recentActivity)

	return recentActivity, nil
}

func InsertRecentActivity(token string, activity string)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))

	db.Create(&RecentActivity{
		UserId:    userid,
		Activity:  activity,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return true, nil
}


func InsertRecentActivityById(userid int, activity string)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&RecentActivity{
		UserId:    userid,
		Activity:  activity,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return true, nil
}

func GetRecentActivityOffsetLimit(id int, offset int, limit int)([]RecentActivity, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var recentActivity []RecentActivity
	db.Where("user_id = ?", id).Order("created_at desc").Limit(limit).Offset(offset).Find(&recentActivity)

	return recentActivity, nil
}