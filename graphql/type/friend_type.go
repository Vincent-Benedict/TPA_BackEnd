package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var friendType *graphql.Object

func GetFriendType() *graphql.Object{

	if friendType == nil{
		friendType = graphql.NewObject(graphql.ObjectConfig{
			Name:"friendType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"userid": &graphql.Field{
					Type: graphql.Int,
				},
				"friendid": &graphql.Field{
					Type: graphql.String,
				},
				"friend": &graphql.Field{
					Type: graphql.NewList(GetUserType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					db, _:= database.Connect()
					defer db.Close()
					friend := params.Source.(models.Friend)

					var user []models.User
					db.Where("id = ?", friend.FriendId).Find(&user)

					return user, nil
					},
				},
			},
		})
	}

	return friendType
}