package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var itemNotificationType *graphql.Object

func GetItemNotificationType() *graphql.Object{

	if itemNotificationType == nil{
		itemNotificationType = graphql.NewObject(graphql.ObjectConfig{
			Name:"itemNotificationType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"userid": &graphql.Field{
					Type: graphql.Int,
				},
				"user":&graphql.Field{
					Type: graphql.NewList(GetUserType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						us := params.Source.(models.ItemNotification)

						var user []models.User
						db.Where("id = ?", us.UserId).Find(&user)

						return user, nil
					},
				},
				"itemid": &graphql.Field{
					Type: graphql.Int,
				},
				"item": &graphql.Field{
					Type: graphql.NewList(GetGameItemType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						item := params.Source.(models.ItemNotification)

						var items []models.GameItem
						db.Where("id = ?", item.ItemId).Find(&items)

						return items, nil
					},
				},
			},
		})
	}

	return itemNotificationType
}
