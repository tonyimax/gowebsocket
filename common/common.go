package common

import (
	"fmt"
	"math/rand"
	"time"
)

//游戏公共业务
type GameCommon struct {
	//初始化游戏公共业务信息
}

func New() IBase {
	return &GameCommon{} //返回实现化接口的对象
}

//实现接口方法-构造
func (common *GameCommon) Init() {
	fmt.Println("构造游戏公共业务")
}

//实现接口方法-消毁
func (common *GameCommon) Destroy() {
	fmt.Println("消毁游戏公共业务")
}

//取扑克花色
func GetPokerColor(cardData byte) uint8 {
	return cardData / 16
}

//取扑克牌值
func GetPokerValue(cardData byte) uint8 {
	return cardData & 255 % 16
}

//取扑克牌值,花色
func GetPokerValueAndColor(cardData byte) (uint8, uint8) {
	return cardData & 255 % 16, cardData / 16
}

//生成随机数
func GetRandNum(theMaxNum uint32) int32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int31n(int32(theMaxNum))
}

//取不重复的牌
func GetCards(count uint32, cardMaxCount uint32, playerCount int) []byte {
	cards := make([]byte, count)
	for j := 0; j < playerCount; j++ {
		for i := 0; i < 3; i++ {
			v := GetRandNum(cardMaxCount)
			for k := 0; k < len(cards); k++ {
				if v == int32(cards[k]) {
					v = GetRandNum(cardMaxCount)
				}
			}
			cards[i+j*3] = byte(v)
		}
	}
	return cards
}

func GetNewGameID() uint32 {
	return uint32(GetRandNum(999999))
}

func GetNewUserID() uint32 {
	return uint32(GetRandNum(99999999))
}

func GetNewUserHead() uint32 {
	return uint32(GetRandNum(999))
}
