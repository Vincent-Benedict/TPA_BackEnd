package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetAllShopChatSticker(p graphql.ResolveParams) (i interface{}, e error){

	chatSticker, _ := models.GetAllShopChatSticker()

	return chatSticker, nil
}

