package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetGamesCart(p graphql.ResolveParams) (i interface{}, e error){

	gamesCart, err := models.GetGamesInCart()

	if err!=nil{
		panic(err)
	}

	return gamesCart, nil
}

func InsertGamesToCart(p graphql.ResolveParams)(i interface{}, e error){

	gameid := p.Args["gameid"].(int)
	name := p.Args["name"].(string)
	image := p.Args["image"].(string)
	price := p.Args["price"].(int)
	discount := p.Args["discount"].(int)

	gameCart, _ := models.InsertGameToCart(gameid ,name, image, price, discount)

	return gameCart, nil
}

func DeleteGameFromCartById(p graphql.ResolveParams)(i interface{}, e error){
	id := p.Args["id"].(int)

	idGameCart,_ := models.DeleteGameFromCartById(id)

	return idGameCart, nil
}

func DeleteGameFromCartById2(p graphql.ResolveParams)(i interface{}, e error){
	id := p.Args["id"].(int)

	idGameCart,_ := models.DeleteGameFromCartById2(id)

	return idGameCart, nil
}

func DeleteAllGameFromCart(p graphql.ResolveParams)(i interface{}, e error){
	idGameCart,_ := models.DeleteAllGameFromCart()

	return idGameCart, nil
}