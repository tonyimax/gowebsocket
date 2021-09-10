package main

import (
	"google.golang.org/protobuf/proto"
	"gowebsocket/impl"
	"gowebsocket/platform"
	"gowebsocket/teenpatti"
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
		if gConnectMax > 0 {
			for i := 0; i < int(gConnectMax); i++ {
				gConnArray[i].WriteMessage(data)
			}
		}
		return
	}
	if gatetype == 4002 && msgtype == 3 {
		data = SendTeenPattiPlayerReady()

	}
	if gatetype == 4002 && msgtype == 7 {

	}
	if gatetype == 4002 && msgtype == 13 {
		data = SendGameBet(protoData)

	}
	if gatetype == 4002 && msgtype == 15 {
		data = SendPlayerLookCard()

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

	//主动推送倒计时信息
	if gatetype == int32(platform.ServerType_SERVER_TYPE_TEEPATTI_+2) &&
		msgtype == int32(teenpatti.TeenpattiCmd_CMD_C_MATCH_READY_REQ) &&
		err == nil {

		//SendResponseByServer(SendGameStartClock())
		//SendResponseByServer(SendPlayersCardData())
		//callByTimeID("GAME_ACTION_NOTIFY_TIMER",
		//	3*time.Second,
		//	func() {
		//	SendResponseByServer(SendPlayersAtions(currentChair%2))
		//})
	}
}
