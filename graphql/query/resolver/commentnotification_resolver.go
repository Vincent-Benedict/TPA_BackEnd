package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetCommentNotificationByToken(p graphql.ResolveParams)(i interface{}, e error){

	jwt := p.Args["token"].(string)

	friend, _ := models.GetCommentNotificationByToken(jwt)

	return friend, nil
}

func InsertCommentNotification(p graphql.ResolveParams)(i interface{}, e error){

	userid := p.Args["userid"].(int)
	friendid := p.Args["friendid"].(int)

	friend, _ := models.InsertCommentNotification(userid, friendid)

	return friend, nil
}

func DeleteCommentNotification(p graphql.ResolveParams)(i interface{}, e error){

	userid := p.Args["userid"].(int)
	friendid := p.Args["friendid"].(int)

	friend, _ := models.DeleteCommentNotification(userid, friendid)

	return friend, nil
}
