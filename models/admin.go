package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Admin struct {
	ID          uint `gorm:"primary_key"`
	Username    string
	Password    string
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

	db.DropTableIfExists(Admin{})
	db.AutoMigrate(Admin{})

	SeedAdmin()
}

func SeedAdmin(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&Admin{
		Username:  "admin1",
		Password:  "admin1",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}

func GetAdmin(username string, password string) (string, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var admin Admin
	db.Where("username = ? and password = ?", username, password).First(&admin)

	if admin.Username == "" {
		return "can't find admin", nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": admin.ID,
	})

	tokenString, _ := token.SignedString([]byte("abc"))

	return tokenString, nil
}
