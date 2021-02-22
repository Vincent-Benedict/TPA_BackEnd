package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type ShopMiniBackground struct {
	ID                  uint `gorm:"primary_key"`
	MiniBackgroundName  string
	MiniBackgroundImage string
	MiniBackgroundPoint int
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time `sql:"index"`
}

func init() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DropTableIfExists(ShopMiniBackground{})
	db.AutoMigrate(ShopMiniBackground{})

	SeedShopMiniBackground()
}

func SeedShopMiniBackground() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&ShopMiniBackground{
		MiniBackgroundName:  "NamDo-Sun",
		MiniBackgroundImage: "mini-profile24.jpg",
		MiniBackgroundPoint: 545,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopMiniBackground{
		MiniBackgroundName:  "Salted Egg",
		MiniBackgroundImage: "mini-profile23.jpg",
		MiniBackgroundPoint: 434,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopMiniBackground{
		MiniBackgroundName:  "Indomie",
		MiniBackgroundImage: "mini-profile22.jpg",
		MiniBackgroundPoint: 222,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopMiniBackground{
		MiniBackgroundName:  "Mubarok",
		MiniBackgroundImage: "mini-profile21.jpg",
		MiniBackgroundPoint: 111,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopMiniBackground{
		MiniBackgroundName:  "Magal",
		MiniBackgroundImage: "mini-profile20.jpg",
		MiniBackgroundPoint: 323,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopMiniBackground{
		MiniBackgroundName:  "Japchae",
		MiniBackgroundImage: "mini-profile19.jpg",
		MiniBackgroundPoint: 4343,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopMiniBackground{
		MiniBackgroundName:  "Chicken Wing",
		MiniBackgroundImage: "mini-profile18.jpg",
		MiniBackgroundPoint: 132,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopMiniBackground{
		MiniBackgroundName:  "Sunlight",
		MiniBackgroundImage: "mini-profile17.jpg",
		MiniBackgroundPoint: 333,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopMiniBackground{
		MiniBackgroundName:  "Bean Bozzle",
		MiniBackgroundImage: "mini-profile16.jpg",
		MiniBackgroundPoint: 5454,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopMiniBackground{
		MiniBackgroundName:  "Bamboo",
		MiniBackgroundImage: "mini-profile15.jpg",
		MiniBackgroundPoint: 1111,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopMiniBackground{
		MiniBackgroundName:  "Dan Wik",
		MiniBackgroundImage: "mini-profile14.jpg",
		MiniBackgroundPoint: 2323,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopMiniBackground{
		MiniBackgroundName:  "Jung-On",
		MiniBackgroundImage: "mini-profile13.jpg",
		MiniBackgroundPoint: 232,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})


}

func GetAllShopMiniBackground() ([]ShopMiniBackground, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var shopMiniBackground []ShopMiniBackground
	db.Find(&shopMiniBackground)

	return shopMiniBackground, nil
}
