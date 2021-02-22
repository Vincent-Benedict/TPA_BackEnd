package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetChatNotificationByToken(p graphql.ResolveParams)(i interface{}, e error){

	jwt := p.Args["token"].(string)

	friend, _ := models.GetChatNotificationByToken(jwt)

	return friend, nil
}

func InsertChatNotification(p graphql.ResolveParams)(i interface{}, e error){

	userid := p.Args["userid"].(int)
	friendid := p.Args["friendid"].(int)

	friend, _ := models.InsertChatNotification(userid, friendid)

	return friend, nil
}

func DeleteChatNotification(p graphql.ResolveParams)(i interface{}, e error){

	userid := p.Args["userid"].(int)
	friendid := p.Args["friendid"].(int)

	friend, _ := models.DeleteChatNotification(userid, friendid)

	return friend, nil
}
