package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type UserTheme struct {
	ID        uint `gorm:"primary_key"`
	Theme     string
	Color     string
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

	db.DropTableIfExists(&UserTheme{})
	db.AutoMigrate(&UserTheme{})

	SeedUserTheme()
}

func SeedUserTheme(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	db.Create(&UserTheme{
		Theme:     "DEFAULT",
		Color:     "rgba(33, 40, 44, 0.93)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserTheme{
		Theme:     "SUMMER",
		Color:     "rgba(181, 114, 19, 0.8)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserTheme{
		Theme:     "MIDNIGHT",
		Color:     "rgba(31, 29, 105, 0.8)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserTheme{
		Theme:     "STEEL",
		Color:     "rgba(61, 74, 89, 0.8)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserTheme{
		Theme:     "COSMIC",
		Color:     "rgba(119, 36, 120, 0.8)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserTheme{
		Theme:     "DARKMODE",
		Color:     "rgbA(17, 17, 17, 0.85)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}

func GetAllTheme()([]UserTheme, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var userTheme []UserTheme
	db.Find(&userTheme)

	return userTheme, nil
}

func GetThemeById(id int)(UserTheme, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var userTheme UserTheme
	db.Where("id = ?", id).First(&userTheme)

	return userTheme, nil
}