package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetGames(p graphql.ResolveParams) (i interface{}, e error){
	games, err := models.GetGames()

	if err!=nil{
		panic(err)
	}

	return games, nil
}

func GetGamesOffsetLimit(p graphql.ResolveParams) (i interface{}, e error){
	offset:= p.Args["offset"].(int)
	limit:= p.Args["limit"].(int)

	games, err := models.GetGamesOffsetLimit(offset, limit)

	if err!=nil{
		panic(err)
	}

	return games, nil
}

func GetGameById(p graphql.ResolveParams) (i interface{}, e error){

	id:= p.Args["id"].(int)

	game, _:= models.GetGame(id)

	return game, nil
}

func GetGamesByStatus(p graphql.ResolveParams) (i interface{}, e error){

	status := p.Args["status"].(string)

	games, err := models.GetGamesByStatus(status)

	if err!=nil{
		panic(err)
	}

	return games, nil
}

func GetGamesByStatusLimit(p graphql.ResolveParams) (i interface{}, e error){

	status := p.Args["status"].(string)
	limit := p.Args["limit"].(int)

	games, err := models.GetGamesByStatusLimit(status,limit)

	if err!=nil{
		panic(err)
	}

	return games, nil
}

func GetGamesByStatusDiscountLimit(p graphql.ResolveParams) (i interface{}, e error){

	status := p.Args["status"].(string)
	discount := p.Args["discount"].(int)
	limit := p.Args["limit"].(int)

	games, err := models.GetGamesSpecialsSpecificDiscount(status, discount,limit)

	if err!=nil{
		panic(err)
	}

	return games, nil
}

func GetGamesByNameLimit(p graphql.ResolveParams) (i interface{}, e error){

	name := p.Args["name"].(string)
	limit := p.Args["limit"].(int)

	games, err := models.GetGamesByNameLimit(name,limit)

	if err!=nil{
		panic(err)
	}

	return games, nil
}

func GetGamesByName(p graphql.ResolveParams) (i interface{}, e error){

	name := p.Args["name"].(string)

	games, err := models.GetGamesByName(name)

	if err!=nil{
		panic(err)
	}

	return games, nil
}

func GetGamesByFilter(p graphql.ResolveParams) (i interface{}, e error){

	name := p.Args["name"].(string)
	price := p.Args["price"].(int)
	category := p.Args["category"].(string)
	genre := p.Args["genre"].(string)
	offset:= p.Args["offset"].(int)
	limit:= p.Args["limit"].(int)

	games, err := models.GetGamesByFilter(offset, limit, name, price, category, genre)

	if err!=nil{
		panic(err)
	}

	return games, nil
}

func GetGamesFeatureAndRecommended(p graphql.ResolveParams) (i interface{}, e error){
	limit:= p.Args["limit"].(int)

	games, _ := models.GetGamesByFeaturedAndRecommended(limit);

	return games, nil
}

func GetGamesSpecialOffer(p graphql.ResolveParams) (i interface{}, e error){
	limit:= p.Args["limit"].(int)

	games, _ := models.GetGamesSpecialOffers(limit)

	return games, nil
}

func GetGamesNewAndTrending(p graphql.ResolveParams) (i interface{}, e error){
	limit:= p.Args["limit"].(int)

	games, _ := models.GetGamesNewAndTrending(limit)

	return games, nil
}

func GetAllReviewCommentByReviewId(p graphql.ResolveParams) (i interface{}, e error){
	reviewid:= p.Args["reviewid"].(int)

	reviewcomment, _ := models.GetAllReviewCommentByReviewId(reviewid)

	return reviewcomment, nil
}

func GetReviewCommentByReviewIdOffsetLimit(p graphql.ResolveParams) (i interface{}, e error){
	reviewid:= p.Args["reviewid"].(int)
	offset:= p.Args["offset"].(int)
	limit:= p.Args["limit"].(int)

	reviewcomment, _ := models.GetReviewCommentByReviewIdOffsetLimit(reviewid, offset, limit)

	return reviewcomment, nil
}


// 1
func InsertGame(p graphql.ResolveParams) (i interface{}, e error){
	encoded:= p.Args["encoded"].(string)
	filename:= p.Args["filename"].(string)
	title:= p.Args["title"].(string)
	desc:= p.Args["description"].(string)
	price:= p.Args["price"].(int)
	tags:= p.Args["tags"].(string)

	reviewcomment, _ := models.InsertGame(encoded, filename, title, desc, price, tags)

	return reviewcomment, nil
}

func UpdateGame(p graphql.ResolveParams) (i interface{}, e error){
	id:= p.Args["id"].(int)
	encoded:= p.Args["encoded"].(string)
	filename:= p.Args["filename"].(string)
	title:= p.Args["title"].(string)
	desc:= p.Args["description"].(string)
	price:= p.Args["price"].(int)
	tags:= p.Args["tags"].(string)

	reviewcomment, _ := models.UpdateGame(id, encoded, filename, title, desc, price, tags)

	return reviewcomment, nil
}

func DeleteGame(p graphql.ResolveParams) (i interface{}, e error){
	id:= p.Args["id"].(int)

	bolean, _ := models.DeleteGame(id)

	return bolean, nil
}



// 2
func GetAllPromoGames(p graphql.ResolveParams) (i interface{}, e error){
	auth:= p.Args["auth"].(string)

	games, _ := models.GetAllPromoGames(auth)

	return games, nil
}

func GetPromoGameOffsetLimit(p graphql.ResolveParams) (i interface{}, e error){
	auth:= p.Args["auth"].(string)
	offset:= p.Args["offset"].(int)
	limit:= p.Args["limit"].(int)

	games, _ := models.GetPromoGameOffsetLimit(auth, offset, limit)

	return games, nil
}

func InsertUpdateDeletePromo(p graphql.ResolveParams) (i interface{}, e error){
	id:= p.Args["id"].(int)
	promo:= p.Args["promo"].(int)

	games, _ := models.InsertUpdateDeletePromo(id, promo)

	return games, nil
}



