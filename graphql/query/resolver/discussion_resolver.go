package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetDiscussionCommentByDiscussionId(p graphql.ResolveParams)(i interface{}, e error){

	discussionid := p.Args["discussionid"].(int)

	comment, _ := models.GetDiscussionCommentByDiscussionId(discussionid)

	return comment, nil
}

func InsertDiscussionComment(p graphql.ResolveParams)(i interface{}, e error){

	token := p.Args["token"].(string)
	discussionid := p.Args["discussionid"].(int)
	comment := p.Args["comment"].(string)

	comments, _ := models.InsertDiscussionComment(token, discussionid, comment)

	return comments, nil
}

func GetAllDiscussion(p graphql.ResolveParams)(i interface{}, e error){

	discuss, _ := models.GetAllDiscussion()

	return discuss, nil
}

func InsertDiscussion(p graphql.ResolveParams)(i interface{}, e error){

	token := p.Args["token"].(string)
	gameid := p.Args["gameid"].(int)
	discussiontitle := p.Args["discussiontitle"].(string)
	discussioncontent := p.Args["discussioncontent"].(string)

	discuss, _ := models.InsertDiscussion(token, gameid, discussiontitle, discussioncontent)

	return discuss, nil
}

func GetDiscussionsWithGameName(p graphql.ResolveParams)(i interface{}, e error){

	gameName := p.Args["gamename"].(string)

	discuss, _ := models.GetDiscussionsWithGameName(gameName)

	return discuss, nil
}

func GetDiscussionWithId(p graphql.ResolveParams)(i interface{}, e error){

	id := p.Args["id"].(int)

	discuss, _ := models.GetDiscussionWithId(id)

	return discuss, nil
}