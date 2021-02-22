package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var chatNotificationType *graphql.Object

func GetChatNotificationType() *graphql.Object{

	if chatNotificationType == nil{
		chatNotificationType = graphql.NewObject(graphql.ObjectConfig{
			Name:"chatNotificationType",
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
						friend := params.Source.(models.ChatNotification)

						var user []models.User
						db.Where("id = ?", friend.FriendId).Find(&user)

						return user, nil
					},
				},
			},
		})
	}

	return chatNotificationType
}
