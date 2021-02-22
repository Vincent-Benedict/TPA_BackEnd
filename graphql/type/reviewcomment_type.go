package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var reviewCommentType *graphql.Object

func GetReviewCommentType() *graphql.Object{

	if reviewCommentType == nil{
		reviewCommentType = graphql.NewObject(graphql.ObjectConfig{
			Name:"reviewCommentType",
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
						reviewComment := params.Source.(models.ReviewComment)

						var user []models.User
						db.Where("id = ?", reviewComment.UserId).Find(&user)

						return user, nil
					},
				},
				"reviewid": &graphql.Field{
					Type: graphql.String,
				},
				"comment": &graphql.Field{
					Type: graphql.String,
				},

			},
		})
	}

	return reviewCommentType
}

