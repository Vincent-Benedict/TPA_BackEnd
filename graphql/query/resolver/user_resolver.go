package resolver

import (
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

func GetUsers(p graphql.ResolveParams) (i interface{}, e error){

	user, err := models.GetUsers()

	if err!=nil{
		panic(err)
	}

	return user, nil
}

func GetUserOffsetLimit(p graphql.ResolveParams) (i interface{}, e error){

	offset := p.Args["offset"].(int)
	limit := p.Args["limit"].(int)

	user, err := models.GetUserOffsetLimit(offset, limit)

	if err!=nil{
		panic(err)
	}

	return user, nil
}

func GetUser(p graphql.ResolveParams) (i interface{}, e error){

	username := p.Args["username"].(string)
	password := p.Args["password"].(string)

	user, err := models.GetUser(username, password)

	if err!=nil{
		panic(err)
	}

	return user, nil
}

func GetUserByFriendCode(p graphql.ResolveParams) (i interface{}, e error){

	friendcode := p.Args["friendcode"].(string)

	user, err := models.GetUserByFriendCode(friendcode)

	if err!=nil{
		panic(err)
	}

	return user, nil
}

func GetIdFromJwtToken(p graphql.ResolveParams)(i interface{}, e error){
	token := p.Args["token"].(string)

	id, _ := models.GetIdFromJwtToken(token);

	return id, nil
}

func GetUserById(p graphql.ResolveParams)(i interface{}, e error){

	id := p.Args["id"].(int)

	user, err := models.GetUserById(id)

	if err!=nil{
		panic(err)
	}

	return user, nil
}

func CheckUserByUsername(p graphql.ResolveParams)(i interface{}, e error){

	username := p.Args["username"].(string)

	user, err := models.CheckUserByUsername(username)

	if err!=nil{
		panic(err)
	}

	return user, nil
}

func GetUserByToken(p graphql.ResolveParams)(i interface{}, e error){

	token := p.Args["token"].(string)

	user, err := models.GetUserByToken(token)

	if err!=nil{
		panic(err)
	}

	return user, nil
}

func DecreaseBalance(p graphql.ResolveParams)(i interface{}, e error){
	token := p.Args["token"].(string)
	balance := p.Args["balance"].(int)

	user, _ := models.DecreaseBalance(token, balance)

	return user, nil
}

func DecreasePoint(p graphql.ResolveParams)(i interface{}, e error){
	token := p.Args["token"].(string)
	point := p.Args["point"].(int)

	user, _ := models.DecreasePoint(token, point)

	return user, nil
}

func IncreaseBalance(p graphql.ResolveParams)(i interface{}, e error){
	token := p.Args["token"].(string)
	balance := p.Args["balance"].(int)

	user, _ := models.IncreaseBalance(token, balance)

	return user, nil
}

func InsertUser(p graphql.ResolveParams)(i interface{}, e error){

	username := p.Args["username"].(string)
	password := p.Args["password"].(string)
	email := p.Args["email"].(string)
	country := p.Args["country"].(string)

	user, _ := models.InsertUser(username, password, email, country)

	return user, nil
}

func GetIdFromCustomUrl(p graphql.ResolveParams)(i interface{}, e error){

	customurl := p.Args["customurl"].(string)

	user, _ := models.GetIdFromCustomUrl(customurl)

	return user, nil
}



// UPDATE PROFILE
func UpdateUserAbout(p graphql.ResolveParams)(i interface{}, e error){

	id := p.Args["id"].(int)
	username := p.Args["username"].(string)
	realname := p.Args["realname"].(string)
	customurl := p.Args["customurl"].(string)
	country := p.Args["country"].(string)
	summary := p.Args["summary"].(string)

	user, _ := models.UpdateUserAbout(id, username, realname, customurl, country, summary)

	return user, nil
}

func UpdateUserAvatar(p graphql.ResolveParams)(i interface{}, e error){

	id := p.Args["id"].(int)
	avatar := p.Args["avatar"].(string)
	frame := p.Args["frame"].(string)

	user, _ := models.UpdateUserAvatar(id, avatar, frame)

	return user, nil
}

func UpdateUserBackground(p graphql.ResolveParams)(i interface{}, e error){

	id := p.Args["id"].(int)
	background := p.Args["background"].(string)

	user, _ := models.UpdateUserBackground(id, background)

	return user, nil
}

func UpdateUserMiniProfile(p graphql.ResolveParams)(i interface{}, e error){

	id := p.Args["id"].(int)
	miniprofile := p.Args["miniprofile"].(string)

	user, _ := models.UpdateUserMiniProfile(id, miniprofile)

	return user, nil
}

func UpdateUserTheme(p graphql.ResolveParams)(i interface{}, e error){

	id := p.Args["id"].(int)
	theme := p.Args["theme"].(string)

	user, _ := models.UpdateUserTheme(id, theme)

	return user, nil
}

func UpdateUserBadge(p graphql.ResolveParams)(i interface{}, e error){

	id := p.Args["id"].(int)
	badgename := p.Args["badgename"].(string)

	user, _ := models.UpdateUserBadge(id, badgename)

	return user, nil
}



func SaveAvatar(p graphql.ResolveParams)(i interface{}, e error){

	encoded := p.Args["encoded"].(string)
	filename := p.Args["filename"].(string)

	user, _ := models.SaveAvatar(encoded, filename)

	return user, nil
}


