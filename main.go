package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"gowebsocket/impl"
	"net/http"
)
var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	gConn  *impl.Connection
	gConnArray = make([]*impl.Connection,10)
	gConnectMax uint32 = 0
)
func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		conn   *impl.Connection
		data   []byte
	)
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	if conn, err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}
	gConn = conn
	for i := 0; i < 10; i++ {
		if nil == gConnArray[i]{
			gConnArray[i] = conn
			gConnectMax = uint32(i)+1
			break
		}
	}
	fmt.Println("gConnectMax:",gConnectMax)
	fmt.Println("gConnArray:",gConnArray)
	Init()
	go func() {

	}()

	for {
		data, err = conn.ReadMessage() //接收数据
		if len(data) == 0{
			client:= conn.OnWebsocketClose()
			fmt.Println(client)
			for i := 0; i < int(gConnectMax); i++ {
				if gConnArray[i] == conn {
					gConnArray[i] = nil
					gConnectMax = gConnectMax -1
					break
				}
			}
			fmt.Println("gConnectMax",gConnectMax)
			fmt.Println(gConnArray)
			return
		}
		l, mCmd, sCmd, uid, msg := unpkgData(data) //解包
		//fmt.Println(fmt.Sprintf("长度=%v,主协议=%v,消息协议=%v,UID=%v,protobuf结构=%v",l, mCmd, sCmd, uid, msg))
		if err != nil { //错误处理
			goto ERR
		}
		s1,s2:=conn.RemoteAddr()
		fmt.Println(fmt.Sprintf("client[%v-%v]:Connected",s1,s2))
		getDataByProtocol(conn, l, mCmd, sCmd, uid, msg) //处理协议数据
	}

ERR:
	conn.Close()

}

func main() {
	fmt.Println("WebSocket服务已启动: 0.0.0.0:6615")
	http.HandleFunc("/", wsHandler)
	http.ListenAndServe("0.0.0.0:6615", nil)
}
