package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var communityImageAndVideoType *graphql.Object

func GetCommunityImageAndVideoType() *graphql.Object{

	if communityImageAndVideoType == nil{
		communityImageAndVideoType = graphql.NewObject(graphql.ObjectConfig{
			Name:"communityImageAndVideoType",
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
						userid := params.Source.(models.CommunityImageAndVideo)

						var user []models.User
						db.Where("id = ?", userid.UserId).Find(&user)

						return user, nil
					},
				},
				"source": &graphql.Field{
					Type: graphql.String,
				},
				"poster": &graphql.Field{
					Type: graphql.String,
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
				"liked": &graphql.Field{
					Type: graphql.Int,
				},
				"disliked": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}

	return communityImageAndVideoType
}
