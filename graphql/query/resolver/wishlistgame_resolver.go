package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func InsertGamesToWishlist(p graphql.ResolveParams)(i interface{}, e error){

	token := p.Args["token"].(string)
	gameid := p.Args["gameid"].(int)

	gameWishlist, _ := models.InsertGameToWishlist(token, gameid)

	return gameWishlist, nil
}

func DeleteGameFromWishlist(p graphql.ResolveParams)(i interface{}, e error){

	gameid := p.Args["gameid"].(int)

	gameWishlist, _ := models.DeleteGameFromWishlist(gameid)

	return gameWishlist, nil
}

func IsAddedGameToWishlist(p graphql.ResolveParams)(i interface{}, e error){
	token := p.Args["token"].(string)
	gameid := p.Args["gameid"].(int)

	gameWishlist, _ := models.IsAddedGameToWishlist(token, gameid)

	return gameWishlist, nil
}


