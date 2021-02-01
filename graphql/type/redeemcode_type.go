package _type

import "github.com/graphql-go/graphql"

var redeemCodeType *graphql.Object

func GetRedeemCodeType() *graphql.Object{

	if redeemCodeType == nil{
		redeemCodeType = graphql.NewObject(graphql.ObjectConfig{
			Name:"redeemCodeType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"code": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}

	return redeemCodeType
}
