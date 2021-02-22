package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type ShopChatSticker struct {
	ID               uint `gorm:"primary_key"`
	ChatStickerName  string
	ChatStickerImage string
	ChatStickerPoint int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time `sql:"index"`
}

func init() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DropTableIfExists(ShopChatSticker{})
	db.AutoMigrate(ShopChatSticker{})

	SeedShopChatSticker()
}

func SeedShopChatSticker() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&ShopChatSticker{
		ChatStickerName:  "Zuma",
		ChatStickerImage: "sticker1.png",
		ChatStickerPoint: 545,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopChatSticker{
		ChatStickerName:  "Chicky Pox",
		ChatStickerImage: "sticker2.png",
		ChatStickerPoint: 434,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopChatSticker{
		ChatStickerName:  "Shilin",
		ChatStickerImage: "sticker3.png",
		ChatStickerPoint: 111,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopChatSticker{
		ChatStickerName:  "Rice Bowl",
		ChatStickerImage: "sticker4.png",
		ChatStickerPoint: 323,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopChatSticker{
		ChatStickerName:  "Joji",
		ChatStickerImage: "sticker5.png",
		ChatStickerPoint: 545,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopChatSticker{
		ChatStickerName:  "Niki",
		ChatStickerImage: "sticker6.png",
		ChatStickerPoint: 434,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopChatSticker{
		ChatStickerName:  "Ngabbb",
		ChatStickerImage: "sticker7.png",
		ChatStickerPoint: 323,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopChatSticker{
		ChatStickerName:  "Audi Marissa",
		ChatStickerImage: "sticker8.png",
		ChatStickerPoint: 222,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopChatSticker{
		ChatStickerName:  "Jenny",
		ChatStickerImage: "sticker9.png",
		ChatStickerPoint: 543,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopChatSticker{
		ChatStickerName:  "Zina",
		ChatStickerImage: "sticker10.png",
		ChatStickerPoint: 434,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopChatSticker{
		ChatStickerName:  "Densu",
		ChatStickerImage: "sticker11.png",
		ChatStickerPoint: 111,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})
	db.Create(&ShopChatSticker{
		ChatStickerName:  "Takoyaki",
		ChatStickerImage: "sticker12.png",
		ChatStickerPoint: 323,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

}

func GetAllShopChatSticker() ([]ShopChatSticker, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var shopChatSticker []ShopChatSticker
	db.Find(&shopChatSticker)

	return shopChatSticker, nil
}
