package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetReviewsByGameIdEqualsUpvotedAndLimit(p graphql.ResolveParams) (i interface{}, e error){

	gameid := p.Args["gameid"].(int)
	upvoted := p.Args["reviewupvoted"].(int)
	limit := p.Args["limit"].(int)

	reviews, err := models.GetReviewByEqualsUpvotedLimitAndGameId(gameid, upvoted, limit)

	if err!=nil{
		panic(err)
	}

	return reviews, nil
}

func GetReviewsByGameIdGreaterUpvotedAndLimit(p graphql.ResolveParams) (i interface{}, e error){

	gameid := p.Args["gameid"].(int)
	upvoted := p.Args["reviewupvoted"].(int)
	limit := p.Args["limit"].(int)

	reviews, err := models.GetReviewByGreaterUpvotedLimitAndGameId(gameid, upvoted, limit)

	if err!=nil{
		panic(err)
	}

	return reviews, nil
}



func UpdateReviewUpvoted(p graphql.ResolveParams) (i interface{}, e error){

	id := p.Args["id"].(int)
	upvote := p.Args["upvote"].(int)

	reviews, _:= models.UpdateReviewUpvote(upvote, uint(id))

	return reviews, nil
}

func UpdateReviewDownvoted(p graphql.ResolveParams) (i interface{}, e error){

	id := p.Args["id"].(int)
	upvote := p.Args["downvote"].(int)

	reviews, _:= models.UpdateReviewDownvote(upvote, uint(id))

	return reviews, nil
}

func InsertReview(p graphql.ResolveParams) (i interface{}, e error){

	token := p.Args["token"].(string)
	reviewdesc := p.Args["reviewdesc"].(string)
	gameid := p.Args["gameid"].(int)

	reviews, _:= models.InsertReview(token, reviewdesc, gameid)

	return reviews, nil
}