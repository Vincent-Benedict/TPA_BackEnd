package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var buyListingType *graphql.Object

func GetBuyListingType() *graphql.Object{

	if buyListingType == nil{
		buyListingType = graphql.NewObject(graphql.ObjectConfig{
			Name:"buyListingType",
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
				"quantity": &graphql.Field{
					Type: graphql.Int,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
				"user": &graphql.Field{
					Type: graphql.NewList(GetUserType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						user := params.Source.(models.BuyListing)

						var users []models.User
						db.Where("id = ?", user.UserId).Find(&users)

						return users, nil
					},
				},
			},
		})
	}

	return buyListingType
}



