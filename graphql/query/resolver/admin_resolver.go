package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetAdmin(p graphql.ResolveParams)(i interface{}, e error){

	username := p.Args["username"].(string)
	password := p.Args["password"].(string)

	bol, _ := models.GetAdmin(username, password)

	return bol, nil
}

