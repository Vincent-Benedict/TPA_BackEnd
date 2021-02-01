package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetUser(p graphql.ResolveParams) (i interface{}, e error){

	username := p.Args["username"].(string)
	password := p.Args["password"].(string)

	user, err := models.GetUser(username, password)

	if err!=nil{
		panic(err)
	}

	return user, nil
}

func GetIdFromJwtToken(p graphql.ResolveParams)(i interface{}, e error){
	token := p.Args["token"].(string)

	id, _ := models.GetIdFromJwtToken(token);

	return id, nil
}

func GetUserById(p graphql.ResolveParams)(i interface{}, e error){

	id := p.Args["id"].(int)

	user, err := models.GetUserById(id)

	if err!=nil{
		panic(err)
	}

	return user, nil
}

func GetUserByToken(p graphql.ResolveParams)(i interface{}, e error){

	token := p.Args["token"].(string)

	user, err := models.GetUserByToken(token)

	if err!=nil{
		panic(err)
	}

	return user, nil
}

func DecreaseBalance(p graphql.ResolveParams)(i interface{}, e error){
	token := p.Args["token"].(string)
	balance := p.Args["balance"].(int)

	user, _ := models.DecreaseBalance(token, balance)

	return user, nil
}

func IncreaseBalance(p graphql.ResolveParams)(i interface{}, e error){
	token := p.Args["token"].(string)
	balance := p.Args["balance"].(int)

	user, _ := models.IncreaseBalance(token, balance)

	return user, nil
}
