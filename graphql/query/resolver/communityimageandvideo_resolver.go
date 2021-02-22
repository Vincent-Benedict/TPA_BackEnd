package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetCommunityimageandvideo(p graphql.ResolveParams)(i interface{}, e error){

	object, _ := models.GetCommunityImageAndVideo()

	return object, nil
}

func GetCommunityImageAndVideoComment(p graphql.ResolveParams)(i interface{}, e error){

	contentid := p.Args["contentid"].(int)
	offset := p.Args["offset"].(int)
	limit := p.Args["limit"].(int)

	object, _ := models.GetCommunityImageAndVideoComment(contentid, offset, limit)

	return object, nil
}

func GetAllCommunityImageAndVideoComment(p graphql.ResolveParams)(i interface{}, e error){

	contentid := p.Args["contentid"].(int)

	object, _ := models.GetAllCommunityImageAndVideoComment(contentid)

	return object, nil
}

func InsertCommunityImageAndVideoComment(p graphql.ResolveParams)(i interface{}, e error){

	token := p.Args["token"].(string)
	contentid := p.Args["contentid"].(int)
	comment := p.Args["comment"].(string)

	object, _ := models.InsertCommunityImageAndVideoComment(token, contentid, comment)

	return object, nil
}

func InsertCommunityImageAndVideoLike(p graphql.ResolveParams)(i interface{}, e error){

	contentid := p.Args["contentid"].(int)

	object, _ := models.InsertCommunityImageAndVideoLike(contentid)

	return object, nil
}

func InsertCommunityImageAndVideoDislike(p graphql.ResolveParams)(i interface{}, e error){

	contentid := p.Args["contentid"].(int)
	object, _ := models.InsertCommunityImageAndVideoDislike(contentid)

	return object, nil
}