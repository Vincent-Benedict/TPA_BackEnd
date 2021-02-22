package _type

import "github.com/graphql-go/graphql"

var userMiniProfileType *graphql.Object

func GetUserMiniProfileType() *graphql.Object{

	if userMiniProfileType == nil{
		userMiniProfileType = graphql.NewObject(graphql.ObjectConfig{
			Name:"userMiniProfileType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"userid": &graphql.Field{
					Type: graphql.Int,
				},
				"miniprofile": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return userMiniProfileType
}

