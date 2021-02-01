package _type

import (
	"github.com/graphql-go/graphql"
)

var chatType *graphql.Object

func GetChatType() *graphql.Object{

	if chatType == nil{
		chatType = graphql.NewObject(graphql.ObjectConfig{
			Name:"chatType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"senderid": &graphql.Field{
					Type: graphql.Int,
				},
				"recipientid": &graphql.Field{
					Type: graphql.Int,
				},
				"message": &graphql.Field{
					Type: graphql.String,
				},
				"createdat": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return chatType
}
