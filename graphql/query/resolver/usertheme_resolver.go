package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetUserTheme(p graphql.ResolveParams) (i interface{}, e error){

	user, err := models.GetAllTheme()

	if err!=nil{
		panic(err)
	}

	return user, nil
}

func GetUserThemeById(p graphql.ResolveParams) (i interface{}, e error){

	id := p.Args["id"].(int)

	user, err := models.GetThemeById(id)

	if err!=nil{
		panic(err)
	}

	return user, nil
}
