package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"gowebsocket/platform"
	"gowebsocket/teenpatti"
)

func Init() {
	GLoginUsers = make(map[string]uint32, 2)
	//游戏级别的基础定义
	gameLevelData = &platform.GameLevelDesc{
		LevelId:       0,  //房间类型ID
		CurrencyKind:  0,  //当前游戏房间类型-> 0：未知，1：金币，2: 练习
		CurrencyLimit: 0,  //进入限制（这个逻辑客户端不需要判定，只显示，服务器会下发，免得以后要修改逻辑）
		LevelName:     "", //级别名字，例如（中级房）
		UserCount:     0,  //这个级别有多少个用户在玩
		TaxPermillage: 0,  //抽水比率（千分比）
		GameId:        0,  //room number
	}
	//每一个级别的数据描叙
	roomInfo = &platform.TeepattiLevelDesc{
		GameLevel:     &platform.GameLevelDesc{}, //游戏级别的基础定义
		Blind:         0,                         //最大盲注次数
		SingleMaxBet:  0,                         //单人最大下注金额限制
		TableMaxBet:   0,                         //桌子最大下注金额限制
		TimesFeeRatio: 0,                         //台费系数
		Fantasytime:   0,                         //是否支持ft，0：不支持，1：支持
	}
	//匹配成功通知
	tableInfo = platform.MatchOKResponse{
		Result:    0, //結果      0：成功：100：用户余额不够；101：申请的游戏已经过期；102：账户被限制，不能匹配；103：服务器正在维护中:105：人数太多
		GameType:  0, //这个值就是动态ServerType，后续我们就正式开始游戏了
		TableId:   0, //桌号
		GameKind:  0, //游戏种类
		GameLevel: 0, //级别
	}

	//桌子上的用户
	tableUser = &platform.GameUser{
		Uid:      0,  //用户ID
		RealUser: 0,  //是否真实玩家
		Coin:     0,  //金币
		UserNick: "", //昵称
		UserHead: "", //头像
	}
	//匹配通知
	matchResponse = &platform.MatchResponse{
		Result:      0, //0：成功申请匹配；1：成功取消匹配；100：余额不够；101：目前已经在匹配队列中，不能重复匹配；102：不在匹配队列中，无法取消
		MaxTime:     0, //最大匹配时间：单位秒
		AverageTime: 0, //平均匹配时间：单位秒
	}
	//游戏类型
	gameKindResponse = platform.GameKindResponse{
		GameKind:       []platform.GameKind{},           //游戏类型 platform.GameKind_GAME_KIND_TEEPATTI
		TeepattiLevels: []*platform.TeepattiLevelDesc{}, //游戏房间明细 roomInfo
	}
	//用户金币结构
	balanceInfo = platform.GetPlayerBalanceResponse{
		Result:       0, //请求结果 0:成功
		Balance:      0, //账户余额
		BalanceWins:  0, //提现余额
		Partices:     0, //练习币余额
		GameCurrency: platform.CurrencyKind_CK_INVALID,
	}
	//用户属性
	userAttrib = platform.UserAttri{
		UserId: 0,  //用户ID
		Nick:   "", //昵称
		Head:   "", //头像
	}
	//登陆结果
	loginResponse = platform.LoginResponse{
		Result: 0,
		UserId: 0,
	}
	//座位属性
	chair = &teenpatti.ChairStatus{
		BUser:      0,                    //此座位是否有人：0：没有人；1：有人
		BGame:      0,                    //此座位是否参与本局游戏：0：不参与；1：参与（不参与的情况可能是刚进入游戏，或者没有钱站起等等）
		BDrop:      0,                    //此座位用户是否drop了：0：没有drop，正常游戏中；1：主动drop；2：超时drop；3：sideshow PK失败；4：show失败；5：达到封顶后强行结算失败
		ChairIndex: 0,                    //座位号：0/1/2/3/4/5
		Cards:      []uint32{0, 0, 0},    //座位上的牌
		PT:         0,                    //牌型 此座位是否有人：0：没有人；1：有人
		Score:      0,                    //牌分
		User:       &platform.GameUser{}, //座位上的用户
		BSee:       0,                    //用户是否看牌：0没有看；1看牌
		Bet:        0,                    //用户的当前局总计投注
		LastBet:    0,                    //用户最近一次投注
		Bet_A:      0,                    //用户的当前局总计投注：来自存款账户
		Bet_B:      0,                    //用户的当前局总计投注：来自可提现账户
	}
	//桌子的全部信息
	tableStatus = &teenpatti.MSG_C_GET_TABLE_STATUS_RESP{
		LevelDesc:         &platform.TeepattiLevelDesc{},   //teepatti每一个级别的数据描叙
		GamePhase:         teenpatti.GamePhase_PHS_INVALID, //游戏进度状态: 0:无效 1:资源加载(10秒) 2:等待用户确认(3秒) 3:发牌 4:提示用户操作 5:下注中 6:单开 7:结算 8:继续下一局(5秒)
		Charis:            []*teenpatti.ChairStatus{},      //座位属性
		SelfIndex:         0,                               //自己的座位
		GameID:            "",                              //本局牌局ID
		Dealer:            0,                               //庄家位置
		TotalCurrency:     0,                               //总共下注
		CurrentRoundValue: 0,                               //当前轮注（以盲为标准）
		CurrentRoundAct:   0,                               //当前轮次的实际叫分人
	}
	//發牌機構
	dealcard = &teenpatti.MSG_C_GAME_DEALCARDS_RESP{
		TableId: 0,                          //桌号
		GameID:  "",                         //本局牌局ID
		Charis:  []*teenpatti.ChairStatus{}, //座位上参与的人数，如果是2人局，则会有2个数据，如果是6人局，有6个数据
		Dealer:  0,                          //本局dealer（庄家）
	}
	//提示用戶操作
	gameNotice = &teenpatti.MSG_C_GAME_NOTICE_RESP{
		ChairIndex:        0, //座位编号
		BSelf:             0, //是否是本人  0不是  1是
		UserCurrency:      0, //当前操作人的货币
		CurrentRoundValue: 0, //当前轮注
		BSee:              0, //用户是否看过牌  0没有看过牌，1已经看过牌
		BBet:              0, //是否可以跟注
		BDouble:           0, //是否可以加注
		BShow:             0, //是否可以show
		BSideshow:         0, //是否可以sideshow
		TimeLeft:          0, //操作时间
		BUpdate:           0, //本次操作是否是因为外部条件变化而引发的重置，默认为0，例如其他用户看牌了，此值就会是1;2:表示是自己看牌
	}
	//看牌結構
	gameSeeCard = &teenpatti.MSG_C_GAME_SEE_RESP{
		ChairIndex: 65535,
		Cards:      []uint32{}, //0x02B, 0x2C, 0x2D
		PokerType:  -1,         //teenpatti.PokersType_PT_Set
	}
	//下注结构
	gameBet = &teenpatti.MSG_C_GAME_BET_RESP{
		ChairIndex:    0, //座位编号
		Result:        0, //0：正确；103：用户余额不够；其他错误
		TotalCurrency: 0, //所有用户总共下注
		BSee:          0, //用户是否看过牌  0没有看过牌，1已经看过牌
		BetOdd:        0, //0：放弃；1：跟注；2：加注；
		BetValue:      0, //下注的值
		UserBetValue:  0, //用户累计下注
		UserCurrency:  0, //下注后，用户的剩下的钱
	}

	//结算数据
	settles = make([]*teenpatti.UserSettle, GamePlayerCount)
	for i := 0; i < GamePlayerCount; i++ {
		settle = &teenpatti.UserSettle{
			ChairIndex:   uint32(i), //座位编号
			BGame:        0,         //用户是否参与
			BDrop:        0,         //用户是否drop，0：没有drop，正常游戏中；1：主动drop；2：超时drop；3：sideshow PK失败；4：show失败；5：达到封顶后强行结算失败
			Pokers:       []uint32{},
			PokersTypes:  -1, //teenpatti.PokersType_PT_Pure_Sequence
			WinCurrency:  0,
			UserCurrency: 0,
		}
		settles[i] = settle
	}
	//结算结构
	gameSettle = &teenpatti.MSG_C_GAME_SETTLE_RESP{
		UsersSettle: settles,
		WinIndex:    0, //胜利的位置
		WinReason:   0, //胜利的原因：1：drop后剩一个人；2：牌型比较胜；3：相同牌后手胜；4：相同牌封顶Dealer下家胜
		IsMax:       0, //是否达到了最大值：0没有1达到
	}

	gameEnd = &teenpatti.MSG_C_MATCH_FINISH_RESP{
		TableId: 0, //桌子ID
		Reason:  0, //解散原因
	}

	userAttribArray = platform.MSG_GET_USER_ATTRI_RESP{
		UserAttris: []*platform.UserAttri{},
	}
}

func buildGameRoomData(gameKind int32) proto.Message {
	gameLevelData.LevelId = 5
	gameLevelData.CurrencyKind = platform.CurrencyKind_CK_Money
	gameLevelData.CurrencyLimit = 100
	gameLevelData.LevelName = "Level I"
	gameLevelData.UserCount = 888
	gameLevelData.TaxPermillage = 0
	gameLevelData.GameId = 0

	roomInfo.GameLevel = gameLevelData
	roomInfo.Blind = 10
	roomInfo.SingleMaxBet = 1280
	roomInfo.TableMaxBet = 10240
	roomInfo.TimesFeeRatio = 10
	roomInfo.Fantasytime = 0

	gameKindResponse.GameKind = []platform.GameKind{platform.GameKind_GAME_KIND_TEEPATTI}
	gameKindResponse.TeepattiLevels = []*platform.TeepattiLevelDesc{roomInfo}

	return &gameKindResponse
}

func buildUserInfo(uid uint32) proto.Message {
	users := map[uint32]*platform.UserAttri{}
	users[573232] = &platform.UserAttri{UserId: 573232, Head: "111", Nick: "U573232"}
	users[602684] = &platform.UserAttri{UserId: 602684, Head: "118", Nick: "U602684"}
	userAttribArray.UserAttris = []*platform.UserAttri{users[uid]}
	return &userAttribArray
}

func buildBalanceData(uid uint32) proto.Message {
	balanceInfo.Balance = 888888
	balanceInfo.Result = 0
	balanceInfo.BalanceWins = 10000
	balanceInfo.GameCurrency = platform.CurrencyKind_CK_Money
	balanceInfo.Partices = 1000000
	return &balanceInfo
}

func buildTableInfo(tableId int) proto.Message {
	tableInfo.Result = 0
	tableInfo.TableId = 587643
	tableInfo.GameKind = 4
	tableInfo.GameType = 4002
	tableInfo.GameLevel = 5
	return &tableInfo
}

func buildLoginResult() proto.Message {
	loginResponse.Result = 0
	if gConnectMax > 1 {
		loginResponse.UserId = 602684
	} else {
		loginResponse.UserId = 573232
	}
	fmt.Println(loginResponse.UserId, "is Login!")
	for i := 0; i < int(gConnectMax); i++ {
		_, y := gConnArray[i].RemoteAddr()
		GLoginUsers[y] = loginResponse.UserId
	}
	return &loginResponse
}
