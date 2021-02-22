package _type

import "github.com/graphql-go/graphql"

var shopAvatarType *graphql.Object

func GetShopAvatarType() *graphql.Object{

	if shopAvatarType == nil{
		shopAvatarType = graphql.NewObject(graphql.ObjectConfig{
			Name:"shopAvatarType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"avatarname": &graphql.Field{
					Type: graphql.String,
				},
				"avatarimage": &graphql.Field{
					Type: graphql.String,
				},
				"avatarpoint": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}

	return shopAvatarType
}

