package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type ShopAvatar struct {
	ID          uint `gorm:"primary_key"`
	AvatarName  string
	AvatarImage string
	AvatarPoint int
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

	db.DropTableIfExists(ShopAvatar{})
	db.AutoMigrate(ShopAvatar{})

	SeedShopAvatar()
}

func SeedShopAvatar(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&ShopAvatar{
		AvatarName:  "Stim",
		AvatarImage: "avatar-animated1.gif",
		AvatarPoint: 2347,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopAvatar{
		AvatarName:  "Temp",
		AvatarImage: "avatar-animated2.gif",
		AvatarPoint: 2059,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopAvatar{
		AvatarName:  "Greenlam",
		AvatarImage: "avatar-animated3.gif",
		AvatarPoint: 2031,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopAvatar{
		AvatarName:  "Overhold",
		AvatarImage: "avatar-animated4.gif",
		AvatarPoint: 1503,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopAvatar{
		AvatarName:  "Stronghold",
		AvatarImage: "avatar-animated5.gif",
		AvatarPoint: 322,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopAvatar{
		AvatarName:  "Subin",
		AvatarImage: "avatar-animated6.gif",
		AvatarPoint: 2050,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopAvatar{
		AvatarName:  "Latlux",
		AvatarImage: "avatar-animated7.gif",
		AvatarPoint: 2888,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopAvatar{
		AvatarName:  "Konklux",
		AvatarImage: "avatar-animated8.gif",
		AvatarPoint: 230,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopAvatar{
		AvatarName:  "Tempflex",
		AvatarImage: "avatar-animated9.gif",
		AvatarPoint: 150,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopAvatar{
		AvatarName:  "Ronstring",
		AvatarImage: "avatar-animated10.gif",
		AvatarPoint: 350,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopAvatar{
		AvatarName:  "Lotstring",
		AvatarImage: "avatar-animated11.gif",
		AvatarPoint: 550,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopAvatar{
		AvatarName:  "Ricky",
		AvatarImage: "avatar-animated12.gif",
		AvatarPoint: 1162,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

}

func GetAllShopAvatar()([]ShopAvatar, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var shopAvatar []ShopAvatar
	db.Find(&shopAvatar)

	return shopAvatar, nil
}


