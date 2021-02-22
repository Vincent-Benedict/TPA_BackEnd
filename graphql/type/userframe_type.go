package _type

import "github.com/graphql-go/graphql"

var userFrameType *graphql.Object

func GetUserFrameType() *graphql.Object{

	if userFrameType == nil{
		userFrameType = graphql.NewObject(graphql.ObjectConfig{
			Name:"userFrameType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"userid": &graphql.Field{
					Type: graphql.Int,
				},
				"frame": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return userFrameType
}
