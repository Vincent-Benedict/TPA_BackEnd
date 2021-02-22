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

			"updateuserabout":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"username": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"realname": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"customurl": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"country": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"summary": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.UpdateUserAbout,
				Description: "To Update User About",
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

			"insertreviewcommentbyreviewid":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"reviewid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"comment": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.InsertReviewCommentByReviewId,
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

			"insertuserframe":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"frame": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.InsertUserFrame,
				Description: "To Insert Frame",
			},
			"insertuserbackground":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"background": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.InsertUserBackground,
				Description: "To Insert Background",
			},

			"insertuserminibackground":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"minibackground": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.InsertUserMiniBackground,
				Description: "To Insert Mini Background",
			},

			"insertusergamebyid":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"gameid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.InsertUserGameById,
				Description: "To Insert usergame by id",
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

			"decreasepoint":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"point": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.DecreasePoint,
				Description: "To Decrease Point in User",
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

			"insertmarketactivity":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"itemid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"activity": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.InsertMarketRecentActivity,
				Description: "To Insert MarketActivity",
			},

			"insertfriendrequest":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"friendid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.InsertFriendRequest,
				Description: "To Insert Friend Request",
			},

			"insertfriendrequestignored":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"friendid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.InsertFriendRequestIgnored,
				Description: "To Insert Friend Request Ignored",
			},

			"insertdiscussion":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"gameid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"discussiontitle": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"discussioncontent": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.InsertDiscussion,
				Description: "To Insert Discussion",
			},

			"insertdiscussioncomment":{
				Type: typ.GetDiscussionCommentType(),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"discussionid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"comment": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.InsertDiscussionComment,
				Description: "To Insert Discussion Comment",
			},


			"insertchatnotification":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"friendid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.InsertChatNotification,
				Description: "To Insert Chat Notification",
			},

			"insertgiftnotification":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"friendid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.InsertGiftNotification,
				Description: "To Insert Gift Notification",
			},

			"insertitemnotification":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"itemid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.InsertItemNotification,
				Description: "To Insert Item Notification",
			},

			"insertcommentnotification":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"friendid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.InsertCommentNotification,
				Description: "To Insert Comment Notification",
			},

			"insertfriend":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"friendid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.InsertFriend,
				Description: "To Insert Friend",
			},

			"insertcomment":{
				Type: typ.GetCommentType(),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"friendid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"commentdesc": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.InsertComment,
				Description: "To Insert Comment",
			},

			"insertreport":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"reportedid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"reportreason": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.InsertReport,
				Description: "To Insert Report",
			},

			"insertrecentactivity":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"activity": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.InsertRecentActivity,
				Description: "To Insert Recent Activity",
			},
			"insertrecentactivitybyid":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"activity": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.InsertRecentActivityById,
				Description: "To Insert Recent Activity By Id",
			},

			"insertuser":{
				Type: typ.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"country": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.InsertUser,
				Description: "To Insert User",
			},


			"insertcommunityimageandvideocomment":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"contentid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"comment": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.InsertCommunityImageAndVideoComment,
				Description: "To Insert community image and video comment",
			},

			"insertselllisting":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"itemid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"quantity": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.InsertSellListing,
				Description: "To Insert Sell Listing",
			},

			"insertbuylisting":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"itemid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"quantity": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.InsertBuyListing,
				Description: "To Insert Buy Listing",
			},

			"buysolditem":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"itemid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"quantity": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.BuySoldItem,
				Description: "To Buy Sold Item",
			},





			"deletefriendrequest":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.DeleteFriendRequest,
				Description: "To Delete Friend Request",
			},

			"deletefriendrequest2":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"friendid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.DeleteFriendRequest2,
				Description: "To Delete Friend Request",
			},

			"deletechatnotification":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"friendid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.DeleteChatNotification,
				Description: "To Delete Chat Notification",
			},

			"deletegiftnotification":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"friendid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.DeleteGiftNotification,
				Description: "To Delete Chat Notification",
			},

			"deletecommentnotification":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"friendid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.DeleteCommentNotification,
				Description: "To Delete Comment Notification",
			},

			"deleteitemnotification":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.DeleteItemNotification,
				Description: "To Delete Item Notification",
			},

			"deletegamefromwishlist":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"gameid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.DeleteGameFromWishlist,
				Description: "To Delete Game From Wishlist",
			},


			"updateuseravatar":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"avatar": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"frame": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.UpdateUserAvatar,
				Description: "To Update User Avatar",
			},
			"saveavatar":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"encoded": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"filename": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.SaveAvatar,
				Description: "To Sava Avatar",
			},


			"updateuserbackground":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"background": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.UpdateUserBackground,
				Description: "To Update User Bacgkround",
			},

			"updateuserminiprofile":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"miniprofile": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.UpdateUserMiniProfile,
				Description: "To Update User Mini Profile",
			},

			"updateusertheme":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"theme": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.UpdateUserTheme,
				Description: "To Update User Theme",
			},

			"updateuserbadge":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"badgename": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.UpdateUserBadge,
				Description: "To Update User Badge",
			},


			"insertcommunityimageandvideolike":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"contentid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.InsertCommunityImageAndVideoLike,
				Description: "To Insert Like Community Image And Video",
			},

			"insertcommunityimageandvideodislike":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"contentid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.InsertCommunityImageAndVideoDislike,
				Description: "To Insert Dislike Community Image And Video",
			},


			// ADMIN
			"insertgame":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"encoded": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"filename": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"tags": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.InsertGame,
				Description: "To Insert Game",
			},

			"updategame":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"encoded": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"filename": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"tags": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.UpdateGame,
				Description: "To Update Game",
			},

			"deletegame":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.DeleteGame,
				Description: "To Delete Game",
			},

			"insertupdatedeletepromo":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"promo": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.InsertUpdateDeletePromo,
				Description: "To Insert Update Delete Promo",
			},


			"insertunsuspendrequest":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.InsertUnsuspendRequest,
				Description: "To Insert Unsuspend Request",
			},

			"deleteunsuspendrequest":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.DeleteUnsuspendRequest,
				Description: "To Delete Unsuspend Request",
			},

			"deletereportbyreportedid":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"reportedid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.DeleteReportByReportedId,
				Description: "To Delete Report By reported id",
			},

		},
	})
}
