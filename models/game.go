package models

import (
	"encoding/base64"
	"fmt"
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"os"
	"time"
)

type Game struct {
	ID            uint `gorm:"primary_key"`
	Name          string
	Image         string
	Category      string
	Price         int64
	Discount      int32 `gorm:"default:0"`
	SideImage     string
	Overall       string
	Status        string
	Description   string
	Developer     string
	Publisher     string
	Inappropriate string
	GameplayHours int `gorm:"default:0"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}

type Review struct {
	ID              uint `gorm:"primary_key"`
	GameID          uint
	ReviewDesc      string
	ReviewUsername  string
	ReviewAvatar    string
	ReviewUpvoted   int `gorm:"default:0"`
	ReviewDownvoted int `gorm:"default:0"`
	ReviewSentiment string `gorm:"default: 'positive'"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}

type ReviewComment struct{
	ID              uint `gorm:"primary_key"`
	UserId          int
	ReviewId        int
	Comment         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}

type VideoGame struct {
	ID          uint `gorm:"primary_key"`
	GameID      uint
	VideoSource string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

type PhotoGame struct {
	ID          uint `gorm:"primary_key"`
	GameID      uint
	PhotoSource string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

type SystemRequirement struct {
	ID        uint
	GameID    uint
	Os        string
	Processor string
	Memory    string
	Graphics  string
	Storage   string
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

	db.DropTableIfExists(&Game{}, &Review{}, &ReviewComment{}, &VideoGame{}, &PhotoGame{}, &SystemRequirement{})
	db.AutoMigrate(&Game{}, &Review{}, &ReviewComment{}, &VideoGame{}, &PhotoGame{}, &SystemRequirement{})
	SeedGameRecommended()
	SeedGameSpecialOffer()
	SeedGameCommunityRecommended()
	SeedGameTopSellers()
	SeedGameSpecials()
	SeedGameNewAndTrending()
}

func SeedGameRecommended() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// CREATE RECOMMENDED GAME
	// ID 1-8

	db.Create(&Game{
		Name:          "CyberPunk 2077",
		Image:         "Featured&Recommended2",
		Category:      "Action",
		Price:         173999,
		Discount:      0,
		SideImage:     "Featured&Recommended2",
		Overall:       "4.0",
		Status:        "recommend",
		Description:   "Cyberpunk 2077 is an open-world, action-adventure story set in Night City, a megalopolis obsessed with power, glamour and body modification. You play as V, a mercenary outlaw going after a one-of-a-kind implant that is the key to immortality.",
		Developer:     "CD PROJEKT RED",
		Publisher:     "CD PROJEKT RED",
		Inappropriate: "yes",
		GameplayHours: 300,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})
	db.Create(&SystemRequirement{
		GameID:    1,
		Os:        "Windows 10 64-bit",
		Processor: "Intel Core i7-4770K@3.5GHz or Ryzen 5 1500X@3.5GHz",
		Memory:    "16 GB RAM",
		Graphics:  "Nvidia GeForce GTX 1060 (6 GB) or AMD Radeon RX 580 (8GB)",
		Storage:   "50 GB available space",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Review{
		GameID:         1,
		ReviewDesc:     "**Did not play beta or alpha or any version of the game prior to 01/18/2021**\n\nThis is it.\n\nIt's not fully baked but boy oh boy is it delicious.\n\nI grew up playing Wing Commander 3 in the late 90's and Freelancer in the early 2000's, birthing my love for space combat RPG's.\n\nI've been chasing the dragon ever since.\n\nI remember trying Everspace for the first time and it was not my cup of tea as a rogue-like. But it was so close to what I was searching for.\n\nEverspace 2 is it.\n\nIt's in EA folks. You should set expectations as such, if you want to support the project. It's incredibly playable, feature rich, and gorgeous in its current state.\n\nIt's fun to play, and wow did Rockfish games get a lot right.\n\nI cannot WAIT to see what is in store for the future of development.\n\nI've had one crash so far, and it was expected, as I was trying to get real fancy exiting cruise speed trying to pull off a disgusting amount of ship spins. My GTX 1070 Rog Strix OC edition running everything at EPIC settings did it's best.\n\nOtherwise the game runs very well for day 1 of EA. No bugs experienced or noticed so far, besides one crash due to pushing the limits of my system.\n\nI have no doubt with some time and more love from the Dev team, this game is going to be the start of something great. Perhaps even a new era of space combat rpg's if other studios take notice.\n\nGreat job Rockfish games.\nFly safe everyone. o7",
		ReviewUsername: "ChronCake",
		ReviewAvatar:   "avatar-chroncake.jpg",
		ReviewUpvoted:  10,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&ReviewComment{
		UserId:    1,
		ReviewId:  1,
		Comment:   "Hello",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&ReviewComment{
		UserId:    2,
		ReviewId:  1,
		Comment:   "Huhhhh",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&ReviewComment{
		UserId:    3,
		ReviewId:  1,
		Comment:   "eweewe",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&ReviewComment{
		UserId:    4,
		ReviewId:  1,
		Comment:   "asdfjklm",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&ReviewComment{
		UserId:    1,
		ReviewId:  1,
		Comment:   "About a Thing",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&ReviewComment{
		UserId:    2,
		ReviewId:  1,
		Comment:   "Be Happy",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&ReviewComment{
		UserId:    3,
		ReviewId:  1,
		Comment:   "Don't Worry",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&ReviewComment{
		UserId:    4,
		ReviewId:  1,
		Comment:   "Avengers Assemble",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&ReviewComment{
		UserId:    1,
		ReviewId:  1,
		Comment:   "Noob Master",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&ReviewComment{
		UserId:    2,
		ReviewId:  1,
		Comment:   "Wo Jiao Cin Ci",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&ReviewComment{
		UserId:    3,
		ReviewId:  1,
		Comment:   "Good Morning",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&ReviewComment{
		UserId:    4,
		ReviewId:  1,
		Comment:   "Haiiii Neighbour",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&ReviewComment{
		UserId:    1,
		ReviewId:  1,
		Comment:   "asdfeg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Review{
		GameID:         1,
		ReviewDesc:     "First space RPG I've played where the ship controls like a dream. Systems are simple to understand but seem like they have a considerable amount of depth.\n\nUI is simple and navigating is easy. I'm constantly finding hidden secrets that take me off my main path. So far so good!",
		ReviewUsername: "Daniel",
		ReviewAvatar:   "avatar-daniel.jpg",
		ReviewUpvoted:  5,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&ReviewComment{
		UserId:    4,
		ReviewId:  2,
		Comment:   "Senja Senja huhuhuhuh...",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&ReviewComment{
		UserId:    1,
		ReviewId:  2,
		Comment:   "Mentari Pagi",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&ReviewComment{
		UserId:    2,
		ReviewId:  2,
		Comment:   "Cin Cau Balsem",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&ReviewComment{
		UserId:    3,
		ReviewId:  2,
		Comment:   "Spongeeeboobbb",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&ReviewComment{
		UserId:    4,
		ReviewId:  2,
		Comment:   "Squidword",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&ReviewComment{
		UserId:    1,
		ReviewId:  2,
		Comment:   "Hail Plankton",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Review{
		GameID:         1,
		ReviewDesc:     "Combat flows nicely, the space is huge and there's always something to explore and the story is solid so far. Extremely well polished for an early access product!",
		ReviewUsername: "KokoPlays",
		ReviewAvatar:   "avatar-kokoplays.jpg",
		ReviewUpvoted:  0,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&VideoGame{
		GameID:      1,
		VideoSource: "https://cdn.cloudflare.steamstatic.com/steam/apps/256805123/movie480_vp9.webm?t=1602863731",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&VideoGame{
		GameID:      1,
		VideoSource: "https://cdn.cloudflare.steamstatic.com/steam/apps/256812920/movie480_vp9.webm?t=1607518395",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      1,
		PhotoSource: "https://cdn.cloudflare.steamstatic.com/steam/apps/256805123/movie.293x165.jpg?t=1602863731",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      1,
		PhotoSource: "https://cdn.cloudflare.steamstatic.com/steam/apps/256812920/movie.293x165.jpg?t=1607518395",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      1,
		PhotoSource: "https://cdn.cloudflare.steamstatic.com/steam/apps/1091500/ss_2171e3f8deefaf544b27710ad127ffa26d3f712c.600x338.jpg?t=1608552868",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Name:          "Football Manager 2021",
		Image:         "Featured&Recommended1",
		Category:      "Sport",
		Price:         100000,
		Discount:      0,
		SideImage:     "Featured&Recommended1",
		Overall:       "4.0",
		Status:        "recommend",
		Description:   "New additions and game upgrades deliver added levels of depth, drama and football authenticity. FM21 empowers you like never before to develop your skills and command success at your club.",
		Developer:     "Sports Interactive",
		Publisher:     "SEGA",
		Inappropriate: "no",
		GameplayHours: 200,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})
	db.Create(&SystemRequirement{
		GameID:    2,
		Os:        "Windows 10 64-bit",
		Processor: "Intel Core i7-4770K@3.5GHz or Ryzen 5 1500X@3.5GHz",
		Memory:    "16 GB RAM",
		Graphics:  "Nvidia GeForce GTX 1060 (6 GB) or AMD Radeon RX 580 (8GB)",
		Storage:   "50 GB available space",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Review{
		GameID:         2,
		ReviewDesc:     "**Did not play beta or alpha or any version of the game prior to 01/18/2021**\n\nThis is it.\n\nIt's not fully baked but boy oh boy is it delicious.\n\nI grew up playing Wing Commander 3 in the late 90's and Freelancer in the early 2000's, birthing my love for space combat RPG's.\n\nI've been chasing the dragon ever since.\n\nI remember trying Everspace for the first time and it was not my cup of tea as a rogue-like. But it was so close to what I was searching for.\n\nEverspace 2 is it.\n\nIt's in EA folks. You should set expectations as such, if you want to support the project. It's incredibly playable, feature rich, and gorgeous in its current state.\n\nIt's fun to play, and wow did Rockfish games get a lot right.\n\nI cannot WAIT to see what is in store for the future of development.\n\nI've had one crash so far, and it was expected, as I was trying to get real fancy exiting cruise speed trying to pull off a disgusting amount of ship spins. My GTX 1070 Rog Strix OC edition running everything at EPIC settings did it's best.\n\nOtherwise the game runs very well for day 1 of EA. No bugs experienced or noticed so far, besides one crash due to pushing the limits of my system.\n\nI have no doubt with some time and more love from the Dev team, this game is going to be the start of something great. Perhaps even a new era of space combat rpg's if other studios take notice.\n\nGreat job Rockfish games.\nFly safe everyone. o7",
		ReviewUsername: "ChronCake",
		ReviewAvatar:   "avatar-chroncake.jpg",
		ReviewUpvoted:  10,
		ReviewSentiment: "negative",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&Review{
		GameID:         2,
		ReviewDesc:     "First space RPG I've played where the ship controls like a dream. Systems are simple to understand but seem like they have a considerable amount of depth.\n\nUI is simple and navigating is easy. I'm constantly finding hidden secrets that take me off my main path. So far so good!",
		ReviewUsername: "Daniel",
		ReviewAvatar:   "avatar-daniel.jpg",
		ReviewUpvoted:  5,
		ReviewSentiment: "negative",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&Review{
		GameID:         2,
		ReviewDesc:     "Combat flows nicely, the space is huge and there's always something to explore and the story is solid so far. Extremely well polished for an early access product!",
		ReviewUsername: "KokoPlays",
		ReviewAvatar:   "avatar-kokoplays.jpg",
		ReviewUpvoted:  0,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&VideoGame{
		GameID:      2,
		VideoSource: "https://cdn.akamai.steamstatic.com/steam/apps/256803629/movie480_vp9.webm?t=1601922367",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&VideoGame{
		GameID:      2,
		VideoSource: "https://cdn.akamai.steamstatic.com/steam/apps/256808959/movie480_vp9.webm?t=1605059339",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      2,
		PhotoSource: "https://cdn.akamai.steamstatic.com/steam/apps/1263850/ss_435711f30c32974d611c7015aaa04b5888ae1a83.600x338.jpg?t=1606827432",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      2,
		PhotoSource: "https://cdn.akamai.steamstatic.com/steam/apps/1263850/ss_fd42ad95cb1a3e7d9e48ef393362147210020ba7.600x338.jpg?t=1606827432",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      2,
		PhotoSource: "https://cdn.akamai.steamstatic.com/steam/apps/1263850/ss_0e08de5f60e5a2465c896015368b66ac69056d30.600x338.jpg?t=1606827432",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Name:          "Phasmophobia",
		Image:         "Featured&Recommended3",
		Category:      "Horror",
		Price:         201999,
		Discount:      0,
		SideImage:     "Featured&Recommended3",
		Overall:       "4.5",
		Status:        "recommend",
		Description:   "Phasmophobia is a 4 player online co-op psychological horror. Paranormal activity is on the rise and it’s up to you and your team to use all the ghost hunting equipment at your disposal in order to gather as much evidence as you can.",
		Developer:     "Kinetic Games",
		Publisher:     "Kinetic Games",
		Inappropriate: "no",
		GameplayHours: 250,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})
	db.Create(&SystemRequirement{
		GameID:    3,
		Os:        "Windows 10 64-bit",
		Processor: "Intel Core i7-4770K@3.5GHz or Ryzen 5 1500X@3.5GHz",
		Memory:    "16 GB RAM",
		Graphics:  "Nvidia GeForce GTX 1060 (6 GB) or AMD Radeon RX 580 (8GB)",
		Storage:   "50 GB available space",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Review{
		GameID:         3,
		ReviewDesc:     "Oddly enough, the phrase \"hunt me daddy\" is a very effective way to get a lot of activity.",
		ReviewUsername: "IFightBears",
		ReviewAvatar:   "avatar-ifightbears.jpg",
		ReviewUpvoted:  9,
		ReviewSentiment: "negative",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&Review{
		GameID:         3,
		ReviewDesc:     "No cheap jumpscares, a genuine creepy and dark atmosphere, and the game is actually scary when you play it alone. I'd say this is horror done right.",
		ReviewUsername: "ScaredVx",
		ReviewAvatar:   "avatar-scaredvx.jpg",
		ReviewUpvoted:  8,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&Review{
		GameID:         3,
		ReviewDesc:     "Scary and fun at the same time!",
		ReviewUsername: "HeadWound1",
		ReviewAvatar:   "avatar-headwoundharry.jpg",
		ReviewUpvoted:  0,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&VideoGame{
		GameID:      3,
		VideoSource: "https://cdn.cloudflare.steamstatic.com/steam/apps/256776660/movie480.webm?t=1583525286",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&VideoGame{
		GameID:      3,
		VideoSource: "https://cdn.cloudflare.steamstatic.com/steam/apps/256776660/movie480.webm?t=1583525286",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      3,
		PhotoSource: "https://cdn.cloudflare.steamstatic.com/steam/apps/739630/ss_91bca60a51dce60d680a8fb4efcdecf740b3a3d1.600x338.jpg?t=1611245377",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      3,
		PhotoSource: "https://cdn.cloudflare.steamstatic.com/steam/apps/739630/ss_277778876e52adc03b3261213e45a06b2e8cd28c.600x338.jpg?t=1611245377",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      3,
		PhotoSource: "https://cdn.cloudflare.steamstatic.com/steam/apps/739630/ss_d48dc33b2d49086841d8361c59752cb1bb3733ac.600x338.jpg?t=1611245377",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Name:          "Raft",
		Image:         "Featured&Recommended4",
		Category:      "Adventure",
		Price:         140000,
		Discount:      0,
		SideImage:     "Featured&Recommended4",
		Overall:       "4.0",
		Status:        "recommend",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		GameplayHours: 200,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})
	db.Create(&SystemRequirement{
		GameID:    4,
		Os:        "Windows 10 64-bit",
		Processor: "Intel Core i7-4770K@3.5GHz or Ryzen 5 1500X@3.5GHz",
		Memory:    "16 GB RAM",
		Graphics:  "Nvidia GeForce GTX 1060 (6 GB) or AMD Radeon RX 580 (8GB)",
		Storage:   "50 GB available space",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Review{
		GameID:         4,
		ReviewDesc:     "**Did not play beta or alpha or any version of the game prior to 01/18/2021**\n\nThis is it.\n\nIt's not fully baked but boy oh boy is it delicious.\n\nI grew up playing Wing Commander 3 in the late 90's and Freelancer in the early 2000's, birthing my love for space combat RPG's.\n\nI've been chasing the dragon ever since.\n\nI remember trying Everspace for the first time and it was not my cup of tea as a rogue-like. But it was so close to what I was searching for.\n\nEverspace 2 is it.\n\nIt's in EA folks. You should set expectations as such, if you want to support the project. It's incredibly playable, feature rich, and gorgeous in its current state.\n\nIt's fun to play, and wow did Rockfish games get a lot right.\n\nI cannot WAIT to see what is in store for the future of development.\n\nI've had one crash so far, and it was expected, as I was trying to get real fancy exiting cruise speed trying to pull off a disgusting amount of ship spins. My GTX 1070 Rog Strix OC edition running everything at EPIC settings did it's best.\n\nOtherwise the game runs very well for day 1 of EA. No bugs experienced or noticed so far, besides one crash due to pushing the limits of my system.\n\nI have no doubt with some time and more love from the Dev team, this game is going to be the start of something great. Perhaps even a new era of space combat rpg's if other studios take notice.\n\nGreat job Rockfish games.\nFly safe everyone. o7",
		ReviewUsername: "ChronCake",
		ReviewAvatar:   "avatar-chroncake.jpg",
		ReviewUpvoted:  10,
		ReviewSentiment: "negative",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&Review{
		GameID:         4,
		ReviewDesc:     "First space RPG I've played where the ship controls like a dream. Systems are simple to understand but seem like they have a considerable amount of depth.\n\nUI is simple and navigating is easy. I'm constantly finding hidden secrets that take me off my main path. So far so good!",
		ReviewUsername: "Daniel",
		ReviewAvatar:   "avatar-daniel.jpg",
		ReviewUpvoted:  5,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&Review{
		GameID:         4,
		ReviewDesc:     "Combat flows nicely, the space is huge and there's always something to explore and the story is solid so far. Extremely well polished for an early access product!",
		ReviewUsername: "KokoPlays",
		ReviewAvatar:   "avatar-kokoplays.jpg",
		ReviewUpvoted:  0,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&VideoGame{
		GameID:      4,
		VideoSource: "https://cdn.akamai.steamstatic.com/steam/apps/256804090/movie480_vp9.webm?t=1602176624",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&VideoGame{
		GameID:      4,
		VideoSource: "https://cdn.akamai.steamstatic.com/steam/apps/256768464/movie480.webm?t=1574953141",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      4,
		PhotoSource: "https://cdn.akamai.steamstatic.com/steam/apps/648800/ss_c22b2ff5ba5609f74e61b5feaa5b7a1d7fd1dbd3.600x338.jpg?t=1602795811",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      4,
		PhotoSource: "https://cdn.akamai.steamstatic.com/steam/apps/648800/ss_9784377e0e3bcfd2be609721326ab336a39c34b8.600x338.jpg?t=1602795811",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      4,
		PhotoSource: "https://cdn.akamai.steamstatic.com/steam/apps/648800/ss_b327651cc0fdd24615b5e9e4f71f53f032ab712a.600x338.jpg?t=1602795811",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Name:          "Hell Let Loose",
		Image:         "Featured&Recommended5",
		Category:      "Action",
		Price:         200000,
		Discount:      0,
		SideImage:     "Featured&Recommended5",
		Overall:       "3.5",
		Status:        "recommend",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "yes",
		GameplayHours: 150,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})
	db.Create(&SystemRequirement{
		GameID:    5,
		Os:        "Windows 10 64-bit",
		Processor: "Intel Core i7-4770K@3.5GHz or Ryzen 5 1500X@3.5GHz",
		Memory:    "16 GB RAM",
		Graphics:  "Nvidia GeForce GTX 1060 (6 GB) or AMD Radeon RX 580 (8GB)",
		Storage:   "50 GB available space",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Review{
		GameID:         5,
		ReviewDesc:     "**Did not play beta or alpha or any version of the game prior to 01/18/2021**\n\nThis is it.\n\nIt's not fully baked but boy oh boy is it delicious.\n\nI grew up playing Wing Commander 3 in the late 90's and Freelancer in the early 2000's, birthing my love for space combat RPG's.\n\nI've been chasing the dragon ever since.\n\nI remember trying Everspace for the first time and it was not my cup of tea as a rogue-like. But it was so close to what I was searching for.\n\nEverspace 2 is it.\n\nIt's in EA folks. You should set expectations as such, if you want to support the project. It's incredibly playable, feature rich, and gorgeous in its current state.\n\nIt's fun to play, and wow did Rockfish games get a lot right.\n\nI cannot WAIT to see what is in store for the future of development.\n\nI've had one crash so far, and it was expected, as I was trying to get real fancy exiting cruise speed trying to pull off a disgusting amount of ship spins. My GTX 1070 Rog Strix OC edition running everything at EPIC settings did it's best.\n\nOtherwise the game runs very well for day 1 of EA. No bugs experienced or noticed so far, besides one crash due to pushing the limits of my system.\n\nI have no doubt with some time and more love from the Dev team, this game is going to be the start of something great. Perhaps even a new era of space combat rpg's if other studios take notice.\n\nGreat job Rockfish games.\nFly safe everyone. o7",
		ReviewUsername: "ChronCake",
		ReviewAvatar:   "avatar-chroncake.jpg",
		ReviewUpvoted:  10,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&Review{
		GameID:         5,
		ReviewDesc:     "First space RPG I've played where the ship controls like a dream. Systems are simple to understand but seem like they have a considerable amount of depth.\n\nUI is simple and navigating is easy. I'm constantly finding hidden secrets that take me off my main path. So far so good!",
		ReviewUsername: "Daniel",
		ReviewAvatar:   "avatar-daniel.jpg",
		ReviewUpvoted:  5,
		ReviewSentiment: "negative",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&Review{
		GameID:         5,
		ReviewDesc:     "Combat flows nicely, the space is huge and there's always something to explore and the story is solid so far. Extremely well polished for an early access product!",
		ReviewUsername: "KokoPlays",
		ReviewAvatar:   "avatar-kokoplays.jpg",
		ReviewUpvoted:  0,
		ReviewSentiment: "negative",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&VideoGame{
		GameID:      5,
		VideoSource: "https://cdn.akamai.steamstatic.com/steam/apps/256792706/movie.184x123.jpg?t=1594723625",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&VideoGame{
		GameID:      5,
		VideoSource: "https://cdn.akamai.steamstatic.com/steam/apps/256792706/movie.184x123.jpg?t=1594723625",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      5,
		PhotoSource: "https://cdn.akamai.steamstatic.com/steam/apps/686810/ss_f0fa9e773551614a2fe2439f29ad6f74e0937120.600x338.jpg?t=1609948545",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      5,
		PhotoSource: "https://cdn.akamai.steamstatic.com/steam/apps/686810/ss_883861c13f66ad21f92cbb1374cca638aba7ff37.600x338.jpg?t=1609948545",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      5,
		PhotoSource: "https://cdn.akamai.steamstatic.com/steam/apps/686810/ss_62d0a56ddf7c1b60bd85b9a414c29594730275d5.600x338.jpg?t=1609948545",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Name:          "Fallout",
		Image:         "Featured&Recommended6",
		Category:      "Adventure",
		Price:         199000,
		Discount:      0,
		SideImage:     "Featured&Recommended6",
		Overall:       "3.2",
		Status:        "recommend",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		GameplayHours: 100,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})
	db.Create(&SystemRequirement{
		GameID:    6,
		Os:        "Windows 10 64-bit",
		Processor: "Intel Core i7-4770K@3.5GHz or Ryzen 5 1500X@3.5GHz",
		Memory:    "16 GB RAM",
		Graphics:  "Nvidia GeForce GTX 1060 (6 GB) or AMD Radeon RX 580 (8GB)",
		Storage:   "50 GB available space",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Review{
		GameID:         6,
		ReviewDesc:     "**Did not play beta or alpha or any version of the game prior to 01/18/2021**\n\nThis is it.\n\nIt's not fully baked but boy oh boy is it delicious.\n\nI grew up playing Wing Commander 3 in the late 90's and Freelancer in the early 2000's, birthing my love for space combat RPG's.\n\nI've been chasing the dragon ever since.\n\nI remember trying Everspace for the first time and it was not my cup of tea as a rogue-like. But it was so close to what I was searching for.\n\nEverspace 2 is it.\n\nIt's in EA folks. You should set expectations as such, if you want to support the project. It's incredibly playable, feature rich, and gorgeous in its current state.\n\nIt's fun to play, and wow did Rockfish games get a lot right.\n\nI cannot WAIT to see what is in store for the future of development.\n\nI've had one crash so far, and it was expected, as I was trying to get real fancy exiting cruise speed trying to pull off a disgusting amount of ship spins. My GTX 1070 Rog Strix OC edition running everything at EPIC settings did it's best.\n\nOtherwise the game runs very well for day 1 of EA. No bugs experienced or noticed so far, besides one crash due to pushing the limits of my system.\n\nI have no doubt with some time and more love from the Dev team, this game is going to be the start of something great. Perhaps even a new era of space combat rpg's if other studios take notice.\n\nGreat job Rockfish games.\nFly safe everyone. o7",
		ReviewUsername: "ChronCake",
		ReviewAvatar:   "avatar-chroncake.jpg",
		ReviewUpvoted:  10,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&Review{
		GameID:         6,
		ReviewDesc:     "First space RPG I've played where the ship controls like a dream. Systems are simple to understand but seem like they have a considerable amount of depth.\n\nUI is simple and navigating is easy. I'm constantly finding hidden secrets that take me off my main path. So far so good!",
		ReviewUsername: "Daniel",
		ReviewAvatar:   "avatar-daniel.jpg",
		ReviewUpvoted:  5,
		ReviewSentiment: "negative",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&Review{
		GameID:         6,
		ReviewDesc:     "Combat flows nicely, the space is huge and there's always something to explore and the story is solid so far. Extremely well polished for an early access product!",
		ReviewUsername: "KokoPlays",
		ReviewAvatar:   "avatar-kokoplays.jpg",
		ReviewUpvoted:  0,
		ReviewSentiment: "negative",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&VideoGame{
		GameID:      6,
		VideoSource: "https://cdn.akamai.steamstatic.com/steam/apps/256781309/movie480_vp9.webm?t=1587072445",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&VideoGame{
		GameID:      6,
		VideoSource: "https://cdn.akamai.steamstatic.com/steam/apps/256806183/movie480_vp9.webm?t=1606274097",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      6,
		PhotoSource: "https://cdn.akamai.steamstatic.com/steam/apps/1151340/ss_ea796c9634d573a72505dd64537d28a128715e29.600x338.jpg?t=1609880734",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      6,
		PhotoSource: "https://cdn.akamai.steamstatic.com/steam/apps/1151340/ss_a6630b37f2d2a5e40c755af9b1dd7986c98c6299.600x338.jpg?t=1609880734",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      6,
		PhotoSource: "https://cdn.akamai.steamstatic.com/steam/apps/1151340/ss_97a97b0681e648b7fa43bfa6f5fd9e703ecfbba8.600x338.jpg?t=1609880734",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Name:          "PlayerUnknown Battlegrounds",
		Image:         "Featured&Recommended7",
		Category:      "Action",
		Price:         200000,
		Discount:      0,
		SideImage:     "Featured&Recommended7",
		Overall:       "4.0",
		Status:        "recommend",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		GameplayHours: 50,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})
	db.Create(&SystemRequirement{
		GameID:    7,
		Os:        "Windows 10 64-bit",
		Processor: "Intel Core i7-4770K@3.5GHz or Ryzen 5 1500X@3.5GHz",
		Memory:    "16 GB RAM",
		Graphics:  "Nvidia GeForce GTX 1060 (6 GB) or AMD Radeon RX 580 (8GB)",
		Storage:   "50 GB available space",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Review{
		GameID:         7,
		ReviewDesc:     "**Did not play beta or alpha or any version of the game prior to 01/18/2021**\n\nThis is it.\n\nIt's not fully baked but boy oh boy is it delicious.\n\nI grew up playing Wing Commander 3 in the late 90's and Freelancer in the early 2000's, birthing my love for space combat RPG's.\n\nI've been chasing the dragon ever since.\n\nI remember trying Everspace for the first time and it was not my cup of tea as a rogue-like. But it was so close to what I was searching for.\n\nEverspace 2 is it.\n\nIt's in EA folks. You should set expectations as such, if you want to support the project. It's incredibly playable, feature rich, and gorgeous in its current state.\n\nIt's fun to play, and wow did Rockfish games get a lot right.\n\nI cannot WAIT to see what is in store for the future of development.\n\nI've had one crash so far, and it was expected, as I was trying to get real fancy exiting cruise speed trying to pull off a disgusting amount of ship spins. My GTX 1070 Rog Strix OC edition running everything at EPIC settings did it's best.\n\nOtherwise the game runs very well for day 1 of EA. No bugs experienced or noticed so far, besides one crash due to pushing the limits of my system.\n\nI have no doubt with some time and more love from the Dev team, this game is going to be the start of something great. Perhaps even a new era of space combat rpg's if other studios take notice.\n\nGreat job Rockfish games.\nFly safe everyone. o7",
		ReviewUsername: "ChronCake",
		ReviewAvatar:   "avatar-chroncake.jpg",
		ReviewUpvoted:  10,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&Review{
		GameID:         7,
		ReviewDesc:     "First space RPG I've played where the ship controls like a dream. Systems are simple to understand but seem like they have a considerable amount of depth.\n\nUI is simple and navigating is easy. I'm constantly finding hidden secrets that take me off my main path. So far so good!",
		ReviewUsername: "Daniel",
		ReviewAvatar:   "avatar-daniel.jpg",
		ReviewUpvoted:  5,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&Review{
		GameID:         7,
		ReviewDesc:     "Combat flows nicely, the space is huge and there's always something to explore and the story is solid so far. Extremely well polished for an early access product!",
		ReviewUsername: "KokoPlays",
		ReviewAvatar:   "avatar-kokoplays.jpg",
		ReviewUpvoted:  0,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&VideoGame{
		GameID:      7,
		VideoSource: "https://cdn.akamai.steamstatic.com/steam/apps/256814351/movie480_vp9.webm?t=1608093283",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&VideoGame{
		GameID:      7,
		VideoSource: "https://cdn.akamai.steamstatic.com/steam/apps/256805874/movie480_vp9.webm?t=1603301348",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      7,
		PhotoSource: "https://cdn.akamai.steamstatic.com/steam/apps/578080/ss_606cee13e97530720c678513cb1138ef9854d7d5.600x338.jpg?t=1608093288",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      7,
		PhotoSource: "https://cdn.akamai.steamstatic.com/steam/apps/578080/ss_c49417566f70eec8bf0ddbb2956b235d91504a09.600x338.jpg?t=1608093288",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      7,
		PhotoSource: "https://cdn.akamai.steamstatic.com/steam/apps/578080/ss_0985fff929498a15793fc3df766607fb54bf5338.600x338.jpg?t=1608093288",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Name:          "Rust",
		Image:         "Featured&Recommended8",
		Category:      "Horror",
		Price:         200000,
		Discount:      0,
		SideImage:     "Featured&Recommended8",
		Overall:       "4.0",
		Status:        "recommend",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		GameplayHours: 25,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})
	db.Create(&SystemRequirement{
		GameID:    8,
		Os:        "Windows 10 64-bit",
		Processor: "Intel Core i7-4770K@3.5GHz or Ryzen 5 1500X@3.5GHz",
		Memory:    "16 GB RAM",
		Graphics:  "Nvidia GeForce GTX 1060 (6 GB) or AMD Radeon RX 580 (8GB)",
		Storage:   "50 GB available space",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Review{
		GameID:         8,
		ReviewDesc:     "**Did not play beta or alpha or any version of the game prior to 01/18/2021**\n\nThis is it.\n\nIt's not fully baked but boy oh boy is it delicious.\n\nI grew up playing Wing Commander 3 in the late 90's and Freelancer in the early 2000's, birthing my love for space combat RPG's.\n\nI've been chasing the dragon ever since.\n\nI remember trying Everspace for the first time and it was not my cup of tea as a rogue-like. But it was so close to what I was searching for.\n\nEverspace 2 is it.\n\nIt's in EA folks. You should set expectations as such, if you want to support the project. It's incredibly playable, feature rich, and gorgeous in its current state.\n\nIt's fun to play, and wow did Rockfish games get a lot right.\n\nI cannot WAIT to see what is in store for the future of development.\n\nI've had one crash so far, and it was expected, as I was trying to get real fancy exiting cruise speed trying to pull off a disgusting amount of ship spins. My GTX 1070 Rog Strix OC edition running everything at EPIC settings did it's best.\n\nOtherwise the game runs very well for day 1 of EA. No bugs experienced or noticed so far, besides one crash due to pushing the limits of my system.\n\nI have no doubt with some time and more love from the Dev team, this game is going to be the start of something great. Perhaps even a new era of space combat rpg's if other studios take notice.\n\nGreat job Rockfish games.\nFly safe everyone. o7",
		ReviewUsername: "ChronCake",
		ReviewAvatar:   "avatar-chroncake.jpg",
		ReviewUpvoted:  10,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&Review{
		GameID:         8,
		ReviewDesc:     "First space RPG I've played where the ship controls like a dream. Systems are simple to understand but seem like they have a considerable amount of depth.\n\nUI is simple and navigating is easy. I'm constantly finding hidden secrets that take me off my main path. So far so good!",
		ReviewUsername: "Daniel",
		ReviewAvatar:   "avatar-daniel.jpg",
		ReviewUpvoted:  5,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&Review{
		GameID:         8,
		ReviewDesc:     "Combat flows nicely, the space is huge and there's always something to explore and the story is solid so far. Extremely well polished for an early access product!",
		ReviewUsername: "KokoPlays",
		ReviewAvatar:   "avatar-kokoplays.jpg",
		ReviewUpvoted:  0,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&VideoGame{
		GameID:      8,
		VideoSource: "https://cdn.akamai.steamstatic.com/steam/apps/256684736/movie480.webm?t=1518121486",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&VideoGame{
		GameID:      8,
		VideoSource: "https://cdn.akamai.steamstatic.com/steam/apps/256791647/movie480_vp9.webm?t=1593743141",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      8,
		PhotoSource: "https://cdn.akamai.steamstatic.com/steam/apps/252490/ss_58c5b4107fbd90f5a56b3adc97a6b61384ce4f12.600x338.jpg?t=1608404151",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      8,
		PhotoSource: "https://cdn.akamai.steamstatic.com/steam/apps/252490/ss_2a8518810024a5fbf9c714e697a43a1201b5d53e.600x338.jpg?t=1608404151",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      8,
		PhotoSource: "https://cdn.akamai.steamstatic.com/steam/apps/252490/ss_326282c7485e8aff1ebf6750c82622afef098998.600x338.jpg?t=1608404151",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

}

func SeedGameSpecialOffer() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// CREATE RECOMMENDED GAME
	// ID 9-14

	db.Create(&Game{
		Name:          "Need for Speed Heat",
		Image:         "OnSale1",
		Category:      "Sport",
		Price:         1009000,
		Discount:      70,
		SideImage:     "OnSale1",
		Overall:       "5.0",
		Status:        "sale",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})
	db.Create(&SystemRequirement{
		GameID:    9,
		Os:        "Windows 10 64-bit",
		Processor: "Intel Core i7-4770K@3.5GHz or Ryzen 5 1500X@3.5GHz",
		Memory:    "16 GB RAM",
		Graphics:  "Nvidia GeForce GTX 1060 (6 GB) or AMD Radeon RX 580 (8GB)",
		Storage:   "50 GB available space",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Review{
		GameID:         9,
		ReviewDesc:     "**Did not play beta or alpha or any version of the game prior to 01/18/2021**\n\nThis is it.\n\nIt's not fully baked but boy oh boy is it delicious.\n\nI grew up playing Wing Commander 3 in the late 90's and Freelancer in the early 2000's, birthing my love for space combat RPG's.\n\nI've been chasing the dragon ever since.\n\nI remember trying Everspace for the first time and it was not my cup of tea as a rogue-like. But it was so close to what I was searching for.\n\nEverspace 2 is it.\n\nIt's in EA folks. You should set expectations as such, if you want to support the project. It's incredibly playable, feature rich, and gorgeous in its current state.\n\nIt's fun to play, and wow did Rockfish games get a lot right.\n\nI cannot WAIT to see what is in store for the future of development.\n\nI've had one crash so far, and it was expected, as I was trying to get real fancy exiting cruise speed trying to pull off a disgusting amount of ship spins. My GTX 1070 Rog Strix OC edition running everything at EPIC settings did it's best.\n\nOtherwise the game runs very well for day 1 of EA. No bugs experienced or noticed so far, besides one crash due to pushing the limits of my system.\n\nI have no doubt with some time and more love from the Dev team, this game is going to be the start of something great. Perhaps even a new era of space combat rpg's if other studios take notice.\n\nGreat job Rockfish games.\nFly safe everyone. o7",
		ReviewUsername: "ChronCake",
		ReviewAvatar:   "avatar-chroncake.jpg",
		ReviewUpvoted:  10,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&Review{
		GameID:         9,
		ReviewDesc:     "First space RPG I've played where the ship controls like a dream. Systems are simple to understand but seem like they have a considerable amount of depth.\n\nUI is simple and navigating is easy. I'm constantly finding hidden secrets that take me off my main path. So far so good!",
		ReviewUsername: "Daniel",
		ReviewAvatar:   "avatar-daniel.jpg",
		ReviewUpvoted:  5,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&Review{
		GameID:         9,
		ReviewDesc:     "Combat flows nicely, the space is huge and there's always something to explore and the story is solid so far. Extremely well polished for an early access product!",
		ReviewUsername: "KokoPlays",
		ReviewAvatar:   "avatar-kokoplays.jpg",
		ReviewUpvoted:  0,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
	db.Create(&VideoGame{
		GameID:      9,
		VideoSource: "https://cdn.cloudflare.steamstatic.com/steam/apps/256786657/movie480_vp9.webm?t=1591282841",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&VideoGame{
		GameID:      9,
		VideoSource: "https://cdn.cloudflare.steamstatic.com/steam/apps/256789562/movie480_vp9.webm?t=1592508913",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      9,
		PhotoSource: "https://cdn.cloudflare.steamstatic.com/steam/apps/1222680/ss_6994870577a41882c458cd00d852d8092116c81c.600x338.jpg?t=1609386875",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      9,
		PhotoSource: "https://cdn.cloudflare.steamstatic.com/steam/apps/1222680/ss_1f752c037d7cbab2e1658f36d5c76d11e91e4fec.600x338.jpg?t=1609386875",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&PhotoGame{
		GameID:      9,
		PhotoSource: "https://cdn.cloudflare.steamstatic.com/steam/apps/1222680/ss_720840b2cb26c38d0e4ad32085afb5f46b2bb6c6.600x338.jpg?t=1609386875",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Name:          "Farcry 5",
		Image:         "OnSale2",
		Category:      "Adventure",
		Price:         619000,
		Discount:      80,
		SideImage:     "OnSale2",
		Overall:       "4.0",
		Status:        "sale",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Game{
		Name:          "Mafia Trilogy",
		Image:         "OnSale3",
		Category:      "Action",
		Price:         699000,
		Discount:      34,
		SideImage:     "OnSale3",
		Overall:       "4.5",
		Status:        "sale",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Game{
		Name:          "The Outer Worlds",
		Image:         "OnSale4",
		Category:      "Adventure",
		Price:         555000,
		Discount:      50,
		SideImage:     "OnSale4",
		Overall:       "4.0",
		Status:        "sale",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Game{
		Name:          "Crew 2",
		Image:         "OnSale5",
		Category:      "Sport",
		Price:         515000,
		Discount:      80,
		SideImage:     "OnSale5",
		Overall:       "4.0",
		Status:        "sale",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Game{
		Name:          "Warhammer Vermintide",
		Image:         "OnSale6",
		Category:      "Action",
		Price:         139999,
		Discount:      75,
		SideImage:     "OnSale6",
		Overall:       "4.0",
		Status:        "sale",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

}

func SeedGameCommunityRecommended() {

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// CREATE RECOMMENDED GAME
	// ID 15-20

	db.Create(&Game{
		Name:          "The Last of Waifus",
		Image:         "communityrecommended1",
		Category:      "Action",
		Price:         1009000,
		Discount:      0,
		SideImage:     "communityrecommended1",
		Overall:       "4.0",
		Status:        "community recommended",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Review{
		GameID:         15,
		ReviewDesc:     "Gimme my waifus!!! gimme some zombie to shoot for my waifus!!! ill do everything for my waifus and theyll do everything for me!! its amazing i swear... to the WAIFUS AND ZOMBIES!!! cheers.",
		ReviewUsername: "virtyoszt",
		ReviewAvatar:   "avatar-orange.jpg",
		ReviewSentiment: "positive",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})

	db.Create(&Game{
		Name:          "The Old House",
		Image:         "communityrecommended2",
		Category:      "Horror",
		Price:         2009000,
		Discount:      0,
		SideImage:     "communityrecommended2",
		Overall:       "4.5",
		Status:        "community recommended",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Review{
		GameID:         16,
		ReviewDesc:     "I learned about the game recently, but I very-very liked the game. There is nothing complicated in the game, but it is well done, great for killing time. more familiar to me.",
		ReviewUsername: "shuriname",
		ReviewAvatar:   "avatar-orange.jpg",
		ReviewSentiment: "positive",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})

	db.Create(&Game{
		Name:          "Club Girl",
		Image:         "communityrecommended3",
		Category:      "Adventure",
		Price:         89000,
		Discount:      0,
		SideImage:     "communityrecommended3",
		Overall:       "4.0",
		Status:        "community recommended",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Review{
		GameID:         17,
		ReviewDesc:     "A good music game for its price, the only thing I noticed is that some levels are too fast and it is not always possible to press the keys, and most likely the arrows just go at the speed",
		ReviewUsername: "Den Haag",
		ReviewAvatar:   "avatar-orange.jpg",
		ReviewSentiment: "positive",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})

	db.Create(&Game{
		Name:          "Hazardous Space",
		Image:         "communityrecommended4",
		Category:      "Adventure",
		Price:         119000,
		Discount:      0,
		SideImage:     "communityrecommended4",
		Overall:       "4.5",
		Status:        "community recommended",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Review{
		GameID:         18,
		ReviewDesc:     "Pretty fun rogue like survival game with a bit of story to help it along. Had fun playing through, having different monster encounters and finding secret notes to add little details.",
		ReviewUsername: "Jonny Killer",
		ReviewAvatar:   "avatar-orange.jpg",
		ReviewSentiment: "positive",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})

	db.Create(&Game{
		Name:          "The Iaundry",
		Image:         "communityrecommended5",
		Category:      "Horror",
		Price:         39000,
		Discount:      0,
		SideImage:     "communityrecommended5",
		Overall:       "4.5",
		Status:        "community recommended",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Review{
		GameID:         19,
		ReviewDesc:     "The most common horror indie craft with a hidden object, the graphics are simple and will allow you to run the game on the weakest computer.",
		ReviewUsername: "Afraid",
		ReviewAvatar:   "avatar-orange.jpg",
		ReviewSentiment: "positive",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})

	db.Create(&Game{
		Name:          "Four Ways",
		Image:         "communityrecommended6",
		Category:      "Adventure",
		Price:         29000,
		Discount:      0,
		SideImage:     "communityrecommended6",
		Overall:       "4.0",
		Status:        "community recommended",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Review{
		GameID:         20,
		ReviewDesc:     "Well done. Clean, snappy interface, pretty much what you'd expect from a tetris, with many options and variations on four side at the same time. This game made Tetris way more fun for me!",
		ReviewUsername: "jess no limit",
		ReviewAvatar:   "avatar-orange.jpg",
		ReviewSentiment: "positive",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})

}

func SeedGameNewAndTrending() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// CREATE RECOMMENDED GAME
	// ID 21-25

	db.Create(&Game{
		Name:          "Knights College",
		Image:         "trending1",
		Category:      "Dating Sim",
		Price:         120000,
		Discount:      0,
		SideImage:     "side-image1",
		Overall:       "4.5",
		Status:        "new trending",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Game{
		Name:          "Farm Manager 2021 : Prologue",
		Image:         "trending2",
		Category:      "Strategy",
		Price:         154000,
		Discount:      0,
		SideImage:     "side-image2",
		Overall:       "4.5",
		Status:        "new trending",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Game{
		Name:          "Cloud Climber",
		Image:         "trending3",
		Category:      "Adventure",
		Price:         174000,
		Discount:      0,
		SideImage:     "side-image3",
		Overall:       "5.0",
		Status:        "new trending",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Game{
		Name:          "Yaga",
		Image:         "trending4",
		Category:      "RPG",
		Price:         200000,
		Discount:      0,
		SideImage:     "side-image4",
		Overall:       "5.0",
		Status:        "new trending",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Game{
		Name:          "Heroes of the Three Kingdoms 8",
		Image:         "trending5",
		Category:      "Strategy",
		Price:         170000,
		Discount:      0,
		SideImage:     "side-image5",
		Overall:       "4.5",
		Status:        "new trending",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

}

func SeedGameTopSellers() {

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// CREATE RECOMMENDED GAME
	// ID 26-30

	db.Create(&Game{
		Name:          "Rust",
		Image:         "top-seller1",
		Category:      "Multiplayer",
		Price:         169999,
		Discount:      0,
		SideImage:     "side-image11",
		Overall:       "4.5",
		Status:        "top seller",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Game{
		Name:          "CyberPunk",
		Image:         "top-seller2",
		Category:      "Futuristic",
		Price:         699999,
		Discount:      0,
		SideImage:     "side-image12",
		Overall:       "4.7",
		Status:        "top seller",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "yes",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Game{
		Name:          "Red Dead Redemption",
		Image:         "top-seller3",
		Category:      "Western",
		Price:         428800,
		Discount:      0,
		SideImage:     "side-image13",
		Overall:       "4.5",
		Status:        "top seller",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "yes",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Game{
		Name:          "Grand Theft Auto 5",
		Image:         "top-seller4",
		Category:      "Multiplayer",
		Price:         290000,
		Discount:      0,
		SideImage:     "side-image14",
		Overall:       "4.5",
		Status:        "top seller",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Game{
		Name:          "Stardew Valley",
		Image:         "top-seller5",
		Category:      "Farming Sim",
		Price:         115999,
		Discount:      0,
		SideImage:     "side-image15",
		Overall:       "4.0",
		Status:        "top seller",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

}

func SeedGameSpecials() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// CREATE RECOMMENDED GAME
	// ID 31-35

	db.Create(&Game{
		Name:          "UBOAT",
		Image:         "special1",
		Category:      "Strategy",
		Price:         180000,
		Discount:      60,
		SideImage:     "side-image21",
		Overall:       "4.5",
		Status:        "special",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Game{
		Name:          "Mafia : Definitive Edition",
		Image:         "special2",
		Category:      "Action",
		Price:         210000,
		Discount:      70,
		SideImage:     "side-image22",
		Overall:       "4.0",
		Status:        "special",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Game{
		Name:          "The Outer Worlds",
		Image:         "special3",
		Category:      "RPG",
		Price:         130000,
		Discount:      65,
		SideImage:     "side-image23",
		Overall:       "3.5",
		Status:        "special",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Game{
		Name:          "Need for Speed : Playback",
		Image:         "special4",
		Category:      "Sport",
		Price:         150000,
		Discount:      75,
		SideImage:     "side-image24",
		Overall:       "4.0",
		Status:        "special",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Game{
		Name:          "Cooking Simulator",
		Image:         "special5",
		Category:      "Adventure",
		Price:         240000,
		Discount:      60,
		SideImage:     "side-image25",
		Overall:       "4.0",
		Status:        "special",
		Description:   "",
		Developer:     "",
		Publisher:     "",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})
}

// REVIEW COMMENT

func InsertReviewCommentByReviewId(token string, reviewid int, comment string)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))

	db.Create(&ReviewComment{
		UserId:    userid,
		ReviewId:  reviewid,
		Comment:   comment,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return true, nil
}

func GetAllReviewCommentByReviewId(reviewid int)([]ReviewComment, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var reviewComment []ReviewComment
	db.Where("review_id = ?", reviewid).Find(&reviewComment)

	return reviewComment, nil
}

func GetReviewCommentByReviewIdOffsetLimit(reviewid int, offset int, limit int)([]ReviewComment, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var reviewComment []ReviewComment
	db.Where("review_id = ?", reviewid).Offset(offset).Limit(limit).Find(&reviewComment)

	return reviewComment, nil
}


// GET ALL GAMES
func GetGames() ([]Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var games []Game
	db.Find(&games)

	return games, nil
}

func GetGamesOffsetLimit(offset int, limit int) ([]Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var games []Game
	db.Offset(offset).Limit(limit).Find(&games)

	return games, nil
}

// GET GAME BY ID
func GetGame(id int) (Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var game Game
	db.First(&game, id)

	return game, nil
}

// GET GAME BY STATUS
func GetGamesByStatus(status string) ([]Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var games []Game
	db.Where("status = ?", status).Find(&games)

	return games, nil
}

// Get Game Featured and Recommended
func GetGamesByFeaturedAndRecommended(limit int)([]Game, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var games []Game
	db.Order("gameplay_hours desc").Limit(limit).Find(&games)

	return games, nil
}

// Get Game Special Offers
func GetGamesSpecialOffers(limit int)([]Game, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var games []Game
	db.Where("discount > 0").Limit(limit).Find(&games)

	return games, nil
}

// Get Game New and Trending
func GetGamesNewAndTrending(limit int)([]Game, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var games []Game
	db.Order("created_at desc").Limit(limit).Find(&games)

	return games, nil
}

// GET GAME BY STATUS AND LIMIT
func GetGamesByStatusLimit(status string, limit int) ([]Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var games []Game
	db.Where("status = ?", status).Limit(limit).Find(&games)

	return games, nil
}

func GetGamesSpecialsSpecificDiscount(status string, discount int, limit int) ([]Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var games []Game
	db.Where("status = ?", status).Where("discount > ?", discount).Limit(limit).Find(&games)

	return games, nil
}

func GetGamesByNameLimit(name string, limit int) ([]Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var games []Game
	db.Where("name like ? or name like lower(?)", "%"+name+"%", "%"+name+"%").Limit(limit).Find(&games)

	return games, nil
}

func GetGamesByName(name string) ([]Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var games []Game
	db.Where("name like ? or name like lower(?)", "%"+name+"%", "%"+name+"%").Find(&games)

	return games, nil
}

func GetGamesByFilter(offset int, limit int, name string, price int, category string, genre string) ([]Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var games []Game

	if category == "" && genre == "" {
		db.Where("(name like ? or name like lower(?)) and price <= ?", "%"+name+"%", "%"+name+"%", price).Offset(offset).Limit(limit).Find(&games)
	} else if category == "" && genre != "" {
		db.Where("(name like ? or name like lower(?)) and price <= ? and category = ?", "%"+name+"%", "%"+name+"%", price, genre).Offset(offset).Limit(limit).Find(&games)
	} else if genre == "" && category != "" {
		db.Where("(name like ? or name like lower(?)) and price <= ? and status = ?", "%"+name+"%", "%"+name+"%", price, category).Offset(offset).Limit(limit).Find(&games)
	} else {
		db.Where("(name like ? or name like lower(?)) and price <= ? and category = ? and status = ?", "%"+name+"%", "%"+name+"%", price, genre, category).Offset(offset).Limit(limit).Find(&games)
	}

	//db.Where("name like ? and price <= ? and category = ? and status = ?", "%"+name+"%", price, genre, category).Offset(offset).Limit(limit).Find(&games)

	return games, nil
}

func GetReviewUpvoted(gameId int, limit int) ([]Review, error) {
	db, _ := database.Connect()
	defer db.Close()

	var review []Review
	db.Where("game_id = ? and review_upvoted > 0 and DATE_PART( 'day', NOW()) - DATE_PART( 'day', created_at) <= 31", gameId).Order("review_upvoted desc").Limit(limit).Find(&review)

	return review, nil
}

func GetReviewRecently(gameId int, limit int) ([]Review, error) {
	db, _ := database.Connect()
	defer db.Close()

	var review []Review
	db.Where("game_id = ?", gameId).Order("created_at desc").Limit(limit).Find(&review)

	return review, nil
}

func GetAllReview(limit int)([]Review, error){
	db, _ := database.Connect()
	defer db.Close()

	var review []Review
	db.Order("id").Limit(limit).Find(&review)

	return review, nil
}

func GetReviewById(id int)(Review, error){
	db, _ := database.Connect()
	defer db.Close()

	var review Review
	db.Where("id = ?", id).First(&review)

	return review, nil
}

// GET GAME BY STATUS, LIMIT, AND ID
//func GetGamesByStatusId(status string, id int) (Game, error) {
//	db, err := database.Connect()
//	if err != nil {
//		panic(err)
//	}
//	defer db.Close()
//
//	var games Game
//	db.Where("status = ?", status).Where("id = ?", id).First(&games)
//
//	return games, nil
//}

func UpdateReviewUpvote(upvote int, id uint) (bool, error) {

	db, _ := database.Connect()
	defer db.Close()

	db.Model(&Review{ID: id}).UpdateColumn("review_upvoted", gorm.Expr("review_upvoted + ?", upvote))

	return true, nil
}

func UpdateReviewDownvote(downvote int, id uint) (bool, error) {

	db, _ := database.Connect()
	defer db.Close()

	db.Model(&Review{ID: id}).UpdateColumn("review_downvoted", gorm.Expr("review_downvoted + ?", downvote))

	return true, nil
}

func InsertReview(token string, reviewDesc string, gameid int) (Review, error) {
	db, _ := database.Connect()
	defer db.Close()

	userId := getIdFromJWTToken(token)

	var user User
	db.Where("id = ?", userId).First(&user)

	db.Create(&Review{
		GameID:         uint(gameid),
		ReviewDesc:     reviewDesc,
		ReviewUsername: user.Username,
		ReviewAvatar:   user.Avatar,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})

	var review Review
	db.Last(&review)

	return review, nil
}





// ADMIN

// 1
func InsertGame(encodedBase64Img string, filename string, title string, desc string, price int, tags string) (bool, error) {
	decoded, _ := base64.StdEncoding.DecodeString(encodedBase64Img)

	writeFilename := "C:\\Vincent\\TPA Web\\Project\\src\\github.com\\Vincent-Benedict\\Angular\\TPA-Web\\src\\assets\\Search Picture" + "/" + filename
	errs := ioutil.WriteFile(writeFilename, decoded, os.ModePerm)

	db, _ := database.Connect()
	defer db.Close()
	
	db.Create(&Game{
		Name:          title,
		Image:         "dummy",
		Category:      tags,
		SideImage:     "dummy",
		Overall:       "5.0",
		Status:        "none",
		Price:         int64(price),
		Description:   desc,
		Developer:     "none",
		Publisher:     "none",
		Inappropriate: "no",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	if errs != nil {
		fmt.Print("Error Save File: ")
		fmt.Println(errs)
	}

	return true, nil
}

func DeleteGame(id int)(bool, error){
	db, _ := database.Connect()
	defer db.Close()

	db.Where("id = ?", id).Delete(Game{})

	return true, nil
}

func UpdateGame(id int, encodedBase64Img string, filename string, title string, desc string, price int, tags string)(bool, error){
	decoded, _ := base64.StdEncoding.DecodeString(encodedBase64Img)

	writeFilename := "C:\\Vincent\\TPA Web\\Project\\src\\github.com\\Vincent-Benedict\\Angular\\TPA-Web\\src\\assets\\Search Picture" + "/" + filename
	errs := ioutil.WriteFile(writeFilename, decoded, os.ModePerm)
	writeFilename2 := "C:\\Vincent\\TPA Web\\Project\\src\\github.com\\Vincent-Benedict\\Angular\\TPA-Web\\src\\assets" + "/" + filename
	errs2 := ioutil.WriteFile(writeFilename2, decoded, os.ModePerm)

	db, _ := database.Connect()
	defer db.Close()

	db.Model(&Game{}).Where("id = ?", id).Updates(Game{Name: title, Description: desc, Price: int64(price), Category: tags, Image: "dummy1", SideImage: "dummy1"})

	if errs != nil || errs2 != nil{
		fmt.Print("Error Save File: ")
		fmt.Println(errs)
	}

	return true, nil
}

// 2
func GetAllPromoGames(auth string)([]Game, error){
	db, _ := database.Connect()
	defer db.Close()

	// AUTH STRING
	if auth != "abc"{
		return nil, nil
	}

	var games []Game
	db.Where("discount > 0").Find(&games)

	return games, nil
}

func GetPromoGameOffsetLimit(auth string, offset int, limit int)([]Game, error){
	db, _ := database.Connect()
	defer db.Close()

	// AUTH STRING
	if auth != "abc"{
		return nil, nil
	}

	var games []Game
	db.Where("discount > 0").Offset(offset).Limit(limit).Find(&games)

	return games, nil
}

func InsertUpdateDeletePromo(id int, promo int)(bool, error){
	db, _ := database.Connect()
	defer db.Close()

	db.Model(&Game{}).Where("id = ?", id).Update("discount", promo)

	return true, nil
}