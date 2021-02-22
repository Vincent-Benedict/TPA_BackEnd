package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/jinzhu/gorm"
	"time"
)

type SellListing struct {
	ID          uint `gorm:"primary_key"`
	UserId      int
	ItemId      int
	Quantity    int
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

type SellQuantityPrice struct {
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

	db.DropTableIfExists(SellListing{})
	db.AutoMigrate(SellListing{})

	SeedSellListing()
}

func SeedSellListing(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&SellListing{
		UserId:    4,
		ItemId:    13,
		Quantity:  2,
		Price:     10000,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&SellListing{
		UserId:    2,
		ItemId:    13,
		Quantity:  5,
		Price:     30000,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&SellListing{
		UserId:    3,
		ItemId:    13,
		Quantity:  3,
		Price:     20000,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&SellListing{
		UserId:    4,
		ItemId:    13,
		Quantity:  8,
		Price:     20000,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}

func GetSellListingQuantityAndPriceByItemId(itemid int)([]SellQuantityPrice, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Model(&SellListing{}).Select("sum(quantity) as qty, price").Where("item_id = ?", itemid).Group("price").Order("price desc").Limit(5).Rows()
	defer rows.Close()

	var results []SellQuantityPrice
	for rows.Next(){
		var result SellQuantityPrice

		if err := rows.Scan(&result.Qty, &result.Price); err != nil{
			return nil, err
		}

		results = append(results, result)
	}
	return results, nil
}

func InsertSellListing(token string, itemid int, quantity int, price int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))

	db.Create(&SellListing{
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

func GetSellListingByUserId(token string, itemid int)([]SellListing, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))

	var sellListing []SellListing
	db.Where("user_id = ? and item_id = ?", userid, itemid).Find(&sellListing)

	return sellListing, nil
}

func CheckIfThereIsSellListing(itemid int, price int, qty int)(int, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var sell SellListing
	db.Where("item_id = ? and quantity = ? and price = ?", itemid, qty, price).Find(&sell)

	if sell.ID == 0{
		return 0, nil
	}

	return 1, nil
}

func BuySoldItem(token string, itemid int, price int, qty int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))

	// Mendapatkan informasi tentang buyer dan user
	var buyer User
	var seller User
	var sellListing SellListing
	db.Where("id = ?", userid).Find(&buyer)
	db.Where("item_id = ? and price = ? and quantity = ?", itemid, price, qty).Find(&sellListing)
	db.Where("id = ?", sellListing.UserId).Find(&seller)

	// Delete dari sellListing
	db.Where("item_id = ? and price = ? and quantity = ?", itemid, price, qty).Delete(SellListing{})
	// Delete dari buyListing
	db.Where("item_id = ? and price = ? and quantity = ?", itemid, price, qty).Delete(BuyListing{})

	// masukkan item ke inventory
	db.Create(&UserItem{
		UserId:    userid,
		ItemId:    itemid,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	// Menghilangkan user item dari seller
	db.Where("user_id = ? and item_id = ?", sellListing.UserId, itemid).Delete(UserItem{})

	// Mengurangi duitnya user dan menambah duitnya seller
	db.Model(&User{}).Where("id =  ?", userid).UpdateColumn("balance", gorm.Expr("balance - ?", price))
	db.Model(&User{}).Where("id =  ?", sellListing.UserId).UpdateColumn("balance", gorm.Expr("balance + ?", price))

	SendEmailSuccessBuy(buyer.Email)
	SendEmailSuccessSell(seller.Email)

	return true, nil
}
