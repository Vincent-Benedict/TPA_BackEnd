package _type

import (
	"github.com/graphql-go/graphql"
)

var sellQuantityPriceType *graphql.Object

func GetSellQuantityPriceType() *graphql.Object{

	if sellQuantityPriceType == nil{
		sellQuantityPriceType = graphql.NewObject(graphql.ObjectConfig{
			Name:"sellQuantityPriceType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"qty": &graphql.Field{
					Type: graphql.Int,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}

	return sellQuantityPriceType
}

