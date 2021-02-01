package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var gameItemType *graphql.Object

func GetGameItemType() *graphql.Object{

	if gameItemType == nil{
		gameItemType = graphql.NewObject(graphql.ObjectConfig{
			Name:"gameItemType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"gameid": &graphql.Field{
					Type: graphql.Int,
				},
				"itemname": &graphql.Field{
					Type: graphql.String,
				},
				"itemtransaction": &graphql.Field{
					Type: graphql.Int,
				},
				"itemprice": &graphql.Field{
					Type: graphql.Int,
				},
				"itemimage": &graphql.Field{
					Type: graphql.String,
				},
				"itemdescription": &graphql.Field{
					Type: graphql.String,
				},
				"game": &graphql.Field{
					Type: graphql.NewList(GetGameType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						gameP := params.Source.(models.GameItem)

						var game []models.Game
						db.Where("id = ?", gameP.GameId).Find(&game)

						return game, nil
					},
				},
			},
		})
	}

	return gameItemType
}