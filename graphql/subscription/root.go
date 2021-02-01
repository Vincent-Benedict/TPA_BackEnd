package subscription

import (
	res "github.com/Vincent-Benedict/TPA-Web/graphql/query/resolver"
	typ "github.com/Vincent-Benedict/TPA-Web/graphql/type"
	"github.com/graphql-go/graphql"
)

func GetRoot() *graphql.Object{
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootSubscription",
		Fields: graphql.Fields{
			//"tes":&graphql.Field{
			//	Type: graphql.String,
			//	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			//		return "tes-tes", nil
			//	},
			//},

			"getlastchat": &graphql.Field{
				Type: typ.GetChatType(),
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"recipientid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetLastChat,
			},
		},

	})
}
