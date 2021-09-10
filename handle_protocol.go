package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"gowebsocket/impl"
	"gowebsocket/platform"
	"time"
)

func SendResponseByServer(data []byte) {
	if len(data) > 0 {
		err := gConn.WriteMessage(data)
		if err != nil {
			gConn.Close()
		}
	}
}

func BoardCaseMsg(data []byte, title string) {
	fmt.Println("gConnectMax:", gConnectMax, "gConnArray:", gConnArray)
	if gConnectMax > 0 {
		for i := 0; i < int(gConnectMax); i++ {
			x, y := gConnArray[i].RemoteAddr()
			fmt.Println("send ", title, " data to :", x, y)
			gConnArray[i].WriteMessage(data)
		}
	}
}
func getDataByProtocol(conn *impl.Connection, len int32, gatetype int32, msgtype int32, uid int32, protoData []byte) {
	data := []byte{}
	if gatetype == int32(1000) && msgtype == 1 {
		data = SendGatewayLogin()
	}
	if gatetype == 1000 && msgtype == 9 {
		data = SendNetworkPingPong()
	}
	if gatetype == 1100 && msgtype == 9 {
		data = SendPlayerBalanceData()
	}
	if gatetype == 1100 && msgtype == 3 {
		//data = SendUserAttri()
	}
	if gatetype == 1100 && msgtype == 15 {
		data = SendHorseRaceLamp()
	}
	if gatetype == 1101 && msgtype == 1 {
		gameKindRequest := &platform.GameKindRequest{}
		proto.Unmarshal(protoData, gameKindRequest)
		gameKind := int32(gameKindRequest.GameKind)
		data = SendGameRoomList(gameKind)
	}
	if gatetype == 1101 && msgtype == 5 {
		data = SendMatchTable()
		callByTimeID("GAME_TABLE_INFO_TIMER", 10*time.Second, func() {
			SendResponseByServer(SendMatchTableResult())
		})
	}
	if gatetype == 4002 && msgtype == 1 {
		data = SendTeenPattiTableStatus()
		BoardCaseMsg(data, "SendTeenPattiTableStatus")
		return
	}
	if gatetype == 4002 && msgtype == 3 {
		data = SendTeenPattiPlayerReady()
		BoardCaseMsg(data, "SendTeenPattiPlayerReady")
		callByTimeID("GAME_BEGIN_CLOCK_TIMER", 1*time.Second, func() {
			data = SendGameStartClock()
			BoardCaseMsg(data, "SendGameStartClock")
		})
		callByTimeID("GAME_BEGIN_CLOCK_TIMER", 3*time.Second, func() {
			data = SendPlayersCardData()
			BoardCaseMsg(data, "SendPlayersCardData")
		})
		callByTimeID("GAME_ACTION_NOTIFY_TIMER", 3*time.Second, func() {
			data = SendPlayersAtions(1)
			BoardCaseMsg(data, "SendPlayersAtions")
		})
	}
	if gatetype == 4002 && msgtype == 7 {

	}
	if gatetype == 4002 && msgtype == 13 {
		//data = SendGameBet(protoData)

	}
	if gatetype == 4002 && msgtype == 15 {
		//data = SendPlayerLookCard()

	}
	if gatetype == 4002 && msgtype == 17 {

	}
	if gatetype == 4002 && msgtype == 19 {

	}
	if gatetype == 4002 && msgtype == 21 {

	}
	if gatetype == 4002 && msgtype == 23 {

	}
	if gatetype == 4002 && msgtype == 25 {

	}
	if gatetype == 4002 && msgtype == 60 {

	}
	if gatetype == 4002 && msgtype == 62 {
		data = SendPlayerChat()

	}
	if gatetype == 4002 && msgtype == 100 {

	}
	err := conn.WriteMessage(data) //发送数据
	if err != nil {                //错误处理
		conn.Close()
	}
}
