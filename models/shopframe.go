package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type ShopFrame struct {
	ID         uint `gorm:"primary_key"`
	FrameName  string
	FrameImage string
	FramePoint int
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

	db.DropTableIfExists(ShopFrame{})
	db.AutoMigrate(ShopFrame{})

	SeedShopFrame()
}

func SeedShopFrame() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&ShopFrame{
		FrameName:  "Zamit",
		FrameImage: "frame1.png",
		FramePoint: 300,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopFrame{
		FrameName:  "Zion",
		FrameImage: "frame3.png",
		FramePoint: 500,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopFrame{
		FrameName:  "Lone",
		FrameImage: "frame4.png",
		FramePoint: 123,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopFrame{
		FrameName:  "Pestin",
		FrameImage: "frame5.png",
		FramePoint: 432,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopFrame{
		FrameName:  "Superse",
		FrameImage: "frame6.png",
		FramePoint: 654,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopFrame{
		FrameName:  "Seabass",
		FrameImage: "frame7.png",
		FramePoint: 412,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopFrame{
		FrameName:  "Joniiss",
		FrameImage: "frame8.png",
		FramePoint: 654,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopFrame{
		FrameName:  "Trim",
		FrameImage: "frame9.png",
		FramePoint: 4444,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopFrame{
		FrameName:  "Zuppa Zoup",
		FrameImage: "frame10.png",
		FramePoint: 4324,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopFrame{
		FrameName:  "Cucumber",
		FrameImage: "frame11.png",
		FramePoint: 111,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopFrame{
		FrameName:  "Fettuce",
		FrameImage: "frame12.png",
		FramePoint: 123,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&ShopFrame{
		FrameName:  "Zuchinni",
		FrameImage: "frame13.png",
		FramePoint: 432,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})


}

func GetAllShopFrame() ([]ShopFrame, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var shopFrame []ShopFrame
	db.Find(&shopFrame)

	return shopFrame, nil
}
