package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type ShopBackground struct {
	ID         uint `gorm:"primary_key"`
	BackgroundName  string
	BackgroundImage string
	BackgroundPoint int
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

	db.DropTableIfExists(ShopBackground{})
	db.AutoMigrate(ShopBackground{})

	SeedShopBackground()
}

func SeedShopBackground() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&ShopBackground{
		BackgroundName:  "Zum-Zum",
		BackgroundImage: "background-user3.jpg",
		BackgroundPoint: 1550,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopBackground{
		BackgroundName:  "Letuce",
		BackgroundImage: "background-user4.jpg",
		BackgroundPoint: 334,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopBackground{
		BackgroundName:  "Tomato",
		BackgroundImage: "background-user5.jpg",
		BackgroundPoint: 143,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopBackground{
		BackgroundName:  "Salmon",
		BackgroundImage: "background-user6.jpg",
		BackgroundPoint: 222,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopBackground{
		BackgroundName:  "Trout",
		BackgroundImage: "background-user7.jpg",
		BackgroundPoint: 3232,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopBackground{
		BackgroundName:  "Carrot",
		BackgroundImage: "background-user8.jpg",
		BackgroundPoint: 5454,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopBackground{
		BackgroundName:  "Pumpkin",
		BackgroundImage: "background-user9.jpg",
		BackgroundPoint: 1515,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopBackground{
		BackgroundName:  "Honey",
		BackgroundImage: "background-user10.jpg",
		BackgroundPoint: 1111,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopBackground{
		BackgroundName:  "Ovale",
		BackgroundImage: "background-user11.jpg",
		BackgroundPoint: 4343,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopBackground{
		BackgroundName:  "Soundy",
		BackgroundImage: "background-user12.jpg",
		BackgroundPoint: 322,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopBackground{
		BackgroundName:  "Omega-3",
		BackgroundImage: "background-user13.jpg",
		BackgroundPoint: 222,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopBackground{
		BackgroundName:  "Beef",
		BackgroundImage: "background-user14.jpg",
		BackgroundPoint: 5324,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})


}

func GetAllShopBackground() ([]ShopBackground, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var shopBackground []ShopBackground
	db.Find(&shopBackground)

	return shopBackground, nil
}

