package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var recentActivityType *graphql.Object

func GetRecentActivityType() *graphql.Object{

	if recentActivityType == nil{
		recentActivityType = graphql.NewObject(graphql.ObjectConfig{
			Name:"recentActivityType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"userid": &graphql.Field{
					Type: graphql.Int,
				},
				"activity": &graphql.Field{
					Type: graphql.String,
				},
				"createdat": &graphql.Field{
					Type: graphql.String,
				},
				"user": &graphql.Field{
					Type: GetUserType(),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						users := params.Source.(models.RecentActivity)

						var user models.User
						db.Where("id = ?", users.UserId).Find(&user)

						return user, nil
					},
				},

			},
		})
	}

	return recentActivityType
}