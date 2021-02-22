package _type

import "github.com/graphql-go/graphql"

var shopFrameType *graphql.Object

func GetShopFrameType() *graphql.Object{

	if shopFrameType == nil{
		shopFrameType = graphql.NewObject(graphql.ObjectConfig{
			Name:"shopFrameType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"framename": &graphql.Field{
					Type: graphql.String,
				},
				"frameimage": &graphql.Field{
					Type: graphql.String,
				},
				"framepoint": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}

	return shopFrameType
}


