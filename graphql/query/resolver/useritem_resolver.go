package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetAllUserItemWithGameId(p graphql.ResolveParams) (i interface{}, e error){

	token := p.Args["token"].(string)
	gameid := p.Args["gameid"].(int)

	useritem, _ := models.GetAllUserItemWithGameId(token, gameid)


	return useritem, nil
}

func GetUserItemWithGameIdOffsetLimit(p graphql.ResolveParams) (i interface{}, e error){

	token := p.Args["token"].(string)
	gameid := p.Args["gameid"].(int)
	offset := p.Args["offset"].(int)
	limit := p.Args["limit"].(int)

	useritem, _ := models.GetUserItemWithGameIdOffsetLimit(token, gameid, offset, limit)


	return useritem, nil
}

func GetUserItemWithGameIdOffsetLimitByName(p graphql.ResolveParams) (i interface{}, e error){

	token := p.Args["token"].(string)
	gameid := p.Args["gameid"].(int)
	offset := p.Args["offset"].(int)
	limit := p.Args["limit"].(int)
	gamename := p.Args["gamename"].(string)

	useritem, _ := models.GetUserItemWithGameIdOffsetLimitByName(token, gameid, offset, limit, gamename)


	return useritem, nil
}

func GetUserItemWithGameIdByName(p graphql.ResolveParams) (i interface{}, e error){

	token := p.Args["token"].(string)
	gameid := p.Args["gameid"].(int)
	gamename := p.Args["gamename"].(string)

	useritem, _ := models.GetUserItemWithGameIdByName(token, gameid, gamename)


	return useritem, nil
}
