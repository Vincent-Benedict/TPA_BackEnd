package _type

import "github.com/graphql-go/graphql"

var photoType *graphql.Object

func GetPhotoType() *graphql.Object{

	if photoType == nil{
		photoType = graphql.NewObject(graphql.ObjectConfig{
			Name:"photoType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"gameid": &graphql.Field{
					Type: graphql.Int,
				},
				"photosource": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return photoType
}
