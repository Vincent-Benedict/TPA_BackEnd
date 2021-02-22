package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetReportByToken(p graphql.ResolveParams) (i interface{}, e error){

	token := p.Args["token"].(string)

	report, err := models.GetReportedByJWTToken(token)

	if err!=nil{
		panic(err)
	}

	return report, nil
}

func InsertReport(p graphql.ResolveParams) (i interface{}, e error){

	token := p.Args["token"].(string)
	reportedid := p.Args["reportedid"].(int)
	reportreason := p.Args["reportreason"].(string)

	report, err := models.InsertReport(token, reportedid, reportreason)

	if err!=nil{
		panic(err)
	}

	return report, nil
}

func DeleteReportByReportedId(p graphql.ResolveParams) (i interface{}, e error){

	reportedid := p.Args["reportedid"].(int)

	report, err := models.DeleteReportByReportedId(reportedid)

	if err!=nil{
		panic(err)
	}

	return report, nil
}

func GetReportByReportedId(p graphql.ResolveParams) (i interface{}, e error){

	reportedid := p.Args["id"].(int)

	report, err := models.GetReportByReportedId(reportedid)

	if err!=nil{
		panic(err)
	}

	return report, nil
}
