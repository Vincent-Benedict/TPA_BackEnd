package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetAllShopMiniBackground(p graphql.ResolveParams) (i interface{}, e error){

	minibackground, _ := models.GetAllShopMiniBackground()

	return minibackground, nil
}

func InsertUserMiniBackground(p graphql.ResolveParams)(i interface{}, e error){
	token := p.Args["token"].(string)
	minibackground := p.Args["minibackground"].(string)

	g, _:= models.InsertUserMiniProfile(token, minibackground)

	return g, nil
}
