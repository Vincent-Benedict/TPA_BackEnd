package _type

import "github.com/graphql-go/graphql"

var shopBackgroundType *graphql.Object

func GetShopBackgroundType() *graphql.Object{

	if shopBackgroundType == nil{
		shopBackgroundType = graphql.NewObject(graphql.ObjectConfig{
			Name:"shopBackgroundType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"backgroundname": &graphql.Field{
					Type: graphql.String,
				},
				"backgroundimage": &graphql.Field{
					Type: graphql.String,
				},
				"backgroundpoint": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}

	return shopBackgroundType
}



