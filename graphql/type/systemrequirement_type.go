package _type

import "github.com/graphql-go/graphql"

var systemRequirementType *graphql.Object

func GetSystemRequirementType() *graphql.Object{

	if systemRequirementType == nil{
		systemRequirementType = graphql.NewObject(graphql.ObjectConfig{
			Name:"systemRequirementType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"gameid": &graphql.Field{
					Type: graphql.Int,
				},
				"os": &graphql.Field{
					Type: graphql.String,
				},
				"processor": &graphql.Field{
					Type: graphql.String,
				},
				"memory": &graphql.Field{
					Type: graphql.String,
				},
				"graphics": &graphql.Field{
					Type: graphql.String,
				},
				"storage": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return systemRequirementType
}
