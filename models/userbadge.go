package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type UserBadge struct {
	ID        uint `gorm:"primary_key"`
	UserId    int
	BadgeImg  string
	BadgeName string
	BadgeXp   int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type BadgeCard struct {
	ID          uint `gorm:"primary_key"`
	UserBadgeId int
	Card        string
	Status      string
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

	db.DropTableIfExists(&UserBadge{}, &BadgeCard{})
	db.AutoMigrate(&UserBadge{}, &BadgeCard{})

	SeedUserBadge()
}

func SeedUserBadge() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&UserBadge{
		UserId:    1,
		BadgeImg:  "badge10.png",
		BadgeName: "Ganker",
		BadgeXp:   500,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&BadgeCard{
		UserBadgeId: 1,
		Card:        "card1.png",
		Status:      "not-owned",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&BadgeCard{
		UserBadgeId: 1,
		Card:        "card2.png",
		Status:      "owned",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&BadgeCard{
		UserBadgeId: 1,
		Card:        "card3.png",
		Status:      "not-owned",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&BadgeCard{
		UserBadgeId: 1,
		Card:        "card4.png",
		Status:      "owned",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&UserBadge{
		UserId:    1,
		BadgeImg:  "badge1.png",
		BadgeName: "Community Contributor",
		BadgeXp:   200,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&BadgeCard{
		UserBadgeId: 2,
		Card:        "card5.png",
		Status:      "owned",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&BadgeCard{
		UserBadgeId: 2,
		Card:        "card6.png",
		Status:      "not-owned",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&BadgeCard{
		UserBadgeId: 2,
		Card:        "card7.png",
		Status:      "owned",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&BadgeCard{
		UserBadgeId: 2,
		Card:        "card8.png",
		Status:      "not-owned",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&UserBadge{
		UserId:    1,
		BadgeImg:  "badge2.png",
		BadgeName: "Addept Acumulator",
		BadgeXp:   100,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&BadgeCard{
		UserBadgeId: 3,
		Card:        "card9.png",
		Status:      "not-owned",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&BadgeCard{
		UserBadgeId: 3,
		Card:        "card10.png",
		Status:      "owned",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&BadgeCard{
		UserBadgeId: 3,
		Card:        "card11.png",
		Status:      "owned",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&BadgeCard{
		UserBadgeId: 3,
		Card:        "card12.png",
		Status:      "not-owned",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&UserBadge{
		UserId:    1,
		BadgeImg:  "badge3.png",
		BadgeName: "Years of Service",
		BadgeXp:   450,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&BadgeCard{
		UserBadgeId: 4,
		Card:        "card13.png",
		Status:      "owned",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&BadgeCard{
		UserBadgeId: 4,
		Card:        "card14.png",
		Status:      "owned",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&BadgeCard{
		UserBadgeId: 4,
		Card:        "card15.png",
		Status:      "not-owned",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&BadgeCard{
		UserBadgeId: 4,
		Card:        "card16.png",
		Status:      "not-owned",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&UserBadge{
		UserId:    1,
		BadgeImg:  "badge4.png",
		BadgeName: "Winter 2019 Badge",
		BadgeXp:   800,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&UserBadge{
		UserId:    2,
		BadgeImg:  "badge2.png",
		BadgeName: "Addept Acumulator",
		BadgeXp:   100,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserBadge{
		UserId:    2,
		BadgeImg:  "badge5.png",
		BadgeName: "Steam Grand Prix 2019",
		BadgeXp:   450,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserBadge{
		UserId:    2,
		BadgeImg:  "badge6.png",
		BadgeName: "Steam Grand Prix 2020",
		BadgeXp:   800,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&UserBadge{
		UserId:    3,
		BadgeImg:  "badge2.png",
		BadgeName: "Addept Acumulator",
		BadgeXp:   100,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserBadge{
		UserId:    3,
		BadgeImg:  "badge7.png",
		BadgeName: "Star Player",
		BadgeXp:   650,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserBadge{
		UserId:    3,
		BadgeImg:  "badge8.png",
		BadgeName: "Global Sentinel",
		BadgeXp:   100,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&UserBadge{
		UserId:    4,
		BadgeImg:  "badge2.png",
		BadgeName: "Addept Acumulator",
		BadgeXp:   100,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserBadge{
		UserId:    4,
		BadgeImg:  "badge9.png",
		BadgeName: "Hard Carry",
		BadgeXp:   20,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserBadge{
		UserId:    4,
		BadgeImg:  "badge10.png",
		BadgeName: "Ganker",
		BadgeXp:   150,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}

func GetAllUserBadgeByUserId(token string) ([]UserBadge, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := getIdFromJWTToken(token)

	var userBadge []UserBadge
	db.Where("user_id = ?", userid).Find(&userBadge)

	return userBadge, nil
}

func GetUserBadgeById(id int) (UserBadge, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var userBadge UserBadge
	db.Where("id = ?", id).First(&userBadge)

	return userBadge, nil
}

func GetBadgeCardById(id int) ([]BadgeCard, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var badgeCard []BadgeCard
	db.Where("user_badge_id = ?", id).Find(&badgeCard)

	return badgeCard, nil
}
