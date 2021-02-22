package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetItemNotificationByToken(p graphql.ResolveParams)(i interface{}, e error){

	jwt := p.Args["token"].(string)

	friend, _ := models.GetItemNotificationByToken(jwt)

	return friend, nil
}

func InsertItemNotification(p graphql.ResolveParams)(i interface{}, e error){

	userid := p.Args["userid"].(int)
	itemid := p.Args["itemid"].(int)

	friend, _ := models.InsertItemNotification(userid, itemid)

	return friend, nil
}

func DeleteItemNotification(p graphql.ResolveParams)(i interface{}, e error){

	id := p.Args["id"].(int)

	friend, _ := models.DeleteItemNotification(id)

	return friend, nil
}
