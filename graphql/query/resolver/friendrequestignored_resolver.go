package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetFriendRequestIgnoredByUserId(p graphql.ResolveParams)(i interface{}, e error){

	jwt := p.Args["token"].(string)

	friend, _ := models.GetFriendRequestIgnoredById(jwt)

	return friend, nil
}


func InsertFriendRequestIgnored(p graphql.ResolveParams)(i interface{}, e error){

	userid := p.Args["userid"].(int)
	friendid := p.Args["friendid"].(int)

	friend, _ := models.InsertFriendRequestIgnored(userid, friendid)

	return friend, nil
}
