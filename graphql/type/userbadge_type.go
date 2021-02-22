package _type

import "github.com/graphql-go/graphql"

var userBadgeType *graphql.Object

func GetUserBadgeType() *graphql.Object{

	if userBadgeType == nil{
		userBadgeType = graphql.NewObject(graphql.ObjectConfig{
			Name:"userBadgeType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"userid": &graphql.Field{
					Type: graphql.Int,
				},
				"badgeimg": &graphql.Field{
					Type: graphql.String,
				},
				"badgename": &graphql.Field{
					Type: graphql.String,
				},
				"badgexp": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}

	return userBadgeType
}

