package _type

import "github.com/graphql-go/graphql"

var reviewListType *graphql.Object

func GetReviewListType() *graphql.Object{

	if reviewListType == nil{
		reviewListType = graphql.NewObject(graphql.ObjectConfig{
			Name:"reviewListType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"userid": &graphql.Field{
					Type: graphql.Int,
				},
				"gameid": &graphql.Field{
					Type: graphql.Int,
				},
				"reviewid": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}

	return reviewListType
}
