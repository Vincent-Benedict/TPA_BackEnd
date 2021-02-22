package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetAllUserBadgeByUserId(p graphql.ResolveParams)(i interface{}, e error){

	token := p.Args["token"].(string)

	bol, _ := models.GetAllUserBadgeByUserId(token)

	return bol, nil
}

func GetUserBadgeById(p graphql.ResolveParams)(i interface{}, e error){

	id := p.Args["id"].(int)

	bol, _ := models.GetUserBadgeById(id)

	return bol, nil
}

func GetBadgeCardById(p graphql.ResolveParams)(i interface{}, e error){

	id := p.Args["id"].(int)

	bol, _ := models.GetBadgeCardById(id)

	return bol, nil
}