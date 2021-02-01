package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetGames(p graphql.ResolveParams) (i interface{}, e error){
	games, err := models.GetGames()

	if err!=nil{
		panic(err)
	}

	return games, nil
}

func GetGameById(p graphql.ResolveParams) (i interface{}, e error){

	id:= p.Args["id"].(int)

	game, _:= models.GetGame(id)

	return game, nil
}

func GetGamesByStatus(p graphql.ResolveParams) (i interface{}, e error){

	status := p.Args["status"].(string)

	games, err := models.GetGamesByStatus(status)

	if err!=nil{
		panic(err)
	}

	return games, nil
}

func GetGamesByStatusLimit(p graphql.ResolveParams) (i interface{}, e error){

	status := p.Args["status"].(string)
	limit := p.Args["limit"].(int)

	games, err := models.GetGamesByStatusLimit(status,limit)

	if err!=nil{
		panic(err)
	}

	return games, nil
}

func GetGamesByStatusDiscountLimit(p graphql.ResolveParams) (i interface{}, e error){

	status := p.Args["status"].(string)
	discount := p.Args["discount"].(int)
	limit := p.Args["limit"].(int)

	games, err := models.GetGamesSpecialsSpecificDiscount(status, discount,limit)

	if err!=nil{
		panic(err)
	}

	return games, nil
}

func GetGamesByNameLimit(p graphql.ResolveParams) (i interface{}, e error){

	name := p.Args["name"].(string)
	limit := p.Args["limit"].(int)

	games, err := models.GetGamesByNameLimit(name,limit)

	if err!=nil{
		panic(err)
	}

	return games, nil
}

func GetGamesByName(p graphql.ResolveParams) (i interface{}, e error){

	name := p.Args["name"].(string)

	games, err := models.GetGamesByName(name)

	if err!=nil{
		panic(err)
	}

	return games, nil
}

func GetGamesByFilter(p graphql.ResolveParams) (i interface{}, e error){

	name := p.Args["name"].(string)
	price := p.Args["price"].(int)
	category := p.Args["category"].(string)
	genre := p.Args["genre"].(string)

	games, err := models.GetGamesByFilter(name, price, category, genre)

	if err!=nil{
		panic(err)
	}

	return games, nil
}






