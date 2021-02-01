package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var userType *graphql.Object

func GetUserType() *graphql.Object{

	if userType == nil{
		userType = graphql.NewObject(graphql.ObjectConfig{
			Name:"userType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"username": &graphql.Field{
					Type: graphql.String,
				},
				"password": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"country": &graphql.Field{
					Type: graphql.String,
				},
				"balance": &graphql.Field{
					Type: graphql.Int,
				},
				"avatar": &graphql.Field{
					Type: graphql.String,
				},
				"level": &graphql.Field{
					Type: graphql.Int,
				},
				"summary": &graphql.Field{
					Type: graphql.String,
				},
				"featuredbadge": &graphql.Field{
					Type: graphql.String,
				},
				"status": &graphql.Field{
					Type: graphql.String,
				},
				"wishlist": &graphql.Field{
					Type: graphql.NewList(GetWishlistGameType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						user := params.Source.(models.User)

						var wishlistGame []models.WishlistGame
						db.Where("user_id = ?", user.ID).Find(&wishlistGame)

						return wishlistGame, nil
					},
				},
			},
		})
	}

	return userType
}

