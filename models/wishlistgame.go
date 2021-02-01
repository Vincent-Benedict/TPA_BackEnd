package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type WishlistGame struct{
	ID          uint `gorm:"primary_key"`
	UserID      int
	GameID      int
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

	db.DropTableIfExists(&WishlistGame{})
	db.AutoMigrate(&WishlistGame{})
	SeedWishListGame()
}

func SeedWishListGame(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//db.Create(&WishlistGame{
	//	UserID:    1,
	//	GameID:    1,
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//
	//db.Create(&WishlistGame{
	//	UserID:    1,
	//	GameID:    2,
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
}

func InsertGameToWishlist(jwt string, gameId int) (bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id := getIdFromJWTToken(jwt)

	db.Create(&WishlistGame{
		UserID:    int(id),
		GameID:    gameId,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return true, err
}

func IsAddedGameToWishlist(jwt string, gameId int) (bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id := int(getIdFromJWTToken(jwt))

	var wishlist WishlistGame
	db.Where("user_id = ? and game_id = ?", id, gameId).First(&wishlist)

	if(wishlist.ID != 0){
		return true, nil
	}else{
		return false, nil
	}
}
