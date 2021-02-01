package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type GameCart struct {
	ID        uint `gorm:"primary_key"`
	GameId    int
	Name      string
	Image     string
	Price     int
	Discount  int `gorm:"default:0"`
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

	db.DropTableIfExists(&GameCart{})
	db.AutoMigrate(&GameCart{})
}

// INSERT
func InsertGameToCart(gameid int, name string, image string, price int, discount int) (GameCart, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var gameCart GameCart
	gameCart.Name = name
	gameCart.Image = image
	gameCart.Price = price
	gameCart.Discount = discount

	db.Create(&GameCart{
		GameId:    gameid,
		Name:      name,
		Image:     image,
		Price:     price,
		Discount:  discount,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return gameCart, err
}

func DeleteGameFromCartById(id int) (GameCart, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var gameCart GameCart
	db.Where("id = ?", id).First(&gameCart)

	db.Delete(&GameCart{}, id)

	return gameCart, err
}

func DeleteAllGameFromCart() (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Delete(&GameCart{})

	return true, err
}

func DeleteGameFromCartById2(id int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Delete(&GameCart{}, id)

	return true, err
}

// SELECT ALL GAMES IN CARTS
func GetGamesInCart() ([]GameCart, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var gamesCart []GameCart
	db.Find(&gamesCart)

	return gamesCart, nil
}
