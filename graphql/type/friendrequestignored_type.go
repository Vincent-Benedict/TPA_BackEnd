package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var friendRequestIgnoredType *graphql.Object

func GetFriendRequestIgnoredType() *graphql.Object{

	if friendRequestIgnoredType == nil{
		friendRequestIgnoredType = graphql.NewObject(graphql.ObjectConfig{
			Name:"friendRequestIgnoredType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"userid": &graphql.Field{
					Type: graphql.Int,
				},
				"user":&graphql.Field{
					Type: graphql.NewList(GetUserType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						friend := params.Source.(models.FriendRequestIgnored)

						var user []models.User
						db.Where("id = ?", friend.FriendId).Find(&user)

						return user, nil
					},
				},
				"friendid": &graphql.Field{
					Type: graphql.Int,
				},
				"friend": &graphql.Field{
					Type: graphql.NewList(GetUserType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						friend := params.Source.(models.FriendRequestIgnored)

						var user []models.User
						db.Where("id = ?", friend.UserId).Find(&user)

						return user, nil
					},
				},
			},
		})
	}

	return friendRequestIgnoredType
}

