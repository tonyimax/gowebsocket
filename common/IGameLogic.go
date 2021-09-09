package common

var (
	GAME_PLAYER_COUNT     uint8    //游戏人数 2~5人
	GAME_CARD_COUNT       uint8    //扑克数量 52牌 (一副牌去掉大小王)
	GAME_BASE_SCORE       uint64   //游戏底注
	GAME_CARD_DATA        []byte   //牌的数据 52牌 默认是每人3张
	GAME_PLAYER_CARD      [][]byte //玩家的牌
	GAME_CARD_TYPE        []uint32 //牌型
	GAME_ROLES            []uint32 //游戏规则
	GAME_PLAY_TYPE        uint8    //玩法
	GAME_PLAY_BLIND_TIMES uint8    //每位玩家,盲下注次数最多4次,GAME_PLAY_BLIND_TIMES大于4的时候PlayerBetType由BLIND变成CHAAL
	GAME_PLAY_CHAAL_LIMIT uint64   //看牌下注额度固定等于房间底注128倍,达到最大额度时,看牌下注的玩家固定变为只能下固定额度，额度等于封顶值
	GAME_BET_POOL_LIMIT   []uint64 //奖池最高额度,当总奖池达到最大额度时，所有用户强制进行 show 牌比牌
	GAME_CARD_DEALER      int32    //发牌者
)

//游戏逻辑公共接口
type IGameLogic interface {
	Init()       //构造
	Destroy()    //消毁
	GameStart()  //游戏开始
	GameEnd()    //游戏结束
	SendCards()  //发牌
	GameBet()    //下注
	SeeCard()    //看牌
	ShowCard()   //开牌
	PackCard()   //弃牌
	GameWindUp() //游戏结算
}
