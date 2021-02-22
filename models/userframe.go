package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type UserFrame struct {
	ID        uint `gorm:"primary_key"`
	UserId    int
	Frame     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func init() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DropTableIfExists(&UserFrame{})
	db.AutoMigrate(&UserFrame{})

	SeedUserFrame()
}

func SeedUserFrame(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	db.Create(&UserFrame{
		UserId:    1,
		Frame:     "framedefault.png",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserFrame{
		UserId:    1,
		Frame:     "frame1.png",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserFrame{
		UserId:    1,
		Frame:     "frame2.png",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserFrame{
		UserId:    1,
		Frame:     "frame3.png",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserFrame{
		UserId:    1,
		Frame:     "frame4.png",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserFrame{
		UserId:    1,
		Frame:     "frame5.png",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserFrame{
		UserId:    1,
		Frame:     "frame6.png",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})


	db.Create(&UserFrame{
		UserId:    2,
		Frame:     "framedefault.png",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&UserFrame{
		UserId:    3,
		Frame:     "framedefault.png",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}

func GetUserFrameById(userid int)([]UserFrame, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var userFrame []UserFrame
	db.Where("user_id = ?", userid).Find(&userFrame)

	return userFrame, nil
}

func InsertUserFrame(token string, frame string)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))

	db.Create(&UserFrame{
		UserId:    userid,
		Frame:     frame,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return true, nil
}


