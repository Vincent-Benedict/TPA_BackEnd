package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func InsertBuyListing(p graphql.ResolveParams)(i interface{}, e error){

	token := p.Args["token"].(string)
	itemid := p.Args["itemid"].(int)
	quantity := p.Args["quantity"].(int)
	price := p.Args["price"].(int)

	bol, _ := models.InsertBuyListing(token, itemid, quantity, price)

	return bol, nil
}

func GetBuyListingByUserId(p graphql.ResolveParams)(i interface{}, e error){

	token := p.Args["token"].(string)
	itemid := p.Args["itemid"].(int)

	bol, _ := models.GetBuyListingByUserId(token, itemid)

	return bol, nil
}

func BuySoldItem(p graphql.ResolveParams)(i interface{}, e error){

	token := p.Args["token"].(string)
	itemid := p.Args["itemid"].(int)
	price := p.Args["price"].(int)
	quantity := p.Args["quantity"].(int)

	bol, _ := models.BuySoldItem(token, itemid, price, quantity)

	return bol, nil
}
