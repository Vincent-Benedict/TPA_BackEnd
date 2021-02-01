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

