package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetAllGameItem(p graphql.ResolveParams)(i interface{}, e error){

	gameItem, _ := models.GetAllGameItem()

	return gameItem, nil
}

func GetGameItemById(p graphql.ResolveParams) (i interface{}, e error){

	id:= p.Args["id"].(int)

	gameItem, _:= models.GetItemGameById(id)

	return gameItem, nil
}
