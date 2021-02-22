package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetSellListingQuantityAndPriceByItemId(p graphql.ResolveParams)(i interface{}, e error){

	itemid := p.Args["itemid"].(int)

	qtyprice, _ := models.GetSellListingQuantityAndPriceByItemId(itemid)

	return qtyprice, nil
}


