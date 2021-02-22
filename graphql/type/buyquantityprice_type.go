package _type

import (
	"github.com/graphql-go/graphql"
)

var buyQuantityPriceType *graphql.Object

func GetBuyQuantityPriceType() *graphql.Object{

	if buyQuantityPriceType == nil{
		buyQuantityPriceType = graphql.NewObject(graphql.ObjectConfig{
			Name:"buyQuantityPriceType",
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

	return buyQuantityPriceType
}


