package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func SendEmail(p graphql.ResolveParams)(i interface{}, e error){

	email := p.Args["email"].(string)

	user, _ := models.SendEmail(email)

	return user, nil
}

func SendEmailTransaction(p graphql.ResolveParams)(i interface{}, e error){

	email := p.Args["email"].(string)
	firstname := p.Args["firstname"].(string)
	message := p.Args["message"].(string)
	sentiment := p.Args["sentiment"].(string)
	signature := p.Args["signature"].(string)

	user, _ := models.SendEmailTransaction(email, firstname, message, sentiment, signature)

	return user, nil
}

func SendEmailWishlist(p graphql.ResolveParams)(i interface{}, e error){

	email := p.Args["email"].(string)
	gamename := p.Args["gamename"].(string)

	user, _ := models.SendEmailWishlist(email, gamename)

	return user, nil
}