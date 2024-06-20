package network

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	. "go_chat_project/types"
	"log"
	"net/http"
	"time"
)

var upgrader = &websocket.Upgrader{ReadBufferSize: SocketBufferSize, WriteBufferSize: MessageBufferSize, CheckOrigin: func(r *http.Request) bool { return true }}

type message struct {
	Name    string
	Message string
	Time    int64
}
type Room struct {
	Forward chan *message // 수신 메세지 보관하는 값
	// 들어오는 메세지를 다른 클라이언트에게 전송

	Join   chan *client     // Socket 연결되는 경우에 작동
	Leave  chan *client     // Socket 끊어지는 경우에 작동
	Client map[*client]bool // 현재 방에 있는 Client 정보를 저장
}

type client struct {
	Send   chan *message
	Room   *Room
	Name   string
	Socket *websocket.Conn
}

func NewRoom() *Room {
	return &Room{
		Forward: make(chan *message),
		Join:    make(chan *client),
		Leave:   make(chan *client),
		Client:  make(map[*client]bool),
	}
}

func (c *client) Read() {
	// 클라이언트가 들어오는 메세지를 읽는 함수
	defer c.Socket.Close()
	for {
		var msg *message
		err := c.Socket.ReadJSON(&msg)
		if err != nil {
			if !websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				break
			} else {
				panic(err)
			}
		} else {
			log.Println("READ : ", msg, "client", c.Name)
			msg.Time = time.Now().Unix()
			msg.Name = c.Name

			c.Room.Forward <- msg
		}
	}
}

func (c *client) Write() {
	defer c.Socket.Close()
	// 클라이언트가 메세지를 전송하는 함수
	for msg := range c.Send {
		log.Println("WRITE : ", msg, "client", c.Name)
		err := c.Socket.WriteJSON(msg)
		if err != nil {
			panic(err)
		}
	}
}

func (r *Room) RunInit() {
	// Room 에 있는 모든 채널값들을 받는 역할
	for {
		select {
		case client := <-r.Join:
			r.Client[client] = true
		case client := <-r.Leave:
			r.Client[client] = false
			close(client.Send)
			delete(r.Client, client)
		case msg := <-r.Forward:
			for client := range r.Client {
				client.Send <- msg
			}
		}
	}
}

func (r *Room) SocketServe(c *gin.Context) {
	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	userCookie, err := c.Request.Cookie("auth")
	if err != nil {
		panic(err)
	}

	client := &client{
		Socket: socket,
		Send:   make(chan *message, MessageBufferSize),
		Room:   r,
		Name:   userCookie.Value,
	}

	r.Join <- client

	defer func() { r.Leave <- client }()

	go client.Write()

	client.Read()
}
