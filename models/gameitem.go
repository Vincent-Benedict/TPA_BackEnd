package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type GameItem struct {
	ID              uint `gorm:"primary_key"`
	GameId          int
	ItemName        string
	ItemTransaction int
	ItemPrice       int
	ItemImage       string
	ItemDescription string `gorm:"default:'This Item can be applied to any weapon you own and can be scraped to look more worn. You can scrape the same sticker multiple times, making it a bit more worn each time, until it is removed from the weapon.'"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}

func init() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DropTableIfExists(&GameItem{})
	db.AutoMigrate(&GameItem{})
	SeedGameItem()
}

func SeedGameItem(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&GameItem{
		GameId:          1,
		ItemName:        "2020 RMR Legends",
		ItemTransaction: 156,
		ItemPrice:       300000,
		ItemImage:       "item1.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          1,
		ItemName:        "2020 RMR Challengers",
		ItemTransaction: 172,
		ItemPrice:       300000,
		ItemImage:       "item2.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          1,
		ItemName:        "2020 RMR Contenders",
		ItemTransaction: 300,
		ItemPrice:       200000,
		ItemImage:       "item3.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          1,
		ItemName:        "Operation Broken Fang Case",
		ItemTransaction: 35,
		ItemPrice:       100000,
		ItemImage:       "item4.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          1,
		ItemName:        "Sticker | Natus Vincere | 2020 RMR",
		ItemTransaction: 16,
		ItemPrice:       600000,
		ItemImage:       "item5.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          2,
		ItemName:        "Glove Case",
		ItemTransaction: 456,
		ItemPrice:       500000,
		ItemImage:       "item6.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          2,
		ItemName:        "Sticker | Natus Vincere (Holo) | 2020 RMR",
		ItemTransaction: 126,
		ItemPrice:       400000,
		ItemImage:       "item7.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          2,
		ItemName:        "Fracture Case",
		ItemTransaction: 154,
		ItemPrice:       100000,
		ItemImage:       "item8.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          2,
		ItemName:        "Sticker | Natus Vincere (Gold) | 2020 RMR",
		ItemTransaction: 112,
		ItemPrice:       200000,
		ItemImage:       "item9.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          2,
		ItemName:        "Clutch Case",
		ItemTransaction: 54,
		ItemPrice:       700000,
		ItemImage:       "item10.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	/* 10 item */

	db.Create(&GameItem{
		GameId:          3,
		ItemName:        "UMP-45 | Gold Bismuth",
		ItemTransaction: 125,
		ItemPrice:       400000,
		ItemImage:       "item11.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          3,
		ItemName:        "Sealed Graffiti | Rage Mode (Bazooka Pink)",
		ItemTransaction: 35,
		ItemPrice:       400000,
		ItemImage:       "item12.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          3,
		ItemName:        "FAMAS | Crypsis",
		ItemTransaction: 1336,
		ItemPrice:       500000,
		ItemImage:       "item13.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          3,
		ItemName:        "Bewitched (Foil)",
		ItemTransaction: 256,
		ItemPrice:       700000,
		ItemImage:       "item14.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          3,
		ItemName:        "RAVEN SHOTGUN | Cloaker Gold",
		ItemTransaction: 65,
		ItemPrice:       200000,
		ItemImage:       "item15.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          4,
		ItemName:        "The Steam Awards - 2020 Trading Card",
		ItemTransaction: 556,
		ItemPrice:       100000,
		ItemImage:       "item16.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          4,
		ItemName:        "MAG-7 | Rust Coat",
		ItemTransaction: 75,
		ItemPrice:       900000,
		ItemImage:       "item17.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          4,
		ItemName:        "Outstanding Visual Style",
		ItemTransaction: 77,
		ItemPrice:       100000,
		ItemImage:       "item18.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          4,
		ItemName:        "Sticker | Stone Scales",
		ItemTransaction: 34,
		ItemPrice:       150000,
		ItemImage:       "item19.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          4,
		ItemName:        "Stone Gold Water",
		ItemTransaction: 21,
		ItemPrice:       200000,
		ItemImage:       "item20.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	// 20 ITEM

	db.Create(&GameItem{
		GameId:          2,
		ItemName:        "Sinovac",
		ItemTransaction: 777,
		ItemPrice:       770000,
		ItemImage:       "item21.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          2,
		ItemName:        "Sticker | Natus Vincere (Holo) | 2010 RMR",
		ItemTransaction: 888,
		ItemPrice:       800000,
		ItemImage:       "item22.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          2,
		ItemName:        "Fracture Case 2018",
		ItemTransaction: 133,
		ItemPrice:       150000,
		ItemImage:       "item23.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          2,
		ItemName:        "Sticker | Natus Vincere (Gold) | 2030 RMR",
		ItemTransaction: 323,
		ItemPrice:       200000,
		ItemImage:       "item24.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          2,
		ItemName:        "Clutch Case 2026",
		ItemTransaction: 102,
		ItemPrice:       434,
		ItemImage:       "item25.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          2,
		ItemName:        "UMP-45 | Gold Bismuth 2077",
		ItemTransaction: 1252,
		ItemPrice:       200000,
		ItemImage:       "item26.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          2,
		ItemName:        "Sealed Graffiti | Rage Mode (Bazooka Pink) 3100",
		ItemTransaction: 352,
		ItemPrice:       200000,
		ItemImage:       "item27.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          7,
		ItemName:        "FAMAS | Crypsis 400",
		ItemTransaction: 1326,
		ItemPrice:       400000,
		ItemImage:       "item28.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          7,
		ItemName:        "Bewitched (Foil) 1000",
		ItemTransaction: 434,
		ItemPrice:       800000,
		ItemImage:       "item29.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          8,
		ItemName:        "RAVEN SHOTGUN | Cloaker Gold 2",
		ItemTransaction: 65,
		ItemPrice:       300000,
		ItemImage:       "item30.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          8,
		ItemName:        "The Steam Awards - 2020 Trading Card ",
		ItemTransaction: 321,
		ItemPrice:       500000,
		ItemImage:       "item31.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          8,
		ItemName:        "MAG-7 | Rust Coat 10000",
		ItemTransaction: 240,
		ItemPrice:       600000,
		ItemImage:       "item32.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          8,
		ItemName:        "Outstanding Visual Style 6000",
		ItemTransaction: 77,
		ItemPrice:       100000,
		ItemImage:       "item33.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          8,
		ItemName:        "Sticker Stone Scale 20s",
		ItemTransaction: 75,
		ItemPrice:       200000,
		ItemImage:       "item34.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          9,
		ItemName:        "Stone Gold Water 74",
		ItemTransaction: 40,
		ItemPrice:       800000,
		ItemImage:       "item35.png",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

}

func GetAllGameItem()([]GameItem, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var gameItem []GameItem
	db.Order("item_transaction desc").Find(&gameItem)

	return gameItem, nil
}

func GetGameItemOffsetLimit(offset int, limit int)([]GameItem, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var gameItem []GameItem
	db.Order("item_transaction desc").Limit(limit).Offset(offset).Find(&gameItem)

	return gameItem, nil
}

func GetItemGameById(id int) (GameItem, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var gameItem GameItem
	db.First(&gameItem, id)

	return gameItem, nil
}
