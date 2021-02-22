package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func InsertSellListing(p graphql.ResolveParams)(i interface{}, e error){

	token := p.Args["token"].(string)
	itemid := p.Args["itemid"].(int)
	quantity := p.Args["quantity"].(int)
	price := p.Args["price"].(int)

	bol, _ := models.InsertSellListing(token, itemid, quantity, price)

	return bol, nil
}

func GetSellListingByUserId(p graphql.ResolveParams)(i interface{}, e error){

	token := p.Args["token"].(string)
	itemid := p.Args["itemid"].(int)

	bol, _ := models.GetSellListingByUserId(token, itemid)

	return bol, nil
}

func CheckIfThereIsSellListing(p graphql.ResolveParams)(i interface{}, e error){

	itemid := p.Args["itemid"].(int)
	price := p.Args["price"].(int)
	qty := p.Args["quantity"].(int)

	bol, _ := models.CheckIfThereIsSellListing(itemid, price, qty)

	return bol, nil
}
