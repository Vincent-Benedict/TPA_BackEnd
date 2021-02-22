package models

import (
	"encoding/base64"
	"fmt"
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"os"
	"time"
)

type User struct {
	ID          uint `gorm:"primary_key"`
	Username    string
	Password    string
	Email       string
	Country     string
	RealName    string
	Balance     int    `gorm:"default:0"`
	Avatar      string `gorm:"default:'avatar-default.jpg'"`
	Level       int    `gorm:"default:0"`
	Summary     string `gorm:"default:'No information given.'"`
	Background  string `gorm:"default:'default.png'"`
	MiniProfile string `gorm:"default:'default.png'"`
	Status      string `gorm:"default:'offline'"`
	CustomUrl   string
	Frame       string `gorm:"default:'framedefault.png'"`
	Theme       string `gorm:"default:'rgba(33, 40, 44, 0.93)'"`
	Badge       string `gorm:"default:'badge2.png'"`
	Point       int    `gorm:"default:0"`
	FriendCode  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
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

func GetUserByFriendCode(friendcode string) (User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var user User
	db.Where("friend_code = ?", friendcode).First(&user)

	return user, nil
}

func InsertUser(username string, password string, email string, country string) (User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	friendcode := generateOTP()

	db.Create(&User{
		Username:   username,
		Password:   password,
		Email:      email,
		Country:    country,
		CustomUrl:  username,
		RealName:   username,
		FriendCode: friendcode,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})

	var user User
	db.Last(&user)

	return user, nil
}

func SeedUser() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&User{
		Username:    "vincentbenedict",
		Password:    "vincentbenedict",
		Email:       "benedictvincentrpj7@gmail.com",
		Country:     "indonesia",
		Avatar:      "avatar-orange.jpg",
		RealName:    "vincent benedict",
		Level:       20,
		Balance:     500000,
		Background:  "background-user1.mp4",
		MiniProfile: "mini-profile8.jpg",
		Status:      "online",
		CustomUrl:   "vincentbenedict",
		Frame:       "frame1.png",
		Badge:       "badge10.png",
		FriendCode:  "AD13A",
		Point:       1000000,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserGame{
		UserId:    1,
		GameId:    1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserGame{
		UserId:    1,
		GameId:    2,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserGame{
		UserId:    1,
		GameId:    3,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserGame{
		UserId:    1,
		GameId:    4,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&RecentActivity{
		UserId:    1,
		Activity:  "vincentbenedict just played Football Manager 2021",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&UserGame{
	//	UserId:    1,
	//	GameId:    2,
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
	//db.Create(&RecentActivity{
	//	UserId:    1,
	//	Activity:  "vincentbenedict just played Football Manager 2021",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//	DeletedAt: nil,
	//})
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
		Username:    "ChronCake",
		Password:    "ChronCake",
		Email:       "chroncake@gmail.com",
		Country:     "indonesia",
		RealName:    "Bobby",
		Avatar:      "avatar-chroncake.jpg",
		Balance:     300000,
		Status:      "online",
		Background:  "background-user2.webm",
		MiniProfile: "mini-profile7.jpg",
		CustomUrl:   "ChronCake",
		FriendCode:  "A1B2A",
		Point:       50000,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserGame{
		UserId:    2,
		GameId:    1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserGame{
		UserId:    2,
		GameId:    2,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserGame{
		UserId:    2,
		GameId:    3,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&UserGame{
		UserId:    2,
		GameId:    4,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
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
	db.Create(&UserGame{
		UserId:    2,
		GameId:    7,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&User{
		Username:    "Daniel",
		Password:    "Daniel",
		Email:       "daniel@gmail.com",
		Country:     "indonesia",
		RealName:    "Daniel",
		Avatar:      "avatar-daniel.jpg",
		MiniProfile: "mini-profile6.jpg",
		Balance:     600000,
		CustomUrl:   "Daniel",
		FriendCode:  "CC13C",
		Point:       20000,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&UserGame{
		UserId:    3,
		GameId:    10,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&User{
		Username:   "KokoPlays",
		Password:   "KokoPlays",
		Email:      "kokoplays@gmail.com",
		Country:    "indonesia",
		RealName:   "Janu",
		Avatar:     "avatar-kokoplays.jpg",
		Balance:    800000,
		CustomUrl:  "KokoPlays",
		FriendCode: "ED1DD",
		Point:      30000,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})
	db.Create(&UserGame{
		UserId:    3,
		GameId:    15,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	// RECENT ACTIVITY BUAT COMMUNITYIMAGEANDVIDEO
	db.Create(&RecentActivity{
		UserId:    1,
		Activity:  "vincentbenedict Posted a Community Image And Video about \"Red Dead Redemption 2 Trailer 2021\"",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&RecentActivity{
		UserId:    2,
		Activity:  "ChronCake Posted a Community Image And Video about \"Narutoooooo !!!\"",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&RecentActivity{
		UserId:    1,
		Activity:  "vincentbenedict Posted a Community Image And Video about \"Counter Strike Strike Again !\"",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&RecentActivity{
		UserId:    3,
		Activity:  "Daniel Posted a Community Image And Video about \"Heroes & General World War II\"",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&RecentActivity{
		UserId:    4,
		Activity:  "KokoPlays Posted a Community Image And Video about \"He seems sad\"",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})



}

func GetUsers()([]User, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var user []User
	db.Find(&user)

	return user, nil
}

func GetUserOffsetLimit(offset int, limit int)([]User, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var user []User
	db.Offset(offset).Limit(limit).Find(&user)

	return user, nil
}

func GetIdFromCustomUrl(customurl string) (int, error) {

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var user User
	db.Where("custom_url = ?", customurl).First(&user)

	return int(user.ID), nil
}

func InsertComment(token string, friendid int, commentdesc string) (Comment, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))

	db.Create(&Comment{
		UserId:      friendid,
		FriendId:    userid,
		CommentDesc: commentdesc,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	var comment Comment
	db.Last(&comment)

	return comment, nil
}

func GetIdFromJwtToken(token string) (int, error) {

	id := int(getIdFromJWTToken(token))

	return id, nil
}

func GetAllCommentsByUserId(userid int) ([]Comment, error) {
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

func DecreasePoint(token string, point int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := uint(getIdFromJWTToken(token))

	db.Model(&User{ID: userid}).UpdateColumn("point", gorm.Expr("point - ?", point))

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

func InsertUserGameById(id int, gameid int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&UserGame{
		UserId:    id,
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

func GetUserGameLimit(token string, limit int) ([]UserGame, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := getIdFromJWTToken(token)

	var userGame []UserGame
	db.Where("user_id = ?", userid).Limit(limit).Find(&userGame)

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

func CheckUserByUsername(username string)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var user User
	db.Where("username = ?", username).First(&user)

	if user.ID == 0{
		return false, nil
	}

	return true, nil
}

// UPDATE PROFILE
func UpdateUserAbout(id int, username string, realname string,
	customurl string, country string, summary string) (bool, error) {

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Model(User{}).Where("id = ?", id).Updates(User{Username: username, RealName: realname, CustomUrl: customurl, Country: country, Summary: summary})

	return true, nil
}

func UpdateUserAvatar(id int, avatar string, frame string) (bool, error) {

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Model(User{}).Where("id = ?", id).Updates(User{Avatar: avatar, Frame: frame})

	return true, nil
}

func UpdateUserBackground(id int, background string) (bool, error) {

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Model(User{}).Where("id = ?", id).Updates(User{Background: background})

	return true, nil
}

func UpdateUserMiniProfile(id int, miniprofile string) (bool, error) {

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Model(User{}).Where("id = ?", id).Updates(User{MiniProfile: miniprofile})

	return true, nil
}

func UpdateUserTheme(id int, theme string) (bool, error) {

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Model(User{}).Where("id = ?", id).Updates(User{Theme: theme})

	return true, nil
}

func UpdateUserBadge(id int, badgename string) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Model(User{}).Where("id = ?", id).Updates(User{Badge: badgename})

	return true, nil
}

func SaveAvatar(encodedBase64Img string, filename string) (bool, error) {
	decoded, _ := base64.StdEncoding.DecodeString(encodedBase64Img)

	writeFilename := "C:\\Vincent\\TPA Web\\Project\\src\\github.com\\Vincent-Benedict\\Angular\\TPA-Web\\src\\assets\\Avatar" + "/" + filename
	errs := ioutil.WriteFile(writeFilename, decoded, os.ModePerm)

	if errs != nil {
		fmt.Print("Error Save File: ")
		fmt.Println(errs)
	}

	return true, nil
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
