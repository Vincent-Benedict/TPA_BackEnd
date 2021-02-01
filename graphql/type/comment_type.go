package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var commentType *graphql.Object

func GetCommentType() *graphql.Object{

	if commentType == nil{
		commentType = graphql.NewObject(graphql.ObjectConfig{
			Name:"commentType",
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
				"commentdesc": &graphql.Field{
					Type: graphql.String,
				},
				"person": &graphql.Field{
					Type: graphql.NewList(GetUserType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						friend := params.Source.(models.Comment)

						var user []models.User
						db.Where("id = ?", friend.FriendId).Find(&user)

						return user, nil
					},
				},
			},
		})
	}

	return commentType
}
