package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetRedeemCodeByString(p graphql.ResolveParams) (i interface{}, e error){

	code:= p.Args["code"].(string)

	redeemCode, _:= models.GetRedeemCodeByString(code)

	return redeemCode, nil
}
