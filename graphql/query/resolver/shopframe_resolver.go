package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetAllShopFrame(p graphql.ResolveParams) (i interface{}, e error){

	frame, _ := models.GetAllShopFrame()

	return frame, nil
}

func InsertUserFrame(p graphql.ResolveParams)(i interface{}, e error){
	token := p.Args["token"].(string)
	frame := p.Args["frame"].(string)

	g, _:= models.InsertUserFrame(token, frame)

	return g, nil
}