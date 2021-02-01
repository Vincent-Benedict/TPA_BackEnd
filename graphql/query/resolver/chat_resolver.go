package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/Global"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/functionalfoundry/graphqlws"
	"github.com/graphql-go/graphql"
	"log"
)

func InsertChat(p graphql.ResolveParams)(i interface{}, e error){

	message := p.Args["message"].(string)
	userid := p.Args["userid"].(int)
	recipientid := p.Args["recipientid"].(int)

	chat, _ := models.InsertChat(message, userid, recipientid)

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

	return chat, nil
}

func GetLastChat(p graphql.ResolveParams)(i interface{}, e error){

	userid := p.Args["userid"].(int)
	recipientid := p.Args["recipientid"].(int)

	chat, _ := models.GetLastChat(userid, recipientid)

	return chat, nil
}


func GetAllChat(p graphql.ResolveParams)(i interface{}, e error){

	userid := p.Args["userid"].(int)
	recipientid := p.Args["recipientid"].(int)

	chat, _ := models.GetAllChat(userid, recipientid)

	return chat, nil
}

