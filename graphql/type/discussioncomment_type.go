package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var discussionCommentType *graphql.Object

func GetDiscussionCommentType() *graphql.Object{

	if discussionCommentType == nil{
		discussionCommentType = graphql.NewObject(graphql.ObjectConfig{
			Name:"discussionCommentType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"userid": &graphql.Field{
					Type: graphql.Int,
				},
				"user": &graphql.Field{
					Type: graphql.NewList(GetUserType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						userid := params.Source.(models.DiscussionComment)

						var user []models.User
						db.Where("id = ?", userid.UserId).First(&user)

						return user, nil
					},
				},
				"discussionid": &graphql.Field{
					Type: graphql.Int,
				},
				"comment": &graphql.Field{
					Type: graphql.String,
				},

			},
		})
	}

	return discussionCommentType
}


