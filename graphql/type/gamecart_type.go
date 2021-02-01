package _type

import (
	"github.com/graphql-go/graphql"
)

var gameCartType *graphql.Object

func GetGameCartType() *graphql.Object{

	if gameCartType == nil{
		gameCartType = graphql.NewObject(graphql.ObjectConfig{
			Name:"gameCartType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"gameid": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"image": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
				"discount": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}

	return gameCartType
}