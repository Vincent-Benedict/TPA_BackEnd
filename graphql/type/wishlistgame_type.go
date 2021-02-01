package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var wishlistGameType *graphql.Object

func GetWishlistGameType() *graphql.Object{

	if wishlistGameType == nil{
		wishlistGameType = graphql.NewObject(graphql.ObjectConfig{
			Name:"wishlistGameType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"userid": &graphql.Field{
					Type: graphql.Int,
				},
				"gameid": &graphql.Field{
					Type: graphql.Int,
				},
				"game": &graphql.Field{
					Type: graphql.NewList(GetGameType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						wishlistGame := params.Source.(models.WishlistGame)

						var game []models.Game
						db.Where("id = ?", wishlistGame.GameID).Find(&game)

						return game, nil
					},
				},
			},
		})
	}

	return wishlistGameType
}
