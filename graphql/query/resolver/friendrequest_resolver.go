package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetFriendRequestByUserId(p graphql.ResolveParams)(i interface{}, e error){

	jwt := p.Args["token"].(string)

	friend, _ := models.GetFriendRequestById(jwt)

	return friend, nil
}

func GetFriendRequestByUserIdSent(p graphql.ResolveParams)(i interface{}, e error){

	jwt := p.Args["token"].(string)

	friend, _ := models.GetFriendRequestByIdSent(jwt)

	return friend, nil
}

func InsertFriendRequest(p graphql.ResolveParams)(i interface{}, e error){

	userid := p.Args["userid"].(int)
	friendid := p.Args["friendid"].(int)

	friend, _ := models.InsertFriendRequest(userid, friendid)

	return friend, nil
}

func DeleteFriendRequest(p graphql.ResolveParams)(i interface{}, e error){

	userid := p.Args["userid"].(int)

	friend, _ := models.DeleteFriendRequest(userid)

	return friend, nil
}

func DeleteFriendRequest2(p graphql.ResolveParams)(i interface{}, e error){

	userid := p.Args["userid"].(int)
	friendid := p.Args["friendid"].(int)

	friend, _ := models.DeleteFriendRequest2(userid, friendid)

	return friend, nil
}