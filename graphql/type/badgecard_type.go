package _type

import "github.com/graphql-go/graphql"

var badgeCardType *graphql.Object

func GetBadgeCardType() *graphql.Object{

	if badgeCardType == nil{
		badgeCardType = graphql.NewObject(graphql.ObjectConfig{
			Name:"badgeCardType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"userbadgeid": &graphql.Field{
					Type: graphql.Int,
				},
				"card": &graphql.Field{
					Type: graphql.String,
				},
				"status": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return badgeCardType
}
