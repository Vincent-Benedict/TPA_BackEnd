package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type BuyListing struct {
	ID          uint `gorm:"primary_key"`
	UserId      int
	ItemId      int
	Quantity    int
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

type BuyQuantityPrice struct {
	ID          uint `gorm:"primary_key"`
	Qty         int
	Price       int
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

	db.DropTableIfExists(BuyListing{})
	db.AutoMigrate(BuyListing{})

	SeedBuyListing()
}

func SeedBuyListing(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	db.Create(&BuyListing{
		UserId:    1,
		ItemId:    13,
		Quantity:  2,
		Price:     40000,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&BuyListing{
		UserId:    2,
		ItemId:    13,
		Quantity:  3,
		Price:     45000,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&BuyListing{
		UserId:    3,
		ItemId:    13,
		Quantity:  1,
		Price:     48000,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&BuyListing{
		UserId:    4,
		ItemId:    13,
		Quantity:  4,
		Price:     40000,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}

func GetBuyListingQuantityAndPriceByItemId(itemid int)([]BuyQuantityPrice, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Model(&BuyListing{}).Select("sum(quantity) as qty, price").Where("item_id = ?", itemid).Group("price").Order("price desc").Limit(5).Rows()
	defer rows.Close()

	var results []BuyQuantityPrice
	for rows.Next(){
		var result BuyQuantityPrice

		if err := rows.Scan(&result.Qty, &result.Price); err != nil{
			return nil, err
		}

		results = append(results, result)
	}
	return results, nil
}

func InsertBuyListing(token string, itemid int, quantity int, price int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))

	db.Create(&BuyListing{
		UserId:    userid,
		ItemId:    itemid,
		Quantity:  quantity,
		Price:     price,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return true, nil
}

func GetBuyListingByUserId(token string, itemid int)([]BuyListing, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))

	var buyListing []BuyListing
	db.Where("user_id = ? and item_id = ?", userid, itemid).Find(&buyListing)

	return buyListing, nil
}