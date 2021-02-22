package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type UnsuspendRequest struct {
	ID         uint `gorm:"primary_key"`
	Username   string
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

	db.DropTableIfExists(&UnsuspendRequest{})
	db.AutoMigrate(&UnsuspendRequest{})

	SeedUnsuspendRequest()
}

func SeedUnsuspendRequest(){
	db,_ := database.Connect()
	defer db.Close()

	db.Create(&UnsuspendRequest{
		Username:  "vincentbenedict",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UnsuspendRequest{
		Username:  "ChronCake",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UnsuspendRequest{
		Username:  "Daniel",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UnsuspendRequest{
		Username:  "KokoPlays",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}

func GetAllUnsuspendRequest()([]UnsuspendRequest, error){
	db,_ := database.Connect()
	defer db.Close()

	var request []UnsuspendRequest
	db.Find(&request)

	return request, nil
}

func InsertUnsuspendRequest(username string)(bool, error){
	db,_ := database.Connect()
	defer db.Close()

	db.Create(&UnsuspendRequest{
		Username:  username,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return true, nil
}

func DeleteUnsuspendRequest(id int)(bool, error){
	db,_ := database.Connect()
	defer db.Close()

	db.Where("id = ?", id).Delete(&UnsuspendRequest{})

	return true, nil
}