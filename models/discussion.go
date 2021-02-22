package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type Discussion struct {
	ID                uint `gorm:"primary_key"`
	UserId            int
	GameId            int
	DiscussionTitle   string
	DiscussionContent string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time `sql:"index"`
}

type DiscussionComment struct {
	ID                uint `gorm:"primary_key"`
	UserId            int
	DiscussionId      int
	Comment           string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time `sql:"index"`
}

func init() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DropTableIfExists(&Discussion{}, &DiscussionComment{})
	db.AutoMigrate(&Discussion{}, &DiscussionComment{})

	SeedDiscussion()
	SeedDiscussionComment()
}

func SeedDiscussion(){

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//db.Create(&Discussion{
	//	UserId:            1,
	//	GameId:            6,
	//	DiscussionTitle:   "Enim non voluptatibus in et ut.",
	//	DiscussionContent: "Rerum laborum vero eum ullam corporis ipsam est expedita quae ea quia qui non dolores ipsa impedit eos et officiis.",
	//	CreatedAt:         time.Time{},
	//	UpdatedAt:         time.Time{},
	//	DeletedAt:         nil,
	//})
	//
	//db.Create(&Discussion{
	//	UserId:            2,
	//	GameId:            6,
	//	DiscussionTitle:   "Enim non voluptatibus in et ut.",
	//	DiscussionContent: "Rerum laborum vero eum ullam corporis ipsam est expedita quae ea quia qui non dolores ipsa impedit eos et officiis.",
	//	CreatedAt:         time.Time{},
	//	UpdatedAt:         time.Time{},
	//	DeletedAt:         nil,
	//})

	db.Create(&Discussion{
		UserId:            1,
		GameId:            1,
		DiscussionTitle:   "Enim non voluptatibus in et ut.",
		DiscussionContent: "Rerum laborum vero eum ullam corporis ipsam est expedita quae ea quia qui non dolores ipsa impedit eos et officiis.",
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
		DeletedAt:         nil,
	})
	db.Create(&Discussion{
		UserId:            2,
		GameId:            1,
		DiscussionTitle:   "Cupiditate et ipsam est asperiores.",
		DiscussionContent: "Dolorem eveniet ratione alias dolorum et quo aperiam adipisci.",
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
		DeletedAt:         nil,
	})
	db.Create(&Discussion{
		UserId:            3,
		GameId:            1,
		DiscussionTitle:   "Consequatur voluptas et dignissimos.",
		DiscussionContent: "Odit eos deleniti dolores tenetur consequatur incidunt occaecati amet quia.",
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
		DeletedAt:         nil,
	})
	db.Create(&Discussion{
		UserId:            4,
		GameId:            1,
		DiscussionTitle:   "Facere et facilis id saepe.",
		DiscussionContent: "Itaque rem vitae illo saepe accusamus et aspernatur officia consequatur.",
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
		DeletedAt:         nil,
	})


	db.Create(&Discussion{
		UserId:            4,
		GameId:            2,
		DiscussionTitle:   "Facere et facilis id saepe.",
		DiscussionContent: "Itaque rem vitae illo saepe accusamus et aspernatur officia consequatur.",
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
		DeletedAt:         nil,
	})
	db.Create(&Discussion{
		UserId:            3,
		GameId:            2,
		DiscussionTitle:   "Est et aperiam et.",
		DiscussionContent: "Vero maxime rerum non sapiente occaecati eum laboriosam aut.",
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
		DeletedAt:         nil,
	})
	db.Create(&Discussion{
		UserId:            2,
		GameId:            2,
		DiscussionTitle:   "Sit rem consectetur beatae modi.",
		DiscussionContent: "Soluta iste itaque vel doloribus qui dolores consectetur fuga sed dicta voluptas iusto possimus et suscipit aut ad quia.",
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
		DeletedAt:         nil,
	})
	db.Create(&Discussion{
		UserId:            1,
		GameId:            2,
		DiscussionTitle:   "Laborum est quasi et.",
		DiscussionContent: "Quaerat vel omnis aut rem sit explicabo iste cum unde saepe animi dignissimos numquam deleniti sequi consequatur sed soluta.",
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
		DeletedAt:         nil,
	})


	db.Create(&Discussion{
		UserId:            2,
		GameId:            3,
		DiscussionTitle:   "Est et aperiam et.",
		DiscussionContent: "Vero maxime rerum non sapiente occaecati eum laboriosam aut.",
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
		DeletedAt:         nil,
	})

	db.Create(&Discussion{
		UserId:            3,
		GameId:            4,
		DiscussionTitle:   "Facere et facilis id saepe.",
		DiscussionContent: "Quaerat vel omnis aut rem sit explicabo iste cum unde saepe animi dignissimos numquam deleniti sequi consequatur sed soluta.",
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
		DeletedAt:         nil,
	})

}

func SeedDiscussionComment(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&DiscussionComment{
		UserId:       1,
		DiscussionId: 4,
		Comment:      "Outstandingly thought out! This is new school.",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       2,
		DiscussionId: 4,
		Comment:      "Nice work you have here.",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       3,
		DiscussionId: 3,
		Comment:      "Avatar, type, camera angle, shot – clean :)",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       4,
		DiscussionId: 3,
		Comment:      "I think I'm crying. It's that magnificent.",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       1,
		DiscussionId: 2,
		Comment:      "Such design, many playfulness, so admirable",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       2,
		DiscussionId: 2,
		Comment:      "Nice use of red in this spaces =)",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})

	db.Create(&DiscussionComment{
		UserId:       3,
		DiscussionId: 1,
		Comment:      "Very classic!!",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       4,
		DiscussionId: 1,
		Comment:      "Such nice.",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       1,
		DiscussionId: 8,
		Comment:      "Alluring idea, friend.",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       2,
		DiscussionId: 8,
		Comment:      "This is neat work!!",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       3,
		DiscussionId: 7,
		Comment:      "Let me take a nap... great job, anyway.",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       4,
		DiscussionId: 7,
		Comment:      "My 24 year old mum rates this shot very graceful mate",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})

	db.Create(&DiscussionComment{
		UserId:       1,
		DiscussionId: 6,
		Comment:      "It's nice not just radiant!",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       2,
		DiscussionId: 6,
		Comment:      "This icons has navigated right into my heart.",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       3,
		DiscussionId: 5,
		Comment:      "Mmh wondering if this comment will hit the generateor as well...",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       4,
		DiscussionId: 5,
		Comment:      "Looks nice and radiant :)",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       1,
		DiscussionId: 4,
		Comment:      "Wow love it!",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       2,
		DiscussionId: 3,
		Comment:      "Love your shot.",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})


	db.Create(&DiscussionComment{
		UserId:       3,
		DiscussionId: 0,
		Comment:      "",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       4,
		DiscussionId: 0,
		Comment:      "",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       1,
		DiscussionId: 0,
		Comment:      "",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       2,
		DiscussionId: 0,
		Comment:      "",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       3,
		DiscussionId: 0,
		Comment:      "",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&DiscussionComment{
		UserId:       4,
		DiscussionId: 0,
		Comment:      "",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})

}

func InsertDiscussion(token string, gameid int, discussiontitle string, discussioncontent string)(bool, error){

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))

	db.Create(&Discussion{
		UserId:            userid,
		GameId:            gameid,
		DiscussionTitle:   discussiontitle,
		DiscussionContent: discussioncontent,
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
		DeletedAt:         nil,
	})

	return true, nil
}

func InsertDiscussionComment(token string, discussionid int, comment string)(DiscussionComment, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userid := int(getIdFromJWTToken(token))

	db.Create(&DiscussionComment{
		UserId:       userid,
		DiscussionId: discussionid,
		Comment:      comment,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})

	var discussionComment DiscussionComment
	db.Last(&discussionComment)

	return discussionComment, nil
}

func GetDiscussionCommentByDiscussionId(discussionid int)([]DiscussionComment, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var discussionComment []DiscussionComment
	db.Where("discussion_id = ?", discussionid).Find(&discussionComment)

	return discussionComment, nil
}

func GetAllDiscussion()([]Discussion, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var discussions []Discussion
	db.Order("created_at desc").Find(&discussions)

	return discussions, nil
}

func GetDiscussionsWithGameName(gameName string)([]Discussion, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var discussions []Discussion
	db.Joins("JOIN games ON games.id = discussions.game_id").Where("games.name like ? or games.name like lower(?)", "%"+gameName+"%", "%"+gameName+"%").Order("created_at desc").Find(&discussions)

	return discussions, nil
}

func GetDiscussionWithId(id int)(Discussion, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var discussions Discussion
	db.Where("id = ?", id).First(&discussions)

	return discussions, nil
}
