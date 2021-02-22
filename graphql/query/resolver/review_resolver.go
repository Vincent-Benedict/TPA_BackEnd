package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetAllReview(p graphql.ResolveParams) (i interface{}, e error){

	limit := p.Args["limit"].(int)

	reviews, err := models.GetAllReview(limit)

	if err!=nil{
		panic(err)
	}

	return reviews, nil
}

func GetReviewById(p graphql.ResolveParams) (i interface{}, e error){

	id := p.Args["id"].(int)

	reviews, err := models.GetReviewById(id)

	if err!=nil{
		panic(err)
	}

	return reviews, nil
}

func GetReviewsRecently(p graphql.ResolveParams) (i interface{}, e error){

	gameid := p.Args["gameid"].(int)
	limit := p.Args["limit"].(int)

	reviews, err := models.GetReviewRecently(gameid, limit)

	if err!=nil{
		panic(err)
	}

	return reviews, nil
}

func GetReviewsUpvoted(p graphql.ResolveParams) (i interface{}, e error){

	gameid := p.Args["gameid"].(int)
	limit := p.Args["limit"].(int)

	reviews, err := models.GetReviewUpvoted(gameid, limit)

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

func InsertReviewCommentByReviewId(p graphql.ResolveParams) (i interface{}, e error){

	token := p.Args["token"].(string)
	comment := p.Args["comment"].(string)
	reviewid := p.Args["reviewid"].(int)

	reviews, _:= models.InsertReviewCommentByReviewId(token, reviewid, comment)

	return reviews, nil
}