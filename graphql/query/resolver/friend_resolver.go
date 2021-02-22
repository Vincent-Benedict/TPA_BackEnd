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

func IsFriend(p graphql.ResolveParams)(i interface{}, e error){

	userid := p.Args["userid"].(int)
	friendid := p.Args["friendid"].(int)

	friend, _ := models.IsFriend(userid, friendid)

	return friend, nil
}

func InsertFriend(p graphql.ResolveParams)(i interface{}, e error){

	userid := p.Args["userid"].(int)
	friendid := p.Args["friendid"].(int)

	friend, _ := models.InsertFriend(userid, friendid)

	return friend, nil
}

func GetFriendsByUsername(p graphql.ResolveParams)(i interface{}, e error){

	userid := p.Args["userid"].(int)
	username := p.Args["username"].(string)

	friend, _ := models.GetFriendsByUsername(userid, username)

	return friend, nil
}
