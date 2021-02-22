package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetAllShopAvatar(p graphql.ResolveParams) (i interface{}, e error){

	avatar, _ := models.GetAllShopAvatar()

	return avatar, nil
}
