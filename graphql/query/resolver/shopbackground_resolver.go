package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetAllShopBackground(p graphql.ResolveParams) (i interface{}, e error){

	background, _ := models.GetAllShopBackground()

	return background, nil
}

func InsertUserBackground(p graphql.ResolveParams)(i interface{}, e error){
	token := p.Args["token"].(string)
	background := p.Args["background"].(string)

	g, _:= models.InsertUserBackground(token, background)

	return g, nil
}

