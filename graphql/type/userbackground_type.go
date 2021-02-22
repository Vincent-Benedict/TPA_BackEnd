package _type

import "github.com/graphql-go/graphql"

var userBackgroundType *graphql.Object

func GetUserBackgroundType() *graphql.Object{

	if userBackgroundType == nil{
		userBackgroundType = graphql.NewObject(graphql.ObjectConfig{
			Name:"userBackgroundType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"userid": &graphql.Field{
					Type: graphql.Int,
				},
				"background": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return userBackgroundType
}

