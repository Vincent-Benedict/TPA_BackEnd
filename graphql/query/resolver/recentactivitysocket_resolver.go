package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/Global"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/functionalfoundry/graphqlws"
	"github.com/graphql-go/graphql"
	"log"
)

func InsertMarketRecentActivity(p graphql.ResolveParams)(i interface{}, e error){


	userid := p.Args["userid"].(int)
	itemid := p.Args["itemid"].(int)
	activity := p.Args["activity"].(string)
	price := p.Args["price"].(int)

	boolean, _ := models.InsertMarketRecentActivity(userid, itemid, activity, price)

	/* Template Broadcast*/
	subscriptions := Global.SubscriptionManager.Subscriptions()
	log.Print(subscriptions)

	for conn, _ := range subscriptions {
		log.Print(subscriptions[conn])
		for _, subscription := range subscriptions[conn] {
			log.Print(subscription)
			params := graphql.Params{
				Schema:         Global.Schema,
				RequestString:  subscription.Query,
				VariableValues: subscription.Variables,
				OperationName:  subscription.OperationName,
			}
			result := graphql.Do(params)

			data := graphqlws.DataMessagePayload{
				Data:   result.Data,
				Errors: graphqlws.ErrorsFromGraphQLErrors(result.Errors),
			}
			log.Print(data)

			subscription.SendData(&data)
		}
	}

	return boolean, nil
}

func GetLastMarketActivity(p graphql.ResolveParams)(i interface{}, e error){

	itemid := p.Args["itemid"].(int)

	activity, _ := models.GetLastMarketActivity(itemid)

	return activity, nil
}


func GetAllMarketActivity(p graphql.ResolveParams)(i interface{}, e error){

	itemid := p.Args["itemid"].(int)

	activity, _ := models.GetAllMarketActivity(itemid)

	return activity, nil
}
