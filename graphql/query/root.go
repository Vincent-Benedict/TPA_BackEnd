package query


import (
	//res "github.com/Vincent-Benedict/TPA-Web/graphql/query/resolver"
	//typ "github.com/Vincent-Benedict/TPA-Web/graphql/type"
	"github.com/graphql-go/graphql"
	res "github.com/Vincent-Benedict/TPA-Web/graphql/query/resolver"
	typ "github.com/Vincent-Benedict/TPA-Web/graphql/type"
)

func GetRoot() *graphql.Object{
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{

			"games":{
				Type: graphql.NewList(typ.GetGameType()),
				Resolve: res.GetGames,
				Description: "To get All Games",
			},
			"game":{
				Type: typ.GetGameType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetGameById,
				Description: "To get singe Game",
			},
			"gamesByStatus":{
				Type: graphql.NewList(typ.GetGameType()),
				Args: graphql.FieldConfigArgument{
					"status": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetGamesByStatus,
				Description: "To get All Games By Status",
			},
			"gamesByStatusLimit":{
				Type: graphql.NewList(typ.GetGameType()),
				Args: graphql.FieldConfigArgument{
					"status": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetGamesByStatusLimit,
				Description: "To get All Games By Status Limit",
			},
			"gamesByStatusDiscountLimit":{
				Type: graphql.NewList(typ.GetGameType()),
				Args: graphql.FieldConfigArgument{
					"discount": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"status": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetGamesByStatusDiscountLimit,
				Description: "To get All Games By Discount Limit",
			},
			"gamesByNameLimit":{
				Type: graphql.NewList(typ.GetGameType()),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetGamesByNameLimit,
				Description: "To get All Games By Name Limit",
			},
			"gamesByName":{
				Type: graphql.NewList(typ.GetGameType()),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetGamesByName,
				Description: "To get All Games By Name",
			},
			"gamesByFilter":{
				Type: graphql.NewList(typ.GetGameType()),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"genre": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"category": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetGamesByFilter,
				Description: "To get All Games By Filter",
			},
			"reviewsByGameIdGreaterUpvotedAndLimit":{
				Type: graphql.NewList(typ.GetReviewType()),
				Args: graphql.FieldConfigArgument{
					"gameid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"reviewupvoted": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetReviewsByGameIdGreaterUpvotedAndLimit,
				Description: "To get All Reviews by gameid and greater upvoted",
			},
			"reviewsByGameIdEqualsUpvotedAndLimit":{
				Type: graphql.NewList(typ.GetReviewType()),
				Args: graphql.FieldConfigArgument{
					"gameid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"reviewupvoted": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetReviewsByGameIdEqualsUpvotedAndLimit,
				Description: "To get All Reviews by gameid and equals upvoted",
			},

			"getAllGamesInCart":{
				Type: graphql.NewList(typ.GetGameCartType()),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetGamesCart,
				Description: "To get All Games In Cart",
			},


			"getUser":{
				//Type: graphql.String,
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetUser,
				Description: "To Get User",
			},


			"getUserByToken":{
				//Type: graphql.String,
				Type: typ.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetUserByToken,
				Description: "To Get User From JWT Token",
			},
			"getUserById":{
				//Type: graphql.String,
				Type: typ.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetUserById,
				Description: "To Get User By Id",
			},

			"isaddedgametowishlist":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"gameid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.IsAddedGameToWishlist,
				Description: "To Get User From JWT Token",
			},

			"checkisaddedtoreviewlist":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"gameid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"reviewid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.CheckIsAdded,
				Description: "To Check Review is added to The List",
			},


			"checkuserishavingagame":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"gameid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.CheckUserIsHavingGame,
				Description: "To Check Review is added to The List",
			},

			"getallgameitem":{
				Type: graphql.NewList(typ.GetGameItemType()),
				Resolve: res.GetAllGameItem,
				Description: "To get All Game Items",
			},

			"getgameitembyid":{
				Type: typ.GetGameItemType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetGameItemById,
				Description: "To get a Single Game Item",
			},

			"getredeemcodebystring":{
				Type: typ.GetRedeemCodeType(),
				Args: graphql.FieldConfigArgument{
					"code": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetRedeemCodeByString,
				Description: "To get a Redeem Code",
			},

			"getusergamebyuserid":{
				Type: graphql.NewList(typ.GetUserGameType()),
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetUserGameByUserId,
				Description: "To Get User Game By Id",
			},

			"getfriendsbyuserid":{
				Type: graphql.NewList(typ.GetFriendType()),
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetFriendsByUserId,
				Description: "To Get Friend By UserId",
			},

			"getallcommentsbyuserid":{
				Type: graphql.NewList(typ.GetCommentType()),
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetCommentByUserId,
				Description: "To Get Comments By UserId",
			},

			"getidfromjwttoken":{
				Type : graphql.Int,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetIdFromJwtToken,
				Description: "To Get ID From JWT Token",
			},

			"getallchat": &graphql.Field{
				Type: graphql.NewList(typ.GetChatType()),
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"recipientid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetAllChat,
			},

		},
	})
}