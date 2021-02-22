package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func CheckUserIsHavingGame(p graphql.ResolveParams) (i interface{}, e error){
	token := p.Args["token"].(string)
	gameid := p.Args["gameid"].(int)

	reviews, _:= models.IsHavingAGame(token, gameid)

	return reviews, nil
}

func InsertUserGame(p graphql.ResolveParams)(i interface{}, e error){
	token := p.Args["token"].(string)
	gameid := p.Args["gameid"].(int)

	g, _:= models.InsertUserGame(token, gameid)

	return g, nil
}

func InsertUserGameById(p graphql.ResolveParams)(i interface{}, e error){
	id := p.Args["userid"].(int)
	gameid := p.Args["gameid"].(int)

	g, _:= models.InsertUserGameById(id, gameid)

	return g, nil
}

func GetUserGameByUserId(p graphql.ResolveParams)(i interface{}, e error){
	userid := p.Args["userid"].(int)

	userGame,_ := models.GetUserGame(userid)

	return userGame, nil
}

func GetUserGameByUserIdLimit(p graphql.ResolveParams)(i interface{}, e error){
	token := p.Args["token"].(string)
	limit := p.Args["limit"].(int)
	userGame,_ := models.GetUserGameLimit(token, limit)

	return userGame, nil
}
