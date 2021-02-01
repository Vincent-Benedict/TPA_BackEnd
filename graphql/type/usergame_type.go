package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var userGameType *graphql.Object

func GetUserGameType() *graphql.Object{

	if userGameType == nil{
		userGameType = graphql.NewObject(graphql.ObjectConfig{
			Name:"userGameType",
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
						user := params.Source.(models.UserGame)

						var game []models.Game
						db.Where("id = ?", user.GameId).Find(&game)

						return game, nil
					},
				},
			},
		})
	}

	return userGameType
}
