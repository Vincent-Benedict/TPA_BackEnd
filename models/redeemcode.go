package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type RedeemCode struct {
	ID              uint `gorm:"primary_key"`
	Code 			string
	Price           int
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

	db.DropTableIfExists(&RedeemCode{})
	db.AutoMigrate(&RedeemCode{})
	SeedRedeemCode()
}

func SeedRedeemCode(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&RedeemCode{
		Code:      "ABCDE",
		Price:     100000,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&RedeemCode{
		Code:      "FGHIJ",
		Price:     200000,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&RedeemCode{
		Code:      "KLMNO",
		Price:     300000,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&RedeemCode{
		Code:      "PQRST",
		Price:     400000,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&RedeemCode{
		Code:      "UVWXY",
		Price:     500000,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}

func GetRedeemCodeByString(code string)(RedeemCode, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var redeemCode RedeemCode;
	db.Where("code = ?", code).First(&redeemCode)

	return redeemCode, nil
}
