package _type

import "github.com/graphql-go/graphql"

var shopChatStickerType *graphql.Object

func GetChatStickerType() *graphql.Object{

	if shopChatStickerType == nil{
		shopChatStickerType = graphql.NewObject(graphql.ObjectConfig{
			Name:"shopChatStickerType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"chatstickername": &graphql.Field{
					Type: graphql.String,
				},
				"chatstickerimage": &graphql.Field{
					Type: graphql.String,
				},
				"chatstickerpoint": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}

	return shopChatStickerType
}




