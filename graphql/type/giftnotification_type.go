package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var giftNotificationType *graphql.Object

func GetGiftNotificationType() *graphql.Object{

	if giftNotificationType == nil{
		giftNotificationType = graphql.NewObject(graphql.ObjectConfig{
			Name:"giftNotificationType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"userid": &graphql.Field{
					Type: graphql.Int,
				},
				"friendid": &graphql.Field{
					Type: graphql.Int,
				},
				"friend": &graphql.Field{
					Type: graphql.NewList(GetUserType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						friend := params.Source.(models.GiftNotification)

						var user []models.User
						db.Where("id = ?", friend.FriendId).Find(&user)

						return user, nil
					},
				},
			},
		})
	}

	return giftNotificationType
}

