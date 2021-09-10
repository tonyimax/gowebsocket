package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"gowebsocket/platform"
	"gowebsocket/teenpatti"
	"time"
)

func packHallData(mCmd int32, sCmd int32, msg proto.Message) []byte {
	data, _ := proto.Marshal(msg)
	return packageData(mCmd, sCmd, data, int64(len(data)))
}

func packGameData(mCmd int32, sCmd int32, msg proto.Message, tableId uint32) []byte {
	byteData, _ := proto.Marshal(msg)
	tb := &teenpatti.TableData{
		TableId: tableId,
		Data:    byteData,
	}
	return packHallData(mCmd, sCmd, tb)
}

func SendGameEnd() []byte {
	sCmd := int32(teenpatti.TeenpattiCmd_CMD_C_MATCH_FINISH_RESP)
	return packGameData(TP, sCmd, gameEnd, tableInfo.TableId)
}

func SendPackedCard() []byte {
	sCmd := int32(teenpatti.TeenpattiCmd_CMD_C_GAME_SETTLE_RESP)
	return packGameData(TP, sCmd, gameSettle, tableInfo.TableId)
}

func SendGameBet(protoData []byte) []byte {
	tbl := &teenpatti.TableData{}
	proto.Unmarshal(protoData, tbl)
	msg := &teenpatti.MSG_C_GAME_BET_REQ{}
	proto.Unmarshal(tbl.Data, msg)
	sCmd := int32(teenpatti.TeenpattiCmd_CMD_C_GAME_BET_RESP)
	if gameBet.BetOdd == 0 {
		callByTimeID("GAME_PACKED_CARD_TIMER", 2*time.Second, func() {
			SendResponseByServer(SendPackedCard())
		})
		callByTimeID("GAME_END_TIMER", 10*time.Second, func() {
			SendResponseByServer(SendGameEnd())
		})
	}
	return packGameData(TP, sCmd, gameBet, tableInfo.TableId)
}

func SendPlayerChat() []byte {
	data := []byte{8, 251, 238, 35, 18, 13, 16, 1, 26, 9, 80, 108, 97, 121, 32, 70, 97, 115, 116}
	pData := packageData(TP,
		int32(teenpatti.TeenpattiCmd_CMD_C_CHAT_RESP),
		data, int64(len(data)))
	return pData
}

func SendPlayerLookCard() []byte {
	sCmd := int32(teenpatti.TeenpattiCmd_CMD_C_GAME_SEE_RESP)
	return packGameData(TP, sCmd, gameSeeCard, tableInfo.TableId)
}

func SendPlayersAtions(iChair uint32) []byte {
	sCmd := int32(teenpatti.TeenpattiCmd_CMD_C_GAME_NOTICE_RESP)
	return packGameData(TP, sCmd, gameNotice, tableInfo.TableId)
}

func SendPlayersCardData() []byte {
	sCmd := int32(teenpatti.TeenpattiCmd_CMD_C_GAME_DEALCARDS_RESP)
	return packGameData(TP, sCmd, dealcard, tableInfo.TableId)
}

func SendGameStartClock() []byte {
	resp := &teenpatti.MSG_C_GAME_READY_3_RESP{Times: 5}
	sCmd := int32(teenpatti.TeenpattiCmd_CMD_C_GAME_READY_3_RESP)
	return packGameData(TP, sCmd, resp, tableInfo.TableId)
}

func SendTeenPattiPlayerReady() []byte {
	resp := &teenpatti.MSG_C_COMMON_RESP{Result: 0}
	sCmd := int32(teenpatti.TeenpattiCmd_CMD_C_MATCH_READY_RESP)
	return packGameData(TP, sCmd, resp, tableInfo.TableId)
}

func SendTeenPattiTableStatus() []byte {
	sCmd := int32(teenpatti.TeenpattiCmd_CMD_C_GET_TABLE_STATUS_RESP)
	for i := 0; i < int(gConnectMax); i++ {
		tableStatus.Charis[i].User = tableUser
		tableStatus.Charis[i].ChairIndex = uint32(i)
	}
	return packGameData(TP, sCmd, tableStatus, tableInfo.TableId)
}

func SendMatchTableResult() []byte {
	sCmd := int32(platform.ServerMatchCmd_CMD_MATCH_OK_RESP)
	return packHallData(MATCH, sCmd, tableInfo)
}

func SendMatchTable() []byte {
	sCmd := int32(platform.ServerMatchCmd_CMD_MATCH_RESP)
	return packHallData(MATCH, sCmd, matchResponse)
}

func SendGameRoomList(gameKind int32) []byte {
	sCmd := int32(platform.ServerMatchCmd_CMD_GET_GAME_KIND_RESP)
	return packHallData(MATCH, sCmd, buildGameRoomData(gameKind))
}

func SendHorseRaceLamp() []byte {
	data := []byte{10, 114, 10, 112, 80, 108, 97, 121, 101, 114, 32, 42, 42, 42, 52, 56, 54, 55, 32, 119, 111, 110, 32, 60, 99, 111, 108, 111, 114, 61, 35, 100, 97, 99, 48, 56, 57, 62, 82, 115, 46, 49, 50, 48, 48, 46, 48, 60, 47, 99, 62, 32, 105, 110, 32, 60, 99, 111, 108, 111, 114, 61, 35, 53, 53, 69, 48, 70, 70, 62, 82, 97, 112, 105, 100, 32, 84, 101, 101, 110, 32, 80, 97, 116, 116, 105, 60, 47, 99, 111, 108, 111, 114, 62, 44, 32, 109, 111, 110, 101, 121, 32, 99, 111, 109, 101, 115, 32, 115, 111, 32, 102, 97, 115, 116, 33, 10, 114, 10, 112, 80, 108, 97, 121, 101, 114, 32, 42, 42, 42, 50, 51, 48, 53, 32, 119, 111, 110, 32, 60, 99, 111, 108, 111, 114, 61, 35, 100, 97, 99, 48, 56, 57, 62, 82, 115, 46, 51, 54, 48, 48, 46, 48, 60, 47, 99, 62, 32, 105, 110, 32, 60, 99, 111, 108, 111, 114, 61, 35, 53, 53, 69, 48, 70, 70, 62, 82, 97, 112, 105, 100, 32, 84, 101, 101, 110, 32, 80, 97, 116, 116, 105, 60, 47, 99, 111, 108, 111, 114, 62, 44, 32, 109, 111, 110, 101, 121, 32, 99, 111, 109, 101, 115, 32, 115, 111, 32, 102, 97, 115, 116, 33, 10, 97, 10, 95, 83, 111, 32, 108, 117, 99, 107, 121, 33, 32, 80, 108, 97, 121, 101, 114, 32, 42, 42, 42, 56, 49, 56, 48, 32, 119, 111, 110, 32, 60, 99, 111, 108, 111, 114, 61, 35, 100, 97, 99, 48, 56, 57, 62, 82, 115, 46, 49, 57, 48, 48, 46, 48, 60, 47, 99, 62, 32, 105, 110, 32, 60, 99, 111, 108, 111, 114, 61, 35, 54, 68, 66, 57, 56, 48, 62, 65, 110, 100, 97, 114, 32, 66, 97, 104, 97, 114, 60, 47, 99, 111, 108, 111, 114, 62, 10, 97, 10, 95, 83, 111, 32, 108, 117, 99, 107, 121, 33, 32, 80, 108, 97, 121, 101, 114, 32, 42, 42, 42, 56, 49, 56, 48, 32, 119, 111, 110, 32, 60, 99, 111, 108, 111, 114, 61, 35, 100, 97, 99, 48, 56, 57, 62, 82, 115, 46, 49, 57, 48, 48, 46, 48, 60, 47, 99, 62, 32, 105, 110, 32, 60, 99, 111, 108, 111, 114, 61, 35, 54, 68, 66, 57, 56, 48, 62, 65, 110, 100, 97, 114, 32, 66, 97, 104, 97, 114, 60, 47, 99, 111, 108, 111, 114, 62, 10, 114, 10, 112, 80, 108, 97, 121, 101, 114, 32, 42, 42, 42, 55, 50, 51, 56, 32, 119, 111, 110, 32, 60, 99, 111, 108, 111, 114, 61, 35, 100, 97, 99, 48, 56, 57, 62, 82, 115, 46, 54, 48, 48, 48, 46, 48, 60, 47, 99, 62, 32, 105, 110, 32, 60, 99, 111, 108, 111, 114, 61, 35, 53, 53, 69, 48, 70, 70, 62, 82, 97, 112, 105, 100, 32, 84, 101, 101, 110, 32, 80, 97, 116, 116, 105, 60, 47, 99, 111, 108, 111, 114, 62, 44, 32, 109, 111, 110, 101, 121, 32, 99, 111, 109, 101, 115, 32, 115, 111, 32, 102, 97, 115, 116, 33, 10, 72, 10, 70, 80, 108, 97, 121, 101, 114, 32, 60, 99, 111, 108, 111, 114, 61, 35, 48, 48, 102, 102, 48, 48, 62, 42, 42, 42, 55, 55, 54, 57, 60, 47, 99, 62, 32, 119, 105, 116, 104, 100, 114, 97, 119, 32, 60, 99, 111, 108, 111, 114, 61, 35, 100, 97, 99, 48, 56, 57, 62, 82, 115, 46, 53, 48, 48, 46, 48, 60, 47, 99, 62, 10, 114, 10, 112, 80, 108, 97, 121, 101, 114, 32, 42, 42, 42, 52, 56, 54, 55, 32, 119, 111, 110, 32, 60, 99, 111, 108, 111, 114, 61, 35, 100, 97, 99, 48, 56, 57, 62, 82, 115, 46, 51, 54, 48, 48, 46, 48, 60, 47, 99, 62, 32, 105, 110, 32, 60, 99, 111, 108, 111, 114, 61, 35, 53, 53, 69, 48, 70, 70, 62, 82, 97, 112, 105, 100, 32, 84, 101, 101, 110, 32, 80, 97, 116, 116, 105, 60, 47, 99, 111, 108, 111, 114, 62, 44, 32, 109, 111, 110, 101, 121, 32, 99, 111, 109, 101, 115, 32, 115, 111, 32, 102, 97, 115, 116, 33, 10, 97, 10, 95, 83, 111, 32, 108, 117, 99, 107, 121, 33, 32, 80, 108, 97, 121, 101, 114, 32, 42, 42, 42, 55, 55, 54, 57, 32, 119, 111, 110, 32, 60, 99, 111, 108, 111, 114, 61, 35, 100, 97, 99, 48, 56, 57, 62, 82, 115, 46, 50, 48, 48, 48, 46, 48, 60, 47, 99, 62, 32, 105, 110, 32, 60, 99, 111, 108, 111, 114, 61, 35, 54, 68, 66, 57, 56, 48, 62, 65, 110, 100, 97, 114, 32, 66, 97, 104, 97, 114, 60, 47, 99, 111, 108, 111, 114, 62, 10, 97, 10, 95, 83, 111, 32, 108, 117, 99, 107, 121, 33, 32, 80, 108, 97, 121, 101, 114, 32, 42, 42, 42, 56, 49, 56, 48, 32, 119, 111, 110, 32, 60, 99, 111, 108, 111, 114, 61, 35, 100, 97, 99, 48, 56, 57, 62, 82, 115, 46, 49, 57, 48, 48, 46, 48, 60, 47, 99, 62, 32, 105, 110, 32, 60, 99, 111, 108, 111, 114, 61, 35, 54, 68, 66, 57, 56, 48, 62, 65, 110, 100, 97, 114, 32, 66, 97, 104, 97, 114, 60, 47, 99, 111, 108, 111, 114, 62, 10, 97, 10, 95, 83, 111, 32, 108, 117, 99, 107, 121, 33, 32, 80, 108, 97, 121, 101, 114, 32, 42, 42, 42, 56, 49, 56, 48, 32, 119, 111, 110, 32, 60, 99, 111, 108, 111, 114, 61, 35, 100, 97, 99, 48, 56, 57, 62, 82, 115, 46, 49, 57, 48, 48, 46, 48, 60, 47, 99, 62, 32, 105, 110, 32, 60, 99, 111, 108, 111, 114, 61, 35, 54, 68, 66, 57, 56, 48, 62, 65, 110, 100, 97, 114, 32, 66, 97, 104, 97, 114, 60, 47, 99, 111, 108, 111, 114, 62}
	pData := packageData(CMD,
		int32(platform.ServerCommonCmd_CMD_BS_RESP),
		data, int64(len(data)))
	return pData
}

func SendNetworkPingPong() []byte {
	pData := packageData(GW,
		int32(platform.ServerGatewayCmd_CMD_GATEWAY_PING_RESP),
		[]byte{}, int64(0))
	return pData
}

func SendGatewayLogin() []byte {
	sCmd := int32(platform.ServerGatewayCmd_CMD_GATEWAY_LOGIN_RESP)
	return packHallData(GW, sCmd, loginResponse)
}

func SendUserAttri() []byte {
	userResp := &platform.MSG_GET_USER_ATTRI_RESP{
		UserAttris: []*platform.UserAttri{
			userAttri,
		},
	}
	data, _ := proto.Marshal(userResp)
	pData := packageData(CMD,
		int32(platform.ServerCommonCmd_CMD_GET_USER_ATTRI_RESP),
		data, int64(len(data)))
	return pData
}

func SendPlayerBalanceData() []byte {
	data, _ := proto.Marshal(balanceInfo)
	pData := packageData(CMD,
		int32(platform.ServerCommonCmd_CMD_GET_PLAYER_BALANCE_RESP),
		data, int64(len(data)))
	return pData
}

func packageData(main_cmd int32, sub_cmd int32, protoData []byte, dataSize int64) []byte {
	pkgsize := dataSize + 10
	datapkg := make([]byte, pkgsize)
	pByteLength := byteEncrypt(pkgsize, 2)
	datapkg[0] = pByteLength[0]
	datapkg[1] = pByteLength[1]
	pGateType := byteEncrypt(int64(main_cmd), 2)
	datapkg[2] = pGateType[0]
	datapkg[3] = pGateType[1]
	pProtoData := byteEncrypt(int64(sub_cmd), 2)
	datapkg[4] = pProtoData[0]
	datapkg[5] = pProtoData[1]
	datapkg[6] = 0
	datapkg[7] = 0
	datapkg[8] = 0
	datapkg[9] = 0
	for i, _ := range protoData {
		datapkg[i+10] = protoData[i]
	}
	return datapkg
}

func byteEncrypt(num int64, len int) []byte {
	x := num
	byteData := make([]byte, len)
	for i := 0; i < len; i++ {
		byteData[i] = byte(x & 255)
		x = x >> 8
	}
	return byteData
}

func byteDencrypt(byteData []byte, len int) int32 {
	x := int32(byteData[len-1])
	for i := 0; i < len-1; i++ {
		index := len - 1 - i
		x = x<<8 + int32(byteData[index-1])
	}
	return x
}

func unpkgData(data []byte) (int32, int32, int32, int32, []byte) {
	l := byteDencrypt([]byte{data[0], data[1]}, 2)
	gw := byteDencrypt([]byte{data[2], data[3]}, 2)
	pb := byteDencrypt([]byte{data[4], data[5]}, 2)
	uid := byteDencrypt([]byte{data[6], data[7], data[8], data[9]}, 4)
	pBytes := make([]byte, l-10)
	for i, _ := range pBytes {
		pBytes[i] = data[i+10]
	}
	return l, gw, pb, uid, pBytes
}

func callByTimeID(timerId string, tick time.Duration, callback func()) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(tick)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			callback()
			return
		case t := <-ticker.C:
			fmt.Println(fmt.Sprintf("TimerId:%s,Date: %v-%v-%v Time: %v:%v:%v",
				timerId, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
		}
	}
}
