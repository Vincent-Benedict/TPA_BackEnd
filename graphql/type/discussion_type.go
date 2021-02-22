package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var discussionType *graphql.Object

func GetDiscussionType() *graphql.Object{

	if discussionType == nil{
		discussionType = graphql.NewObject(graphql.ObjectConfig{
			Name:"discussionType",
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
						userid := params.Source.(models.Discussion)

						var user []models.User
						db.Where("id = ?", userid.UserId).First(&user)

						return user, nil
					},
				},
				"gameid": &graphql.Field{
					Type: graphql.Int,
				},
				"game": &graphql.Field{
					Type: graphql.NewList(GetGameType())  ,
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						game := params.Source.(models.Discussion)

						var games []models.Game
						db.Where("id = ?", game.GameId).First(&games)

						return games, nil
					},
				},
				"discussiontitle": &graphql.Field{
					Type: graphql.String,
				},
				"discussioncontent": &graphql.Field{
					Type: graphql.String,
				},

			},
		})
	}

	return discussionType
}

