package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetCommentByUserId(p graphql.ResolveParams)(i interface{}, e error){

	userid := p.Args["userid"].(int)

	comments, _ := models.GetAllCommentsByUserId(userid)

	return comments, nil
}


func InsertComment(p graphql.ResolveParams)(i interface{}, e error){

	token := p.Args["token"].(string)
	friendid := p.Args["friendid"].(int)
	commentdesc := p.Args["commentdesc"].(string)

	comments, _ := models.InsertComment(token, friendid, commentdesc)

	return comments, nil
}

