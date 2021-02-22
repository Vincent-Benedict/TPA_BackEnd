package _type

import "github.com/graphql-go/graphql"

var reviewType *graphql.Object

func GetReviewType() *graphql.Object{

	if reviewType == nil{
		reviewType = graphql.NewObject(graphql.ObjectConfig{
			Name:"reviewType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"gameid": &graphql.Field{
					Type: graphql.Int,
				},
				"reviewdesc": &graphql.Field{
					Type: graphql.String,
				},
				"reviewusername": &graphql.Field{
					Type: graphql.String,
				},
				"reviewavatar": &graphql.Field{
					Type: graphql.String,
				},
				"reviewupvoted": &graphql.Field{
					Type: graphql.Int,
				},
				"reviewdownvoted": &graphql.Field{
					Type: graphql.Int,
				},
				"reviewsentiment": &graphql.Field{
					Type: graphql.String,
				},

			},
		})
	}

	return reviewType
}
