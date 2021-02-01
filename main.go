package main


import (
	"github.com/Vincent-Benedict/TPA-Web/Global"
	"github.com/Vincent-Benedict/TPA-Web/api"
	"github.com/Vincent-Benedict/TPA-Web/graphql/mutation"
	"github.com/Vincent-Benedict/TPA-Web/graphql/query"
	"github.com/Vincent-Benedict/TPA-Web/graphql/subscription"
	"github.com/Vincent-Benedict/TPA-Web/middleware"
	"github.com/functionalfoundry/graphqlws"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

var err error

func main(){
	Global.Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: query.GetRoot(),
		Mutation: mutation.GetRoot(),
		Subscription: subscription.GetRoot(),
	})

	if err != nil{
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema: &Global.Schema,
		Pretty: true,
		GraphiQL: false,
		Playground: true,
	})

	wrapped:= middleware.CorsMiddleware(h)

	Global.SubscriptionManager = graphqlws.NewSubscriptionManager(&Global.Schema)

	// Create a WebSocket/HTTP handler
	graphqlwsHandler := graphqlws.NewHandler(graphqlws.HandlerConfig{
		// Wire up the GraphqL WebSocket handler with the subscription manager
		SubscriptionManager: Global.SubscriptionManager,

		// Optional: Add a hook to resolve auth tokens into users that are
		// then stored on the GraphQL WS connections
		Authenticate: func(authToken string) (interface{}, error) {
			// This is just a dumb example
			return "Joe", nil
		},
	})

	router:= api.NewRouter()
	router.Handle("/api", wrapped)
	router.Handle("/subscriptions", graphqlwsHandler)
	log.Fatalln(http.ListenAndServe(":2000", router))
}
