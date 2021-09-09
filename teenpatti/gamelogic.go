package teenpatti

import (
	"fmt"
	c "gowebsocket/common"
)

//游戏逻辑
type GameLogic struct {
	//定义额外的游戏中需要的信息IGameLogic没有定义的
}

func New() c.IGameLogic {
	return &GameLogic{} //返回实现化接口的对象
}

//构造
func (logic *GameLogic) Init() {
	fmt.Println("构造游戏逻辑--初始化")
	c.GAME_PLAYER_COUNT = 2 //游戏人数
	c.GAME_CARD_COUNT = 52  //扑克数量
	c.GAME_CARD_DATA = []byte{
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, //方块
		0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, //梅花
		0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, //红桃
		0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4A, 0x4B, 0x4C, 0x4D, //黑桃
	}
	c.GAME_BASE_SCORE = 1000    //游戏底注
	c.GAME_PLAY_BLIND_TIMES = 4 //盲下注次数
	v := uint64(128 * c.GAME_BASE_SCORE)
	c.GAME_PLAY_CHAAL_LIMIT = v
	c.GAME_BET_POOL_LIMIT = []uint64{v, v, v, v}
	tip := fmt.Sprintf("游戏人数:%v,扑克数量%v,游戏底注:%v,盲下注次数:%v,看牌下注额度固:%v", c.GAME_PLAYER_COUNT, c.GAME_CARD_COUNT, c.GAME_BASE_SCORE, c.GAME_PLAY_BLIND_TIMES, c.GAME_PLAY_CHAAL_LIMIT)
	fmt.Println(tip, "\n牌的数据")
}

//消毁
func (logic *GameLogic) Destroy() {
	fmt.Println("消毁游戏逻辑")
}

//实现接口方法-游戏开始
func (logic *GameLogic) GameStart() {
	fmt.Println("游戏开始了")
}

//实现接口方法-游戏结束
func (logic *GameLogic) GameEnd() {
	fmt.Println("游戏结束了")
}

//实现接口方法-发牌
func (logic *GameLogic) SendCards() {
	//生成牌
	colorName := []string{"方块", "梅花", "红桃", "黑桃"}
	cardIndex := c.GetCards(uint32(c.GAME_PLAYER_COUNT)*3, uint32(c.GAME_CARD_COUNT), int(c.GAME_PLAYER_COUNT))
	for i := 0; i < len(cardIndex); i++ {
		value, color := c.GetPokerValueAndColor(c.GAME_CARD_DATA[cardIndex[i]])
		v := fmt.Sprintf("[%v%v],", colorName[color-1], value)
		fmt.Print(v)
	}
	fmt.Println("")

	userIds := []int32{573232, 489866}
	playerCards := make([]GamePlayerCardInfo, c.GAME_PLAYER_COUNT)
	for i := 0; i < len(playerCards); i++ {
		playerCards[i].UserID = uint32(userIds[i])
		for j := 0; j < 3; j++ {
			playerCards[i].Cards[j] = c.GAME_CARD_DATA[cardIndex[j+(3*i)]]
			value, color := c.GetPokerValueAndColor(c.GAME_CARD_DATA[cardIndex[j+(3*i)]])
			v := fmt.Sprintf("[%v%v],", colorName[color-1], value)
			fmt.Print(v)
		}
		fmt.Println("UserID:", playerCards[i].UserID, "Cards", playerCards[i].Cards)
	}

}

//实现接口方法-下注
func (logic *GameLogic) GameBet() {
	fmt.Println("下注了")
}

//实现接口方法-看牌
func (logic *GameLogic) SeeCard() {
	fmt.Println("看牌了")
}

//实现接口方法-开牌
func (logic *GameLogic) ShowCard() {
	fmt.Println("开牌了")
}

//实现接口方法-弃牌
func (logic *GameLogic) PackCard() {
	fmt.Println("弃牌了")
}

//实现接口方法-游戏结算
func (logic *GameLogic) GameWindUp() {
	fmt.Println("游戏结算了")
}
