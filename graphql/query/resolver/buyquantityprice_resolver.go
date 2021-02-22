package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetBuyListingQuantityAndPriceByItemId(p graphql.ResolveParams)(i interface{}, e error){

	itemid := p.Args["itemid"].(int)

	qtyprice, _ := models.GetBuyListingQuantityAndPriceByItemId(itemid)

	return qtyprice, nil
}

