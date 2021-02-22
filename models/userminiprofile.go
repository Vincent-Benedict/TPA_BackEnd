package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type UserMiniProfile struct {
	ID          uint `gorm:"primary_key"`
	UserId      int
	MiniProfile string
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

	db.DropTableIfExists(&UserMiniProfile{})
	db.AutoMigrate(&UserMiniProfile{})

	SeedUserMiniProfile()
}

func SeedUserMiniProfile(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&UserMiniProfile{
		UserId:      1,
		MiniProfile: "default.png",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserMiniProfile{
		UserId:      1,
		MiniProfile: "mini-profile10.jpg",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserMiniProfile{
		UserId:      1,
		MiniProfile: "mini-profile11.jpg",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserMiniProfile{
		UserId:      1,
		MiniProfile: "mini-profile8.jpg",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserMiniProfile{
		UserId:      1,
		MiniProfile: "mini-profile4.jpg",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserMiniProfile{
		UserId:      1,
		MiniProfile: "mini-profile9.jpg",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})



	db.Create(&UserMiniProfile{
		UserId:      2,
		MiniProfile: "default.png",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserMiniProfile{
		UserId:      2,
		MiniProfile: "mini-profile1.jpg",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserMiniProfile{
		UserId:      2,
		MiniProfile: "mini-profile2.jpg",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserMiniProfile{
		UserId:      2,
		MiniProfile: "mini-profile3.jpg",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})



	db.Create(&UserMiniProfile{
		UserId:      3,
		MiniProfile: "default.png",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserMiniProfile{
		UserId:      3,
		MiniProfile: "mini-profile4.jpg",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserMiniProfile{
		UserId:      3,
		MiniProfile: "mini-profile5.jpg",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserMiniProfile{
		UserId:      3,
		MiniProfile: "mini-profile6.jpg",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserMiniProfile{
		UserId:      3,
		MiniProfile: "mini-profile7.jpg",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserMiniProfile{
		UserId:      3,
		MiniProfile: "mini-profile8.jpg",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})




	db.Create(&UserMiniProfile{
		UserId:      4,
		MiniProfile: "default.png",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserMiniProfile{
		UserId:      4,
		MiniProfile: "mini-profile9.jpg",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserMiniProfile{
		UserId:      4,
		MiniProfile: "mini-profile10.jpg",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserMiniProfile{
		UserId:      4,
		MiniProfile: "mini-profile11.jpg",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserMiniProfile{
		UserId:      4,
		MiniProfile: "mini-profile12.jpg",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserMiniProfile{
		UserId:      4,
		MiniProfile: "mini-profile8.jpg",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})


}


func InsertUserMiniProfile(token string, miniprofilename string)(bool, error){

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))

	db.Create(&UserMiniProfile{
		UserId:     userid,
		MiniProfile: miniprofilename,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})

	return true, nil
}

