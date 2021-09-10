package main

import (
	"gowebsocket/platform"
	"gowebsocket/teenpatti"
)

var (
	TP              = int32(platform.ServerType_SERVER_TYPE_TEEPATTI_ + 2)
	MATCH           = int32(platform.ServerType_SERVER_TYPE_MATCH)
	CMD             = int32(platform.ServerType_SERVER_TYPE_COMMON)
	GW              = int32(platform.ServerType_SERVER_TYPE_GATEWAY)
	GamePlayerCount = 5
	GLoginUsers     map[string]uint32
)

var (
	gameLevelData    *platform.GameLevelDesc
	roomInfo         *platform.TeepattiLevelDesc            //房间信息
	tableUser        *platform.GameUser                     //桌子上用户信息
	tableInfo        platform.MatchOKResponse               //桌子信息
	matchResponse    *platform.MatchResponse                //通知匹配
	gameKindResponse platform.GameKindResponse              //
	balanceInfo      platform.GetPlayerBalanceResponse      //用户余额
	userAttrib       platform.UserAttri                     //用户属性
	userAttribArray  platform.MSG_GET_USER_ATTRI_RESP       //用户属性结构
	loginResponse    platform.LoginResponse                 //登陆结果
	chair            *teenpatti.ChairStatus                 //座位属性
	tableStatus      *teenpatti.MSG_C_GET_TABLE_STATUS_RESP //桌子状态
	dealcard         *teenpatti.MSG_C_GAME_DEALCARDS_RESP   //发牌结构
	gameNotice       *teenpatti.MSG_C_GAME_NOTICE_RESP      //提示玩家操作
	gameSeeCard      *teenpatti.MSG_C_GAME_SEE_RESP         //看牌结构
	gameBet          *teenpatti.MSG_C_GAME_BET_RESP         //下注结构
	settle           *teenpatti.UserSettle                  //单个用户结算
	settles          []*teenpatti.UserSettle                //用户结算
	pokerInfo        [][]uint32                             //扑克信息
	gameSettle       *teenpatti.MSG_C_GAME_SETTLE_RESP      //弃牌信息
	gameEnd          *teenpatti.MSG_C_MATCH_FINISH_RESP     //解散桌子
)
