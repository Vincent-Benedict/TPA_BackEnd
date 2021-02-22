package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var unsuspendRequestType *graphql.Object

func GetUnsuspendRequestType() *graphql.Object{

	if unsuspendRequestType == nil{
		unsuspendRequestType = graphql.NewObject(graphql.ObjectConfig{
			Name:"unsuspendRequestType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"username": &graphql.Field{
					Type: graphql.String,
				},
				"user": &graphql.Field{
					Type: GetUserType(),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						user := params.Source.(models.UnsuspendRequest)

						var users models.User
						db.Where("username = ?", user.Username).First(&users)

						return users, nil
					},
				},

			},
		})
	}

	return unsuspendRequestType
}





