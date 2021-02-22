package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type UserBackground struct {
	ID         uint `gorm:"primary_key"`
	UserId     int
	Background string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}

func init() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DropTableIfExists(&UserBackground{})
	db.AutoMigrate(&UserBackground{})

	SeedUserBackground()
}

func SeedUserBackground() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&UserBackground{
		UserId:     1,
		Background: "default.png",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})
	db.Create(&UserBackground{
		UserId:     1,
		Background: "background-user1.mp4",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})
	db.Create(&UserBackground{
		UserId:     1,
		Background: "background-user9.jpg",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})
	db.Create(&UserBackground{
		UserId:     1,
		Background: "background-user5.jpg",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})
	db.Create(&UserBackground{
		UserId:     1,
		Background: "background-user3.jpg",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})
	db.Create(&UserBackground{
		UserId:     1,
		Background: "background-user4.jpg",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})


	db.Create(&UserBackground{
		UserId:     2,
		Background: "default.png",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})
	db.Create(&UserBackground{
		UserId:     2,
		Background: "background-user2.webm",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})
	db.Create(&UserBackground{
		UserId:     2,
		Background: "background-user6.jpg",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})
	db.Create(&UserBackground{
		UserId:     2,
		Background: "background-user7.jpg",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})


	db.Create(&UserBackground{
		UserId:     3,
		Background: "default.png",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})
	db.Create(&UserBackground{
		UserId:     3,
		Background: "background-user7.jpg",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})
	db.Create(&UserBackground{
		UserId:     3,
		Background: "background-user8.jpg",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})


	db.Create(&UserBackground{
		UserId:     4,
		Background: "default.png",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})
	db.Create(&UserBackground{
		UserId:     4,
		Background: "background-user10.jpg",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})
	db.Create(&UserBackground{
		UserId:     4,
		Background: "background-user11.jpg",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})

}

func InsertUserBackground(token string, backgroundname string)(bool, error){

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))

	db.Create(&UserBackground{
		UserId:     userid,
		Background: backgroundname,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})

	return true, nil
}
