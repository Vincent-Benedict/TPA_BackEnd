package _type

import (
	"github.com/graphql-go/graphql"
)

var userThemeType *graphql.Object

func GetUserThemeType() *graphql.Object{

	if userThemeType == nil{
		userThemeType = graphql.NewObject(graphql.ObjectConfig{
			Name:"userThemeType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"theme": &graphql.Field{
					Type: graphql.String,
				},
				"color": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return userThemeType
}
