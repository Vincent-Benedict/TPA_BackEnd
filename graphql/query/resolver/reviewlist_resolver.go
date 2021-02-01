package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func InsertReviewList(p graphql.ResolveParams) (i interface{}, e error){

	token := p.Args["token"].(string)
	gameid := p.Args["gameid"].(int)
	reviewid := p.Args["reviewid"].(int)

	reviews, _:= models.InsertReviewList(token, gameid, reviewid)

	return reviews, nil
}

func CheckIsAdded(p graphql.ResolveParams) (i interface{}, e error){
	token := p.Args["token"].(string)
	gameid := p.Args["gameid"].(int)
	reviewid := p.Args["reviewid"].(int)

	reviews, _:= models.CheckIsAdded(token, gameid, reviewid)

	return reviews, nil
}
