package _type

import "github.com/graphql-go/graphql"

var shopMiniBackgroundType *graphql.Object

func GetShopMiniBackgroundType() *graphql.Object{

	if shopMiniBackgroundType == nil{
		shopMiniBackgroundType = graphql.NewObject(graphql.ObjectConfig{
			Name:"shopMiniBackgroundType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"minibackgroundname": &graphql.Field{
					Type: graphql.String,
				},
				"minibackgroundimage": &graphql.Field{
					Type: graphql.String,
				},
				"minibackgroundpoint": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}

	return shopMiniBackgroundType
}




