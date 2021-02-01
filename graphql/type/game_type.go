package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var gameType *graphql.Object

func GetGameType() *graphql.Object{

	if gameType == nil{
		gameType = graphql.NewObject(graphql.ObjectConfig{
			Name:"gameType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"image": &graphql.Field{
					Type: graphql.String,
				},
				"category": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
				"discount": &graphql.Field{
					Type: graphql.Int,
				},
				"sideimage": &graphql.Field{
					Type: graphql.String,
				},
				"overall": &graphql.Field{
					Type: graphql.String,
				},
				"status": &graphql.Field{
					Type: graphql.String,
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
				"developer": &graphql.Field{
					Type: graphql.String,
				},
				"publisher": &graphql.Field{
					Type: graphql.String,
				},
				"inappropriate": &graphql.Field{
					Type: graphql.String,
				},
				"review": &graphql.Field{
					Type: graphql.NewList(GetReviewType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						gameP := params.Source.(models.Game)

						var review []models.Review
						db.Where("game_id = ?", gameP.ID).Find(&review)

						return review, nil
					},
				},
				"video": &graphql.Field{
					Type: graphql.NewList(GetVideoType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						gameP := params.Source.(models.Game)

						var video []models.VideoGame
						db.Where("game_id = ?", gameP.ID).Find(&video)

						return video, nil
					},
				},
				"photo": &graphql.Field{
					Type: graphql.NewList(GetPhotoType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						gameP := params.Source.(models.Game)

						var photo []models.PhotoGame
						db.Where("game_id = ?", gameP.ID).Find(&photo)

						return photo, nil
					},
				},
				"systemrequirement": &graphql.Field{
					Type: graphql.NewList(GetSystemRequirementType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						gameP := params.Source.(models.Game)

						var systemReq []models.SystemRequirement
						db.Where("game_id = ?", gameP.ID).Find(&systemReq)

						return systemReq, nil
					},
				},
			},
		})
	}

	return gameType
}


