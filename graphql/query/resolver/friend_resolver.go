package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetFriendsByUserId(p graphql.ResolveParams)(i interface{}, e error){

	userid := p.Args["userid"].(int)

	friend, _ := models.GetFriendsByUserId(userid)

	return friend, nil
}
