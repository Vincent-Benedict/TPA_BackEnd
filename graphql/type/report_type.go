package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var reportType *graphql.Object

func GetReportType() *graphql.Object{

	if reportType == nil{
		reportType = graphql.NewObject(graphql.ObjectConfig{
			Name:"reportType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"reportedid": &graphql.Field{
					Type: graphql.Int,
				},
				"reported": &graphql.Field{
					Type: GetUserType(),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					db, _:= database.Connect()
					defer db.Close()
					user := params.Source.(models.Report)

					var users models.User
					db.Where("id = ?", user.ReportedId).First(&users)

					return users, nil
					},
				},
				"reporterid": &graphql.Field{
					Type: graphql.Int,
				},
				"reporter": &graphql.Field{
					Type: GetUserType(),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						user := params.Source.(models.Report)

						var users models.User
						db.Where("id = ?", user.ReporterId).First(&users)

						return users, nil
					},
				},
				"reportreason": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return reportType
}
