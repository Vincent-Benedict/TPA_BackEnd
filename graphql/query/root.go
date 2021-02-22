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
			"getgamesbyfeatureandrecommend":{
				Type: graphql.NewList(typ.GetGameType()),
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetGamesFeatureAndRecommended,
				Description: "To get Game By Featured And Recommended",
			},
			"getgamesspecialoffer":{
				Type: graphql.NewList(typ.GetGameType()),
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetGamesSpecialOffer,
				Description: "To get Game By Special Offer",
			},
			"getgamesnewandtrending":{
				Type: graphql.NewList(typ.GetGameType()),
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetGamesNewAndTrending,
				Description: "To get Game By New And Trending",
			},
			"gamesoffsetlimit":{
				Type: graphql.NewList(typ.GetGameType()),
				Args: graphql.FieldConfigArgument{
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetGamesOffsetLimit,
				Description: "To get game by offset limit",
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
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
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

			"getallreview":{
				Type: graphql.NewList(typ.GetReviewType()),
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetAllReview,
				Description: "To get All Reviews",
			},

			"getreviewbyid":{
				Type: typ.GetReviewType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetReviewById,
				Description: "To get Review by Id",
			},

			"reviewsUpvoted":{
				Type: graphql.NewList(typ.GetReviewType()),
				Args: graphql.FieldConfigArgument{
					"gameid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetReviewsUpvoted,
				Description: "To get All Reviews upvoted",
			},
			"reviewsRecently":{
				Type: graphql.NewList(typ.GetReviewType()),
				Args: graphql.FieldConfigArgument{
					"gameid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetReviewsRecently,
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

			"getusers":{
				Type: graphql.NewList(typ.GetUserType()),
				Resolve: res.GetUsers,
				Description: "To Get All User",
			},
			"getuseroffsetlimit":{
				Type: graphql.NewList(typ.GetUserType()),
				Args: graphql.FieldConfigArgument{
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetUserOffsetLimit,
				Description: "To Get User Offset Limit",
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
			"checkuserbyusername":{
				//Type: graphql.String,
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.CheckUserByUsername,
				Description: "To Check User By Username",
			},
			"getidfromcustomurl":{
				//Type: graphql.String,
				Type: graphql.Int,
				Args: graphql.FieldConfigArgument{
					"customurl": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetIdFromCustomUrl,
				Description: "To Get Id from Custom Url",
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

			"getgameitemoffsetlimit":{
				Type: graphql.NewList(typ.GetGameItemType()),
				Args: graphql.FieldConfigArgument{
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetGameItemOffsetLimit,
				Description: "To get Game Items by offset and limit",
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
			"getusergamebyuseridlimit":{
				Type: graphql.NewList(typ.GetUserGameType()),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetUserGameByUserIdLimit,
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

			"getrecentactivity":{
				Type : graphql.NewList(typ.GetRecentActivityType()),
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetRecentActivity,
				Description: "To Get All RecentActivity",
			},

			"getrecentactivityoffsetlimit":{
				Type : graphql.NewList(typ.GetRecentActivityType()),
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetRecentActivityOffsetLimit,
				Description: "To Get All RecentActivity Offset Limit",
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

			"getallmarketactivity": &graphql.Field{
				Type: graphql.NewList(typ.GetMarketRecentActivityTypeType()),
				Args: graphql.FieldConfigArgument{
					"itemid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetAllMarketActivity,
			},

			"isfriend":&graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"friendid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.IsFriend,
			},
			"getFriendRequest":&graphql.Field{
				Type: graphql.NewList(typ.GetFriendRequestType()),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetFriendRequestByUserId,
				Description: "To Get FriendRequest From JWT Token",
			},
			"getFriendRequestSent":&graphql.Field{
				Type: graphql.NewList(typ.GetFriendRequestType()),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetFriendRequestByUserIdSent,
				Description: "To Get FriendRequest Sent From JWT Token",
			},

			"getFriendRequestIgnored":&graphql.Field{
				Type: graphql.NewList(typ.GetFriendRequestIgnoredType()),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetFriendRequestIgnoredByUserId,
				Description: "To Get FriendRequestIgnored From JWT Token",
			},

			"getreportbytoken":&graphql.Field{
				Type: graphql.NewList(typ.GetReportType()),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetReportByToken,
				Description: "To Get Report From JWT Token",
			},

			"sendemail":&graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.SendEmail,
				Description: "To Send Email",
			},

			"sendemailtransaction":&graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"firstname": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"message": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"sentiment": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"signature": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.SendEmailTransaction,
				Description: "To Send Email Transaction",
			},

			"sendemailwishlist":&graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"gamename": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.SendEmailWishlist,
				Description: "To Send Email Wishlist",
			},

			"getusertheme":&graphql.Field{
				Type: graphql.NewList(typ.GetUserThemeType()),
				Resolve: res.GetUserTheme,
				Description: "To Get All User Theme",
			},

			"getuserthemebyid":&graphql.Field{
				Type: typ.GetUserThemeType(),
				Resolve: res.GetUserThemeById,
				Description: "To Get User Theme by Id",
			},

			"getchatnotification":&graphql.Field{
				Type: graphql.NewList(typ.GetChatNotificationType()),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetChatNotificationByToken,
				Description: "To Get Chat Notification By Token",
			},

			"getgiftnotification":&graphql.Field{
				Type: graphql.NewList(typ.GetGiftNotificationType()),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetGiftNotificationByToken,
				Description: "To Get Gift Notification By Token",
			},

			"getcommentnotification":&graphql.Field{
				Type: graphql.NewList(typ.GetCommentNotificationType()),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetCommentNotificationByToken,
				Description: "To Get Comment Notification By Token",
			},

			"getitemnotification":&graphql.Field{
				Type: graphql.NewList(typ.GetItemNotificationType()),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetItemNotificationByToken,
				Description: "To Get Item Notification By Token",
			},

			"getfriendsbyusername":&graphql.Field{
				Type: graphql.NewList(typ.GetFriendType()),
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"username": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetFriendsByUsername,
				Description: "To Get Friends By Username",
			},

			"getuserbyfriendcode":&graphql.Field{
				Type: typ.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"friendcode": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetUserByFriendCode,
				Description: "To Get User By FriendCode",
			},


			"getallshopavatar":&graphql.Field{
				Type: graphql.NewList(typ.GetShopAvatarType()),
				Resolve: res.GetAllShopAvatar,
				Description: "To Get All Shop Avatar",
			},
			"getallshopframe":&graphql.Field{
				Type: graphql.NewList(typ.GetShopFrameType()),
				Resolve: res.GetAllShopFrame,
				Description: "To Get All Shop Frame",
			},
			"getallshopbackground":&graphql.Field{
				Type: graphql.NewList(typ.GetShopBackgroundType()),
				Resolve: res.GetAllShopBackground,
				Description: "To Get All Shop Background",
			},
			"getallshopminibackground":&graphql.Field{
				Type: graphql.NewList(typ.GetShopMiniBackgroundType()),
				Resolve: res.GetAllShopMiniBackground,
				Description: "To Get All Shop Mini Background",
			},
			"getallshopchatsticker":&graphql.Field{
				Type: graphql.NewList(typ.GetChatStickerType()),
				Resolve: res.GetAllShopChatSticker,
				Description: "To Get All Shop Chat Sticker",
			},

			"getcommunityimageandvideo":&graphql.Field{
				Type: graphql.NewList(typ.GetCommunityImageAndVideoType()),
				Resolve: res.GetCommunityimageandvideo,
				Description: "To Get All Community Image And Video",
			},

			"getcommunityimageandvideocomment":&graphql.Field{
				Type: graphql.NewList(typ.GetCommunityImageAndVideoCommentType()),
				Args: graphql.FieldConfigArgument{
					"contentid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetCommunityImageAndVideoComment,
				Description: "To Get Community Image And Video Comment",
			},

			"getallcommunityimageandvideocomment":&graphql.Field{
				Type: graphql.NewList(typ.GetCommunityImageAndVideoCommentType()),
				Args: graphql.FieldConfigArgument{
					"contentid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetAllCommunityImageAndVideoComment,
				Description: "To Get All Community Image And Video Comment",
			},

			"getallreviewcommentbyreviewid":&graphql.Field{
				Type: graphql.NewList(typ.GetReviewCommentType()),
				Args: graphql.FieldConfigArgument{
					"reviewid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetAllReviewCommentByReviewId,
				Description: "To Get All Review Comment By Review Id",
			},

			"getreviewcommentbyreviewidoffsetlimit":&graphql.Field{
				Type: graphql.NewList(typ.GetReviewCommentType()),
				Args: graphql.FieldConfigArgument{
					"reviewid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetReviewCommentByReviewIdOffsetLimit,
				Description: "To Get Review Comment By Review Id Offset Limit",
			},

			"getalldiscussion":&graphql.Field{
				Type: graphql.NewList(typ.GetDiscussionType()),
				Resolve: res.GetAllDiscussion,
				Description: "To Get All Discussion",
			},

			"getdiscussionswithgamename":&graphql.Field{
				Type: graphql.NewList(typ.GetDiscussionType()),
				Args: graphql.FieldConfigArgument{
					"gamename": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetDiscussionsWithGameName,
				Description: "To Get Discussions With Game Name",
			},
			"getdiscussionswithid":&graphql.Field{
				Type: typ.GetDiscussionType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetDiscussionWithId,
				Description: "To Get Discussion With Id",
			},

			"getdiscussionscommentwithdiscussionid":&graphql.Field{
				Type: graphql.NewList(typ.GetDiscussionCommentType()),
				Args: graphql.FieldConfigArgument{
					"discussionid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetDiscussionCommentByDiscussionId,
				Description: "To Get Discussion Comment With Discussion Id",
			},

			"getalluseritemwithgameid":&graphql.Field{
				Type: graphql.NewList(typ.GetUserItemType()),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"gameid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetAllUserItemWithGameId,
				Description: "To Get All User Item With Game Id",
			},
			"getuseritemwithgameidoffsetlimit":&graphql.Field{
				Type: graphql.NewList(typ.GetUserItemType()),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"gameid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetUserItemWithGameIdOffsetLimit,
				Description: "To Get User Item With Game Id Offset Limit",
			},
			"getuseritemwithgameidoffsetlimitbygamename":&graphql.Field{
				Type: graphql.NewList(typ.GetUserItemType()),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"gameid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"gamename": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetUserItemWithGameIdOffsetLimitByName,
				Description: "To Get User Item With Game Id Offset Limit by Game Name",
			},

			"getuseritemwithgameidbygamename":&graphql.Field{
				Type: graphql.NewList(typ.GetUserItemType()),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"gameid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"gamename": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetUserItemWithGameIdByName,
				Description: "To Get User Item With Game Id by Game Name",
			},




			"getbuylistingquantityandpricebyitemid":&graphql.Field{
				Type: graphql.NewList(typ.GetBuyQuantityPriceType()),
				Args: graphql.FieldConfigArgument{
					"itemid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetBuyListingQuantityAndPriceByItemId,
				Description: "To Get Buy Listing Quantity And Price by Item Id",
			},

			"getselllistingquantityandpricebyitemid":&graphql.Field{
				Type: graphql.NewList(typ.GetSellQuantityPriceType()),
				Args: graphql.FieldConfigArgument{
					"itemid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetSellListingQuantityAndPriceByItemId,
				Description: "To Get Sell Listing Quantity And Price by Item Id",
			},

			"getbuylistingbyuserid":&graphql.Field{
				Type: graphql.NewList(typ.GetBuyListingType()),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"itemid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetBuyListingByUserId,
				Description: "To Get Buy Listing By Userid",
			},

			"getselllistingbyuserid":&graphql.Field{
				Type: graphql.NewList(typ.GetSellListingType()),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"itemid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetSellListingByUserId,
				Description: "To Get Sell Listing By Userid",
			},

			"checkifthereisselllisting":&graphql.Field{
				Type: graphql.Int,
				Args: graphql.FieldConfigArgument{
					"itemid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"quantity": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.CheckIfThereIsSellListing,
				Description: "To Check If There Is Sell Listing",
			},

			"getalluserbadgebyuserid":&graphql.Field{
				Type: graphql.NewList(typ.GetUserBadgeType()),
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetAllUserBadgeByUserId,
				Description: "To Get All User Badge by User ID",
			},
			"getuserbadgebyid":&graphql.Field{
				Type: typ.GetUserBadgeType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetUserBadgeById,
				Description: "To Get User Badge By Id",
			},
			"getbadgecardbyid":&graphql.Field{
				Type: graphql.NewList(typ.GetBadgeCardType()),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetBadgeCardById,
				Description: "To Get Badge Card By Id",
			},

			//ADMIN
			"getadmin":&graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetAdmin,
				Description: "To Get Admin JWT Token",
			},

			"getallpromogames":&graphql.Field{
				Type: graphql.NewList(typ.GetGameType()),
				Args: graphql.FieldConfigArgument{
					"auth": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetAllPromoGames,
				Description: "To Get All Promo Games",
			},
			"getpromogamesoffsetlimit":&graphql.Field{
				Type: graphql.NewList(typ.GetGameType()),
				Args: graphql.FieldConfigArgument{
					"auth": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetPromoGameOffsetLimit,
				Description: "To Get All Promo Games Offset Limit",
			},


			"getreportbyreportedid":&graphql.Field{
				Type: graphql.NewList(typ.GetReportType()),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetReportByReportedId,
				Description: "To Get Report By Report ID",
			},

			"getallunsuspendrequest":&graphql.Field{
				Type: graphql.NewList(typ.GetUnsuspendRequestType()),
				Resolve: res.GetAllUnsupendRequest,
				Description: "To Get All Unsuspend Request",
			},

		},
	})
}