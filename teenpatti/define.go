package teenpatti

//游戏定义
type GameDefine struct {
	//在这里初始化游戏定义
}

//牌花色枚举
type CardColor uint8

//牌花色
const (
	SPADE   CardColor = 0 //黑桃
	HEART   CardColor = 1 //红桃
	DIAMOND CardColor = 2 //梅花
	CLUB    CardColor = 3 //方块
)

//牌型枚举
type CardType uint8

//三条-->同花顺-->顺子-->同花-->对子-->高牌
const (
	SET           CardType = 0 //三条
	PURE_SEQUENCE CardType = 1 //同花顺
	SEQUENCE      CardType = 2 //顺子
	COLOR         CardType = 3 //同花
	PAIR          CardType = 4 //对子
	HIGH_CARD     CardType = 5 //高牌
)

//----------------------------------------

//玩法枚举
type GamePlayType uint8

//玩法定义
const (
	CLASSIC GamePlayType = 0 //经典玩法
)

//----------------------------------------

//动作枚举
type PlayerAction uint8

//玩家动作
const (
	SEE  PlayerAction = 0 //看牌
	SHOW PlayerAction = 1 //开牌
	PACK PlayerAction = 2 //弃牌-投入本局奖池内的金额不会退回
)

//---------------------------------------

//下注标识枚举
type PlayerBetType uint8

//下注标识
const (
	BLIND PlayerBetType = 0 //盲下注   ([当前下注]的1倍或者2倍)
	CHAAL PlayerBetType = 1 //看牌下注 ([当前下注]的2倍或者4倍)
)

//---------------------------------------

//开牌枚举
type ShowCardType uint8

const (
	BLIND_SHOW ShowCardType = 0 //盲玩家-开牌
	CHAAL_SHOW ShowCardType = 1 //看牌玩家-开牌
	SIDE_SHOW  ShowCardType = 2 //单开-自己看牌了,上家也看牌了,自己提出跟自己的上家单独进行比牌 (开牌两玩家都已看牌) ,自己支付【当前叫注】的2倍进奖池才能跟上家比牌,拒绝比牌提出请求的用户仍然会支付Side Show的下注
)

//-----------------------------------------------

//游戏桌子信息
type GameTableInfo struct {
	TableID     uint32            //桌号
	PlayerInfo  []*GamePlayerInfo //桌子上的玩家
	PlayerCards [][]byte          //玩家的牌
}

//游戏玩家信息
type GamePlayerInfo struct {
	UserID  uint32 //用户ID
	Coin    uint32 //金币
	Diamond uint32 //钻石
	Nick    string //昵称
	Head    string //头像
}

//游戏玩家信息
type GamePlayerCardInfo struct {
	UserID uint32  //用户ID
	Cards  [3]byte //扑克
}

//牌局记录信息
type GameRecordInfo struct {
	RecordID          uint32   //记录ID
	TableID           uint32   //桌号
	TableUsers        []uint32 //桌子上的用户
	TableUserDraws    []uint32 //用户本局下注金额
	TableUserWinCoins []int32  //用户本局赢得金额
	TableUserCards    [][]byte //玩家的牌
	Head              string   //头像
}
