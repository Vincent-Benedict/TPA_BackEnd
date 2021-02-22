package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetGiftNotificationByToken(p graphql.ResolveParams)(i interface{}, e error){

	jwt := p.Args["token"].(string)

	friend, _ := models.GetGiftNotificationByToken(jwt)

	return friend, nil
}

func InsertGiftNotification(p graphql.ResolveParams)(i interface{}, e error){

	userid := p.Args["userid"].(int)
	friendid := p.Args["friendid"].(int)

	friend, _ := models.InsertGiftNotification(userid, friendid)

	return friend, nil
}

func DeleteGiftNotification(p graphql.ResolveParams)(i interface{}, e error){

	userid := p.Args["userid"].(int)
	friendid := p.Args["friendid"].(int)

	friend, _ := models.DeleteGiftNotification(userid, friendid)

	return friend, nil
}
