package models

import (
	"fmt"
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID            uint `gorm:"primary_key"`
	Username      string
	Password      string
	Email         string
	Country       string
	Balance       int    `gorm:"default:0"`
	Avatar        string `gorm:"default:'avatar-default'"`
	Level         int    `gorm:"default:0"`
	Summary       string `gorm:"default:'No information given.'"`
	FeaturedBadge string `gorm:"default:'default.png'"`
	Status        string `gorm:"default:'offline'"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}

type UserGame struct {
	ID        uint `gorm:"primary_key"`
	UserId    int
	GameId    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type Comment struct {
	ID          uint `gorm:"primary_key"`
	UserId      int
	FriendId    int
	CommentDesc string
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

	db.DropTableIfExists(&User{}, &UserGame{}, &Comment{})
	db.AutoMigrate(&User{}, &UserGame{}, &Comment{})

	SeedUser()
}

func SeedUser() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&User{
		Username:      "vincentbenedict",
		Password:      "vincentbenedict",
		Email:         "benedictvincentrpj7@gmail.com",
		Country:       "indonesia",
		Avatar:        "avatar-orange",
		Level:         20,
		Balance:       500000,
		FeaturedBadge: "background-user1.mp4",
		Status:        "online",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})
	db.Create(&UserGame{
		UserId:    1,
		GameId:    2,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Comment{
		UserId:      1,
		FriendId:    2,
		CommentDesc: "Heyyy, Your Profile is good !",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&Comment{
		UserId:      1,
		FriendId:    3,
		CommentDesc: "Just The Way You Are, Babe !",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&Comment{
		UserId:      1,
		FriendId:    4,
		CommentDesc: "Don't tell me you're sorry cause you lie !",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&User{
		Username:      "ChronCake",
		Password:      "ChronCake",
		Email:         "chroncake@gmail.com",
		Country:       "indonesia",
		Avatar:        "avatar-chroncake",
		Balance:       300000,
		Status:        "currently-playing",
		FeaturedBadge: "background-user2.webm",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})
	db.Create(&Comment{
		UserId:      2,
		FriendId:    1,
		CommentDesc: "Laskar Pelangi",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&Comment{
		UserId:      2,
		FriendId:    1,
		CommentDesc: "Ku terima Suratmu, dan membelanya",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&User{
		Username:  "Daniel",
		Password:  "Daniel",
		Email:     "daniel@gmail.com",
		Country:   "indonesia",
		Avatar:    "avatar-daniel",
		Balance:   600000,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&User{
		Username:  "KokoPlays",
		Password:  "KokoPlays",
		Email:     "kokoplays@gmail.com",
		Country:   "indonesia",
		Avatar:    "avatar-kokoplays",
		Balance:   800000,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}

func GetIdFromJwtToken(token string)(int, error){

	id := int(getIdFromJWTToken(token))

	return id, nil
}

func GetAllCommentsByUserId(userid int)([]Comment, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var comments []Comment
	db.Where("user_id = ?", userid).Find(&comments)

	return comments, nil
}

func DecreaseBalance(token string, balance int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := uint(getIdFromJWTToken(token))

	db.Model(&User{ID: userid}).UpdateColumn("balance", gorm.Expr("balance - ?", balance))

	return true, nil
}

func IncreaseBalance(token string, balance int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := uint(getIdFromJWTToken(token))

	db.Model(&User{ID: userid}).UpdateColumn("balance", gorm.Expr("balance + ?", balance))

	return true, nil
}

func IsHavingAGame(token string, gameid int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := getIdFromJWTToken(token)

	var userGame UserGame
	db.Where("user_id = ? and game_id = ?", userid, gameid).First(&userGame)

	if userGame.ID != 0 {
		return true, nil
	}

	return false, nil
}

func InsertUserGame(token string, gameid int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))

	db.Create(&UserGame{
		UserId:    userid,
		GameId:    gameid,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return true, nil
}

func GetUserGame(userid int) ([]UserGame, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var userGame []UserGame
	db.Where("user_id = ?", userid).Find(&userGame)

	return userGame, nil
}

//func GetUser(username string, password string)(User, error)  {
//
//	db, err := database.Connect()
//	if err != nil {
//		panic(err)
//	}
//	defer db.Close()
//
//	var user User
//	db.Where("username = ?", username).Where("password = ?", password).First(&user)
//
//	if(user.Username == ""){
//		fmt.Print("Cant' find user")
//	}else{
//		fmt.Print(user)
//	}
//
//
//
//	return user, nil
//}

func GetUser(username string, password string) (string, error) {

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var user User
	db.Where("username = ? and password = ?", username, password).First(&user)

	if user.Username == "" {
		return "can't find user", nil
	}

	//fmt.Print(user)

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte("abc"))

	//claims := jwt.MapClaims{}
	//tes,_ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
	//	return []byte("<YOUR VERIFICATION KEY>"), nil
	//})
	//// ... error handling
	//
	//// do something with decoded claims
	//for key, val := range claims {
	//	fmt.Printf("Key: %v, value: %v\n", key, val)
	//}

	return tokenString, nil
}

func GetUserById(id int) (User, error) {

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var user User
	db.Where("id = ?", id).First(&user)

	return user, nil
}

func GetUserByToken(token string) (User, error) {

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id := getIdFromJWTToken(token)

	var user User
	db.Where("id = ?", id).First(&user)

	return user, nil
}

func getIdFromJWTToken(token string) float64 {

	// DECODE JWT TOKEN
	token1, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("abc"), nil
	})

	var id float64
	if claims, ok := token1.Claims.(jwt.MapClaims); ok && token1.Valid {
		id = claims["userId"].(float64)
		//fmt.Println(claims["userId"])
	}

	return id
}
