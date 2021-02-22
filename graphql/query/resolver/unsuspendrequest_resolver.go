package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetAllUnsupendRequest(p graphql.ResolveParams) (i interface{}, e error){

	unsuspendRequest, _ := models.GetAllUnsuspendRequest()

	return unsuspendRequest, nil
}

func InsertUnsuspendRequest(p graphql.ResolveParams) (i interface{}, e error){
	username := p.Args["username"].(string)

	unsuspendRequest, _ := models.InsertUnsuspendRequest(username)

	return unsuspendRequest, nil
}

func DeleteUnsuspendRequest(p graphql.ResolveParams) (i interface{}, e error){
	id := p.Args["id"].(int)

	unsuspendRequest, _ := models.DeleteUnsuspendRequest(id)

	return unsuspendRequest, nil
}
