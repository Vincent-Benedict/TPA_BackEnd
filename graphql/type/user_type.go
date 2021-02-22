package _type

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"github.com/Vincent-Benedict/TPA-Web/models"
	"github.com/graphql-go/graphql"
)

var userType *graphql.Object

func GetUserType() *graphql.Object{

	if userType == nil{
		userType = graphql.NewObject(graphql.ObjectConfig{
			Name:"userType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"username": &graphql.Field{
					Type: graphql.String,
				},
				"password": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"country": &graphql.Field{
					Type: graphql.String,
				},
				"realname": &graphql.Field{
					Type: graphql.String,
				},
				"balance": &graphql.Field{
					Type: graphql.Int,
				},
				"avatar": &graphql.Field{
					Type: graphql.String,
				},
				"level": &graphql.Field{
					Type: graphql.Int,
				},
				"summary": &graphql.Field{
					Type: graphql.String,
				},
				"background": &graphql.Field{
					Type: graphql.String,
				},
				"miniprofile": &graphql.Field{
					Type: graphql.String,
				},
				"status": &graphql.Field{
					Type: graphql.String,
				},
				"customurl": &graphql.Field{
					Type: graphql.String,
				},
				"frame": &graphql.Field{
					Type: graphql.String,
				},
				"friendcode": &graphql.Field{
					Type: graphql.String,
				},
				"point": &graphql.Field{
					Type: graphql.Int,
				},
				"badge": &graphql.Field{
					Type: GetUserBadgeType(),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						user := params.Source.(models.User)

						var badgeUser models.UserBadge
						db.Where("badge_img = ?", user.Badge).First(&badgeUser)

						return badgeUser, nil
					},
				},
				"usertheme": &graphql.Field{
					Type: GetUserThemeType(),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						theme := params.Source.(models.User)

						var themeUser models.UserTheme
						db.Where("color = ?", theme.Theme).First(&themeUser)

						return themeUser, nil
					},
				},
				"wishlist": &graphql.Field{
					Type: graphql.NewList(GetWishlistGameType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						user := params.Source.(models.User)

						var wishlistGame []models.WishlistGame
						db.Where("user_id = ?", user.ID).Find(&wishlistGame)

						return wishlistGame, nil
					},
				},
				"userframe": &graphql.Field{
					Type: graphql.NewList(GetUserFrameType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						user := params.Source.(models.User)

						var userframe []models.UserFrame
						db.Where("user_id = ?", user.ID).Find(&userframe)

						return userframe, nil
					},
				},

				"userbadge": &graphql.Field{
					Type: graphql.NewList(GetUserBadgeType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						user := params.Source.(models.User)

						var userBadge []models.UserBadge
						db.Where("user_id = ?", user.ID).Find(&userBadge)

						return userBadge, nil
					},
				},

				"userbackground": &graphql.Field{
					Type: graphql.NewList(GetUserBackgroundType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						user := params.Source.(models.User)

						var userbackground []models.UserBackground
						db.Where("user_id = ?", user.ID).Find(&userbackground)

						return userbackground, nil
					},
				},

				"userminiprofile": &graphql.Field{
					Type: graphql.NewList(GetUserMiniProfileType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						db, _:= database.Connect()
						defer db.Close()
						user := params.Source.(models.User)

						var userminiprofile []models.UserMiniProfile
						db.Where("user_id = ?", user.ID).Find(&userminiprofile)

						return userminiprofile, nil
					},
				},
			},
		})
	}

	return userType
}

