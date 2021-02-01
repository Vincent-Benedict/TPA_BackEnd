package _type

import "github.com/graphql-go/graphql"

var videoType *graphql.Object

func GetVideoType() *graphql.Object{

	if videoType == nil{
		videoType = graphql.NewObject(graphql.ObjectConfig{
			Name:"videoType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"gameid": &graphql.Field{
					Type: graphql.Int,
				},
				"videosource": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return videoType
}
