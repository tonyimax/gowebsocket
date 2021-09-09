package impl
import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)
type Connection struct {
	wsConnect *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte
	mutex     sync.Mutex // 对closeChan关闭上锁
	isClosed  bool       // 防止closeChan被关闭多次
}
func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConnect: wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
	}
	go conn.readLoop()
	go conn.writeLoop()
	return
}

func (conn *Connection) getOutChan() chan []byte {
	return conn.outChan
}

func (conn *Connection) getInChan() chan []byte {
	return conn.inChan
}

func GetGoroutineID() uint64 {
	b := make([]byte, 64)
	runtime.Stack(b, false)
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func (conn *Connection) ReadMessage() (data []byte, err error) {

	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("connection is closeed")
	}
	return
}

func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closeed")
	}
	return
}

func (conn *Connection) RemoteAddr() (string,string) {
	ra := conn.wsConnect.RemoteAddr()
	return ra.Network(),ra.String()
}

func (conn *Connection) Close() {
	// 线程安全，可多次调用
	conn.wsConnect.Close()
	// 利用标记，让closeChan只关闭一次
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

func (conn *Connection) OnWebsocketClose() string  {
	ra := conn.wsConnect.RemoteAddr()
	result := fmt.Sprintf("%v-%v",ra.Network(),ra.String())
	fmt.Println(fmt.Sprintf("onWebsocketClose:->client[%v]->closed!",result))
	return result
}

func (conn *Connection) readLoop() {
	fmt.Println("协程ID:", GetGoroutineID())
	var (
		data []byte
		err  error
	)
	for {
		_, data, err = conn.wsConnect.ReadMessage()
		if err != nil {
			//conn.OnWebsocketClose()
			goto ERR
		}
		//阻塞在这里，等待inChan有空闲位置
		select {
		case conn.inChan <- data:
		case <-conn.closeChan: // closeChan 感知 conn断开
			goto ERR
		}

	}

ERR:
	conn.Close()
}

func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)

	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR
		}
		//发二进制数据
		if err = conn.wsConnect.WriteMessage(websocket.BinaryMessage, data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()

}
