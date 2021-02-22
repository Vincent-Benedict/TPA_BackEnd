package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var marketRecentActivityType *graphql.Object

func GetMarketRecentActivityTypeType() *graphql.Object{

	if marketRecentActivityType == nil{
		marketRecentActivityType = graphql.NewObject(graphql.ObjectConfig{
			Name:"marketRecentActivityType",
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
				"activity": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
				"user": &graphql.Field{
					Type: graphql.NewList(GetUserType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						users := params.Source.(models.MarketRecentActivity)

						var user []models.User
						db.Where("id = ?", users.UserId).Find(&user)

						return user, nil
					},
				},
			},
		})
	}

	return marketRecentActivityType
}
