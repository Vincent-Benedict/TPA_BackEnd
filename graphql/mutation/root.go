package mutation

import (
	res "github.com/Vincent-Benedict/TPA-Web/graphql/query/resolver"
	typ "github.com/Vincent-Benedict/TPA-Web/graphql/type"

	//res "github.com/Vincent-Benedict/TPA-Web/graphql/query/resolver"
	//typ "github.com/Vincent-Benedict/TPA-Web/graphql/type"
	"github.com/graphql-go/graphql"
)

func GetRoot() *graphql.Object{
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"insertgametocart":{
				Type: typ.GetGameCartType(),
				Args: graphql.FieldConfigArgument{
					"gameid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"discount": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.InsertGamesToCart,
				Description: "To Insert Games To Cart",
			},

			"deletegamefromcartbyid":{
				Type: typ.GetGameCartType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.DeleteGameFromCartById,
				Description: "To Delete Games From Cart",
			},

			"deletegamefromcartbyid2":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.DeleteGameFromCartById2,
				Description: "To Delete Games From Cart boolean",
			},


			"insertgametowishlist":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"gameid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.InsertGamesToWishlist,
				Description: "To Insert Games To Wishlist",
			},


			"updatereviewupvoted":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"upvote": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.UpdateReviewUpvoted,
				Description: "To Update Review Upvoted",
			},
			"updatereviewdownvoted":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"downvote": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.UpdateReviewDownvoted,
				Description: "To Update Review Downvoted",
			},


			"insertreviewlist":{
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
				Resolve: res.InsertReviewList,
				Description: "To Insert Review to The List",
			},

			"insertreview":{
				Type: typ.GetReviewType(),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"reviewdesc": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"gameid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.InsertReview,
				Description: "To Insert Review",
			},

			"insertusergame":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"gameid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.InsertUserGame,
				Description: "To Insert usergame",
			},

			"deleteallgamefromcart":{
				Type: graphql.Boolean,
				Resolve: res.DeleteAllGameFromCart,
				Description: "Delete all game from cart",
			},

			"decreasebalance":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"balance": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.DecreaseBalance,
				Description: "To Decrease Balance in User",
			},

			"increasebalance":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"balance": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.IncreaseBalance,
				Description: "To Increase Balance in User",
			},

			"insertchat":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"message": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"recipientid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.InsertChat,
				Description: "To Insert Chat",
			},

		},
	})
}
