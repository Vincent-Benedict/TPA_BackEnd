package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var userItem *graphql.Object

func GetUserItemType() *graphql.Object{

	if userItem == nil{
		userItem = graphql.NewObject(graphql.ObjectConfig{
			Name:"userItem",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"userid": &graphql.Field{
					Type: graphql.Int,
				},
				"itemid": &graphql.Field{
					Type: graphql.Int,
				},
				"item": &graphql.Field{
					Type: graphql.NewList(GetGameItemType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						item := params.Source.(models.UserItem)

						var gameItem []models.GameItem
						db.Where("id = ?", item.ItemId).Find(&gameItem)

						return gameItem, nil
					},
				},
			},
		})
	}

	return userItem
}
