package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type UserItem struct {
	ID        uint `gorm:"primary_key"`
	UserId    int
	ItemId    int
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

	db.DropTableIfExists(&UserItem{})
	db.AutoMigrate(&UserItem{})
	SeedUserItem()
}

func SeedUserItem(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// GAME ID 1
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    2,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    3,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    4,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    5,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	// GAME ID 2
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    6,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    7,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    8,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    9,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    10,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    21,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    22,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    23,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    24,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    25,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    26,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    27,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	// GAME ID 3
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    11,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    12,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    13,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    14,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    15,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	// GAME ID 4
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    16,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    17,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    18,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    19,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    1,
		ItemId:    20,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})




	// GAME ID 1
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    2,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    3,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    4,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    5,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	// GAME ID 2
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    6,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    7,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    8,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    9,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    10,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    21,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    22,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    23,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    24,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    25,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    26,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    27,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	// GAME ID 3
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    11,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    12,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    13,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    14,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    15,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	// GAME ID 4
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    16,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    17,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    18,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    19,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserItem{
		UserId:    2,
		ItemId:    20,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}

func GetAllUserItemWithGameId(token string, gameid int)([]UserItem, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))
	var userItem []UserItem
	db.Joins("JOIN game_items ON user_items.item_id = game_items.id").Where("user_items.user_id = ? and game_items.game_id = ?",userid ,gameid).Find((&userItem))

	return userItem, nil
}

func GetUserItemWithGameIdOffsetLimit(token string, gameid int, offset int, limit int)([]UserItem, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))
	var userItem []UserItem
	db.Joins("JOIN game_items ON user_items.item_id = game_items.id").Where("user_items.user_id = ? and game_items.game_id = ?",userid ,gameid).Offset(offset).Limit(limit).Find((&userItem))

	return userItem, nil
}

func GetUserItemWithGameIdByName(token string, gameid int, gamename string)([]UserItem, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))
	var userItem []UserItem
	db.Joins("JOIN game_items ON user_items.item_id = game_items.id").Where("user_items.user_id = ? and game_items.game_id = ? and (game_items.item_name like ? or game_items.item_name like lower(?))",userid ,gameid, "%"+gamename+"%", "%"+gamename+"%").Find(&userItem)

	return userItem, nil
}

func GetUserItemWithGameIdOffsetLimitByName(token string, gameid int, offset int, limit int, gamename string)([]UserItem, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))
	var userItem []UserItem
	db.Joins("JOIN game_items ON user_items.item_id = game_items.id").Where("user_items.user_id = ? and game_items.game_id = ? and (game_items.item_name like ? or game_items.item_name like lower(?))",userid ,gameid, "%"+gamename+"%", "%"+gamename+"%").Offset(offset).Limit(limit).Find(&userItem)

	return userItem, nil
}
