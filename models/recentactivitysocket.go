package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type MarketRecentActivity struct {
	ID          uint `gorm:"primary_key"`
	UserId      int
	ItemId      int
	Activity    string
	Price       int
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

	db.DropTableIfExists(&MarketRecentActivity{})
	db.AutoMigrate(&MarketRecentActivity{})
}

func InsertMarketRecentActivity(userid int, itemid int, activity string, price int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&MarketRecentActivity{
		UserId:    userid,
		Activity:  activity,
		ItemId:    itemid,
		Price:     price,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return true, nil
}

func GetLastMarketActivity(itemid int)(MarketRecentActivity, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var activity MarketRecentActivity
	db.Where("item_id = ?", itemid).Last(&activity)

	return activity, nil
}

func GetAllMarketActivity(itemid int)([]MarketRecentActivity, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var activity []MarketRecentActivity
	db.Where("item_id = ?", itemid).Find(&activity)

	return activity, nil
}
