package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/jinzhu/gorm"
	"time"
)

type CommunityImageAndVideo struct {
	ID             uint `gorm:"primary_key"`
	UserId         int
	Source         string
	Poster         string
	Description    string
	Liked           int `gorm:"default:0"`
	Disliked        int `gorm:"default:0"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
}

type CommunityImageAndVideoComment struct{
	ID             uint `gorm:"primary_key"`
	UserId         int
	ContentId      int
	Comment        string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
}

func init() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DropTableIfExists(&CommunityImageAndVideo{}, &CommunityImageAndVideoComment{})
	db.AutoMigrate(&CommunityImageAndVideo{}, &CommunityImageAndVideoComment{})
	
	SeedCommunityImageAndVideo()
	SeedCommunityImageAndVideoComment()
}

func SeedCommunityImageAndVideo(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	db.Create(&CommunityImageAndVideo{
		UserId:      1,
		Source:      "https://cdn.cloudflare.steamstatic.com/steam/apps/256768371/movie480.webm?t=1574881352",
		Poster:      "https://cdn.cloudflare.steamstatic.com/steam/apps/1174180/ss_668dafe477743f8b50b818d5bbfcec669e9ba93e.600x338.jpg?t=1606935465",
		Description: "Red Dead Redemption 2 Trailer 2021",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&CommunityImageAndVideo{
		UserId:      2,
		Source:      "https://cdn.cloudflare.steamstatic.com/steam/apps/633230/ss_c4db8eb9899c7e9a652e3121b89e4a5d5fb0a3d6.600x338.jpg?t=1611875135",
		Poster:      "https://cdn.cloudflare.steamstatic.com/steam/apps/633230/ss_c4db8eb9899c7e9a652e3121b89e4a5d5fb0a3d6.600x338.jpg?t=1611875135",
		Description: "Narutoooooo !!!",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&CommunityImageAndVideo{
		UserId:      1,
		Source:      "https://cdn.cloudflare.steamstatic.com/steam/apps/555160/ss_980b80950d073298b9e3b57730fece3020f24534.600x338.jpg?t=1608725963",
		Poster:      "https://cdn.cloudflare.steamstatic.com/steam/apps/555160/ss_980b80950d073298b9e3b57730fece3020f24534.600x338.jpg?t=1608725963",
		Description: "Counter Strike Strike Again !",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&CommunityImageAndVideo{
		UserId:      3,
		Source:      "https://cdn.cloudflare.steamstatic.com/steam/apps/256671253/movie480.webm?t=1562674212",
		Poster:      "https://cdn.cloudflare.steamstatic.com/steam/apps/393380/ss_6fabf783d6897ad15486b2051c997d9f9a8f2ab0.600x338.jpg?t=1605299997",
		Description: "Heroes & General World War II",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&CommunityImageAndVideo{
		UserId:      4,
		Source:      "https://cdn.cloudflare.steamstatic.com/steam/apps/1238810/ss_409e2c952aedae360bb2f64736cad845c3cae510.600x338.jpg?t=1609386727",
		Poster:      "https://cdn.cloudflare.steamstatic.com/steam/apps/1238810/ss_409e2c952aedae360bb2f64736cad845c3cae510.600x338.jpg?t=1609386727",
		Description: "He seems sad",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

}

func SeedCommunityImageAndVideoComment(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&CommunityImageAndVideoComment{
		UserId:    1,
		ContentId: 1,
		Comment:   "A place for new players to find friends to play games with. :)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&CommunityImageAndVideoComment{
		UserId:    2,
		ContentId: 1,
		Comment:   "Hi guys :)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&CommunityImageAndVideoComment{
		UserId:    3,
		ContentId: 1,
		Comment:   "Sent a request.",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&CommunityImageAndVideoComment{
		UserId:    4,
		ContentId: 1,
		Comment:   "Team Fortress 2 [Free!] (hat-based cartoony FPS)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&CommunityImageAndVideoComment{
		UserId:    1,
		ContentId: 1,
		Comment:   "Unturned [free] (blocky zombie survival)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&CommunityImageAndVideoComment{
		UserId:    2,
		ContentId: 1,
		Comment:   "Terraria [like $2.50 on sale] (a classic 2D sidescroller sandbox)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&CommunityImageAndVideoComment{
		UserId:    3,
		ContentId: 1,
		Comment:   "Starbound [$15] (space Terraria, play Terraria first :P)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&CommunityImageAndVideoComment{
		UserId:    4,
		ContentId: 1,
		Comment:   "Don't Starve Together [$15] (quirky survival game)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&CommunityImageAndVideoComment{
		UserId:    1,
		ContentId: 1,
		Comment:   "Viscera Cleanup Detail [$13] (janitor sim)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&CommunityImageAndVideoComment{
		UserId:    2,
		ContentId: 1,
		Comment:   "Sid Meier's Civilization [$60 + DLC] (I think they're on like number 6 now?)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&CommunityImageAndVideoComment{
		UserId:    3,
		ContentId: 1,
		Comment:   "Tropico [$50 +DLC] (quirky dictator sim)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&CommunityImageAndVideoComment{
		UserId:    4,
		ContentId: 1,
		Comment:   "Space Engineers [$20] (voxel based spaceship building game)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&CommunityImageAndVideoComment{
		UserId:    1,
		ContentId: 1,
		Comment:   "ARK: Survival Evolved [$50] (fight/tame dinosaurs, takes up A LOT OF TIME)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})


	db.Create(&CommunityImageAndVideoComment{
		UserId:    1,
		ContentId: 2,
		Comment:   "Warframe [FREE] (space robot ninjas, pretty fun!)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&CommunityImageAndVideoComment{
		UserId:    2,
		ContentId: 2,
		Comment:   "Check my games page to see which games I've spent the most time playing!",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&CommunityImageAndVideoComment{
		UserId:    3,
		ContentId: 2,
		Comment:   "Thanks lol didnt mean to come off as cute but thanks ;)",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&CommunityImageAndVideoComment{
		UserId:    4,
		ContentId: 2,
		Comment:   "Same, minus wife and kids.",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&CommunityImageAndVideoComment{
		UserId:    1,
		ContentId: 2,
		Comment:   "Feel free to add me everyone!:D",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&CommunityImageAndVideoComment{
		UserId:    2,
		ContentId: 2,
		Comment:   "Yo guys looking for new friends and gamers",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

}

func GetCommunityImageAndVideo()([]CommunityImageAndVideo, error){

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var communityImageAndVideo []CommunityImageAndVideo
	db.Find(&communityImageAndVideo)

	return communityImageAndVideo, nil
}

func GetCommunityImageAndVideoComment(contentid int, offset int, limit int)([]CommunityImageAndVideoComment, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var comment []CommunityImageAndVideoComment
	db.Where("content_id = ?", contentid).Limit(limit).Offset(offset).Find(&comment)

	return comment, nil
}

func GetAllCommunityImageAndVideoComment(contentid int)([]CommunityImageAndVideoComment, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var comment []CommunityImageAndVideoComment
	db.Where("content_id = ?", contentid).Find(&comment)

	return comment, nil
}

func InsertCommunityImageAndVideoComment(token string, contentid int, comment string)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))

	db.Create(&CommunityImageAndVideoComment{
		UserId:    userid,
		ContentId: contentid,
		Comment:   comment,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return true, nil
}

func InsertCommunityImageAndVideoLike(contentid int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Model(&CommunityImageAndVideo{ID: uint(contentid)}).UpdateColumn("liked", gorm.Expr("liked + ?", 1))
	return true, nil
}

func InsertCommunityImageAndVideoDislike(contentid int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Model(&CommunityImageAndVideo{ID: uint(contentid)}).UpdateColumn("disliked", gorm.Expr("disliked + ?", 1))
	return true, nil
}