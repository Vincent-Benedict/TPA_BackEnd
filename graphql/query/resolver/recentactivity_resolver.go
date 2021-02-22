package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetRecentActivity(p graphql.ResolveParams) (i interface{}, e error){

	userid := p.Args["userid"].(int)

	activity, err := models.GetAllRecentActivity(userid)

	if err!=nil{
		panic(err)
	}

	return activity, nil
}

func InsertRecentActivity(p graphql.ResolveParams) (i interface{}, e error){

	token := p.Args["token"].(string)
	activity := p.Args["activity"].(string)

	act, err := models.InsertRecentActivity(token, activity)

	if err!=nil{
		panic(err)
	}

	return act, nil
}

func InsertRecentActivityById(p graphql.ResolveParams) (i interface{}, e error){

	userid := p.Args["userid"].(int)
	activity := p.Args["activity"].(string)

	act, err := models.InsertRecentActivityById(userid, activity)

	if err!=nil{
		panic(err)
	}

	return act, nil
}

func GetRecentActivityOffsetLimit(p graphql.ResolveParams) (i interface{}, e error){

	userid := p.Args["userid"].(int)
	offset := p.Args["offset"].(int)
	limit := p.Args["limit"].(int)

	act, err := models.GetRecentActivityOffsetLimit(userid, offset, limit)

	if err!=nil{
		panic(err)
	}

	return act, nil
}
