package main

import (
	"gowebsocket/platform"
	"gowebsocket/teenpatti"
)

func Init()  {
	isMatchOK = false
	isPacked=false
	isNotify=false
	currentChair=0
	dealer=0
	//游戏级别的基础定义
	gameLevelData = &platform.GameLevelDesc{
		LevelId:       5,                           //房间类型ID
		CurrencyKind:  platform.CurrencyKind_CK_Money, //当前游戏房间类型-> 0：未知，1：金币，2: 练习
		CurrencyLimit: 100,//进入限制（这个逻辑客户端不需要判定，只显示，服务器会下发，免得以后要修改逻辑）
		LevelName:     "Level I",//级别名字，例如（中级房）
		UserCount:     2626,//这个级别有多少个用户在玩
	}
	//房间信息
	roomInfo=&platform.TeepattiLevelDesc{
		GameLevel:     gameLevelData, //游戏级别的基础定义
		Blind:         10,         //最大盲注次数
		SingleMaxBet:  1280,       //单人最大下注金额限制
		TableMaxBet:   10240,      //桌子最大下注金额限制
		TimesFeeRatio: 10,         //teenpatti的台费系数
	}
	//fmt.Println("房间信息:",roomInfo)
	tableInfo = &platform.MatchOKResponse{
		Result:    0,
		GameType:  4002, //游戏主协议
		TableId:   587643,//桌号
		GameKind:  4,//游戏种类 4:TeenPatti
		GameLevel: 5,//房间级别
	}
	//fmt.Println("桌子信息:",tableInfo)
	//桌子上的用户
	tableUsers=[]*platform.GameUser{
		{
			Uid:      602684,        //用户ID
			RealUser: 0,             //是否真实玩家
			Coin:     8146000,        //金币
			UserNick: "U602684", //昵称
			UserHead: "118",         //头像
		},
		{
			Uid:      573232,        //用户ID
			RealUser: 0,             //是否真实玩家
			Coin:     4027000,        //金币
			UserNick: "U573232", //昵称
			UserHead: "111",         //头像
		},
	}
	//fmt.Println("桌子上的玩家:",tableUsers)

	matchResponse = &platform.MatchResponse{
		Result:      0,
		MaxTime:     10,
		AverageTime: 5,
	}
	//游戏类型
	gameKindResponse = &platform.GameKindResponse{
		GameKind: []platform.GameKind{platform.GameKind_GAME_KIND_TEEPATTI}, //游戏类型
		TeepattiLevels: []*platform.TeepattiLevelDesc{ roomInfo,},//游戏房间明细
	}
	//fmt.Println("游戏类型",gameKindResponse)

	//用户金币结构
	balanceInfo = &platform.GetPlayerBalanceResponse{
		Result:       0,         //请求结果 0:成功
		Balance:      99999,     //账户余额
		BalanceWins:  100,       //提现余额
		Partices:     500000000, //练习币余额
		GameCurrency: platform.CurrencyKind_CK_Money,
	}

	//用户属性
	userAttris = []*platform.UserAttri{
		{
			UserId: tableUsers[0].Uid,  //用户ID
			Nick:   tableUsers[0].UserNick, //昵称
			Head:  tableUsers[0].UserHead, //头像
		},
		{
			UserId: tableUsers[1].Uid,  //用户ID
			Nick:   tableUsers[1].UserNick, //昵称
			Head:  tableUsers[1].UserHead, //头像
		},
	}
	//登陆结果
	loginResponse = &platform.LoginResponse{
		Result:0,
		UserId:0,
	}

	//座位属性
	chairs = []*teenpatti.ChairStatus{
		{
			BUser:      1,                 //此座位是否有人：0：没有人；1：有人
			BGame:      0,                 //此座位是否参与本局游戏：0：不参与；1：参与（不参与的情况可能是刚进入游戏，或者没有钱站起等等）
			BDrop:      0,                 //此座位用户是否drop了：0：没有drop，正常游戏中；1：主动drop；2：超时drop；3：sideshow PK失败；4：show失败；5：达到封顶后强行结算失败
			ChairIndex: 0,                 //座位号：0/1/2/3/4/5
			Cards:      []uint32{0, 0, 0}, //座位上的牌
			PT:         0,                 //牌型 此座位是否有人：0：没有人；1：有人
			Score:      1000,                 //牌分
			//座位上的用户
			User: tableUsers[0],
			BSee:    0, //用户是否看牌：0没有看；1看牌
			Bet:     0, //用户的当前局总计投注
			LastBet: 0, //用户最近一次投注
			Bet_A:   0, //用户的当前局总计投注：来自存款账户
			Bet_B:   0, //用户的当前局总计投注：来自可提现账户
		},
		{
			BUser:      1,                 //此座位是否有人：0：没有人；1：有人
			BGame:      0,                 //此座位是否参与本局游戏：0：不参与；1：参与（不参与的情况可能是刚进入游戏，或者没有钱站起等等）
			BDrop:      0,                 //此座位用户是否drop了：0：没有drop，正常游戏中；1：主动drop；2：超时drop；3：sideshow PK失败；4：show失败；5：达到封顶后强行结算失败
			ChairIndex: 1,                 //座位号：0/1/2/3/4/5
			Cards:      []uint32{0, 0, 0}, //座位上的牌
			PT:         0,                 //牌型 此座位是否有人：0：没有人；1：有人
			Score:      1000,                 //牌分
			//座位上的用户
			User: tableUsers[1],
			BSee:    0, //用户是否看牌：0没有看；1看牌
			Bet:     0, //用户的当前局总计投注
			LastBet: 0, //用户最近一次投注
			Bet_A:   0, //用户的当前局总计投注：来自存款账户
			Bet_B:   0, //用户的当前局总计投注：来自可提现账户
		},
	}

	//桌子的全部信息
	tableStatus = &teenpatti.MSG_C_GET_TABLE_STATUS_RESP{
		LevelDesc: 		   roomInfo, //teepatti每一个级别的数据描叙
		GamePhase:         teenpatti.GamePhase_PHS_Match_Ready,//游戏进度状态: 0:无效 1:资源加载(10秒) 2:等待用户确认(3秒) 3:发牌 4:提示用户操作 5:下注中 6:单开 7:结算 8:继续下一局(5秒)
		Charis:            chairs, //座位属性
		SelfIndex:         0,      //自己的座位
		GameID:            "4",     //本局牌局ID
		Dealer:            1,      //庄家位置
		TotalCurrency:     0,      //总共下注
		CurrentRoundValue: 10,     //当前轮注（以盲为标准）
		CurrentRoundAct:   0,      //当前轮次的实际叫分人
	}

	dealcard = &teenpatti.MSG_C_GAME_DEALCARDS_RESP{
		TableId: 587643,     //桌号
		GameID:  "4", //本局牌局ID
		Charis:  chairs,   //座位上参与的人数，如果是2人局，则会有2个数据，如果是6人局，有6个数据
		Dealer:  0,        //本局dealer（庄家）
	}

	gameNotice = &teenpatti.MSG_C_GAME_NOTICE_RESP{
		ChairIndex:        0, //座位编号
		BSelf:             1, //是否是本人  0不是  1是
		UserCurrency:      1000, //当前操作人的货币
		CurrentRoundValue: 1000, //当前轮注
		BSee:              0, //用户是否看过牌  0没有看过牌，1已经看过牌
		BBet:              1, //是否可以跟注
		BDouble:           1, //是否可以加注
		BShow:             1, //是否可以show
		BSideshow:         1, //是否可以sideshow
		TimeLeft:          10, //操作时间
		BUpdate:           0, //本次操作是否是因为外部条件变化而引发的重置，默认为0，例如其他用户看牌了，此值就会是1;2:表示是自己看牌
	}

	gameSeeCard = &teenpatti.MSG_C_GAME_SEE_RESP{
		ChairIndex:0,
		Cards: []uint32{0x02B,0x2C,0x2D},
		PokerType:teenpatti.PokersType_PT_Set,
	}
	//下注结构
	gameBet = &teenpatti.MSG_C_GAME_BET_RESP{
		ChairIndex:currentChair,//座位编号
		Result:0,//0：正确；103：用户余额不够；其他错误
		TotalCurrency:0,//所有用户总共下注
		BSee:0,//用户是否看过牌  0没有看过牌，1已经看过牌
		BetOdd:1,//0：放弃；1：跟注；2：加注；
		BetValue:1000,//下注的值
		UserBetValue:1000,//用户累计下注
		UserCurrency: int32(tableUsers[currentChair].Coin-1000),//下注后，用户的剩下的钱
	}

	pokerInfo = [][]uint32{{0x2B,0x2C,0x2D},{0x3B,0x3C,0x3D}} //牌数据
	//结算数据
	settles = make([]*teenpatti.UserSettle,2)
	for i := 0; i < 2; i++ {
		settle = &teenpatti.UserSettle{
			ChairIndex:uint32(i),//座位编号
			BGame:1,//用户是否参与
			BDrop:1,//用户是否drop，0：没有drop，正常游戏中；1：主动drop；2：超时drop；3：sideshow PK失败；4：show失败；5：达到封顶后强行结算失败
			Pokers:pokerInfo[i],
			PokersTypes:teenpatti.PokersType_PT_Pure_Sequence,
			WinCurrency:-1000,
			UserCurrency:int32(tableUsers[i].Coin-1000),
		}
		settles[i] = settle
	}
	//结算结构
	gameSettle = &teenpatti.MSG_C_GAME_SETTLE_RESP{
		UsersSettle:settles,
		WinIndex:1,//胜利的位置
		WinReason:1,//胜利的原因：1：drop后剩一个人；2：牌型比较胜；3：相同牌后手胜；4：相同牌封顶Dealer下家胜
		IsMax:0,//是否达到了最大值：0没有1达到
	}

	gameEnd =&teenpatti.MSG_C_MATCH_FINISH_RESP{
		TableId:587643,//桌子ID
		Reason:1,//解散原因
	}
}
