package main

import (
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type websocketClientManager struct {
	conn        *websocket.Conn
	addr        *string
	path        string
	sendMsgChan chan []byte //string
	recvMsgChan chan []byte //string
	isAlive     bool
	timeout     int
}

// 构造函数
func NewWsClientManager(addrIp, addrPort, path string, timeout int) *websocketClientManager {
	addrString := addrIp + ":" + addrPort
	var sendChan = make(chan []byte, 65535)
	var recvChan = make(chan []byte, 65535)
	var conn *websocket.Conn
	return &websocketClientManager{
		addr:        &addrString,
		path:        path,
		conn:        conn,
		sendMsgChan: sendChan,
		recvMsgChan: recvChan,
		isAlive:     false,
		timeout:     timeout,
	}
}

// 链接服务端
func (wsc *websocketClientManager) dail() {
	var err error
	u := url.URL{Scheme: "ws", Host: *wsc.addr, Path: wsc.path}
	log.Printf("connecting to %s", u.String())
	wsc.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return

	}
	wsc.isAlive = true
	log.Printf("connecting to %s 链接成功！！！", u.String())

}

// 发送消息
func (wsc *websocketClientManager) sendMsgThread() {
	go func() {
		for {
			msg := <-wsc.sendMsgChan
			fmt.Println("发送给服务器:", msg)
			err := wsc.conn.WriteMessage(websocket.BinaryMessage, []byte(msg))
			if err != nil {
				log.Println("write:", err)
				continue
			}
		}
	}()
}

// 读取消息
func (wsc *websocketClientManager) readMsgThread() {
	go func() {
		for {
			if wsc.conn != nil {
				_, message, err := wsc.conn.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					wsc.isAlive = false
					// 出现错误，退出读取，尝试重连
					break
				}
				log.Printf("接收到服务器发来的:", message)
				log.Printf("============================================================")
				// 需要读取数据，不然会阻塞
				wsc.recvMsgChan <- []byte(message)
			}

		}
	}()
}

// 开启服务并重连
func (wsc *websocketClientManager) start() {
	for {
		if wsc.isAlive == false {
			wsc.dail()
			wsc.sendMsgThread()
			wsc.readMsgThread()
		}
		time.Sleep(time.Second * time.Duration(wsc.timeout))
	}
}

//15,0,76,4,9,0,0,0,0,0,10,3,-80,-2,34
func makeAction(mCmd int, sCmd int, len int) []byte {
	data := make([]byte, 10)
	for i := 0; i < 10; i++ {
		data[0] = byte(len & 255)
		data[1] = byte(len >> 8)
		data[2] = byte(mCmd & 255)
		data[3] = byte(mCmd >> 8)
		data[4] = byte(sCmd & 255)
		data[5] = byte(sCmd >> 8)
		data[6] = 0
		data[7] = 0
		data[8] = 0
		data[9] = 0
	}

	return data
}

func main() {
	//wsc := NewWsClientManager("192.168.1.164", "6615", "", 10)
	wsc := NewWsClientManager("127.0.0.1", "6615", "", 10)
	//wsc.sendMsgChan <- makeAction(1000, 1, 10) //login
	// wsc.sendMsgChan <- makeAction(1000, 3, 10)  //login
	// wsc.sendMsgChan <- makeAction(1000, 5, 10)  //balance
	// wsc.sendMsgChan <- makeAction(1000, 7, 10)  //repeat login
	// wsc.sendMsgChan <- makeAction(1000, 9, 10)  //send ping pong
	// wsc.sendMsgChan <- makeAction(1100, 9, 15) //money
	//wsc.sendMsgChan <- makeAction(1100, 3, 10) //userinfo
	// wsc.sendMsgChan <- makeAction(1100, 11, 10) //userinfo
	// wsc.sendMsgChan <- makeAction(1100, 13, 10) //userinfo
	// wsc.sendMsgChan <- makeAction(1100, 15, 10) //horse race lamp
	//wsc.sendMsgChan <- makeAction(1101, 1, 10) //room data
	//wsc.sendMsgChan <- makeAction(1101, 5, 10)  //table data

	//wsc.sendMsgChan <- makeAction(4002, 1, 10) //table status
	// wsc.sendMsgChan <- makeAction(4002, 3, 10)  //ready
	// wsc.sendMsgChan <- makeAction(4002, 7, 10)  //time end
	// wsc.sendMsgChan <- makeAction(4002, 13, 10) //bet
	// wsc.sendMsgChan <- makeAction(4002, 15, 10) //see
	// wsc.sendMsgChan <- makeAction(4002, 17, 10) //show
	// wsc.sendMsgChan <- makeAction(4002, 19, 10) //sideshow
	// wsc.sendMsgChan <- makeAction(4002, 21, 10) //sideshow
	// wsc.sendMsgChan <- makeAction(4002, 23, 10) //
	// wsc.sendMsgChan <- makeAction(4002, 25, 10) //continue game

	// wsc.sendMsgChan <- makeAction(4002, 60, 10)  //live
	// wsc.sendMsgChan <- makeAction(4002, 62, 10)  //chat
	// wsc.sendMsgChan <- makeAction(4002, 100, 10) //freetable

	v:=0
	for  {
		fmt.Println(v,v%2)
		time.Sleep(2*time.Second)
		v=v+1
	}

	//d := []byte{8, 251, 238, 35, 18, 2, 8, 3}
	// p := &teenpatti.MSG_C_GAME_READY_3_RESP{}
	// proto.Unmarshal(d, p)
	// fmt.Println(p)
	// fmt.Println(1000&255, 1000>>8)
	// fmt.Println(1100&255, 1100>>8)
	wsc.start()
	var w1 sync.WaitGroup
	w1.Add(1)
	w1.Wait()
}
