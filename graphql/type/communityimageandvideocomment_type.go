package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var communityImageAndVideoCommentType *graphql.Object

func GetCommunityImageAndVideoCommentType() *graphql.Object{

	if communityImageAndVideoCommentType == nil{
		communityImageAndVideoCommentType = graphql.NewObject(graphql.ObjectConfig{
			Name:"communityImageAndVideoCommentType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"userid": &graphql.Field{
					Type: graphql.Int,
				},
				"user": &graphql.Field{
					Type: graphql.NewList(GetUserType()) ,
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						userid := params.Source.(models.CommunityImageAndVideoComment)

						var user []models.User
						db.Where("id = ?", userid.UserId).Find(&user)

						return user, nil
					},
				},
				"contentid": &graphql.Field{
					Type: graphql.Int,
				},
				"comment": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return communityImageAndVideoCommentType
}

