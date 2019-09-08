package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	// "bytes"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size allowed from peer.
	maxMessageSize = 8192
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  8192,
	WriteBufferSize: 8192,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 클라이언트 정보
type Client struct {
	Conn                   *websocket.Conn
	LocalDataChannelLabel  string
	RemoteDataChannelLabel string
	Send                   chan []byte
}
type Hub struct {
	Clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}
type WebSocketMessage struct {
	Service string      `json:"service,omitempty"`
	Label   string      `json:"label,omitempty"`
	Type    string      `json:"type,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (w *WebSocketMessage) Encode() (b []byte, err error) {
	b, err = json.Marshal(w)
	if err != nil {
		return
	}

	return
}

var WebSocketHub *Hub

func init() {
	WebSocketHub = newHub()
	go WebSocketHub.run()
}

func (c *Client) readPump() {
	defer func() {
		WebSocketHub.unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		var message WebSocketMessage
		err := c.Conn.ReadJSON(&message)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}
		fmt.Fprintf(os.Stderr, "RECEIVE: %v\n", message.Service)
		label := message.Label
		switch message.Service {
		case "Connect":
			fmt.Fprintf(os.Stderr, "Create Offer\n")
			//id := message.Data.(string)
			if WebRTCMap == nil {
				WebRTCMap = make(map[string]WebRTC)
			}
			webRTC := WebRTC{}
			webRTC.CreateDataChannel(label, message.Type)
			webRTC.CreateOffer()
			WebRTCMap[label] = webRTC
			if message.Type == "remote" {
				c.RemoteDataChannelLabel = label
			} else {
				c.LocalDataChannelLabel = label
			}

			message.Service = "Offer"
			message.Data = webRTC.Offer
			response, err := message.Encode()
			if err != nil {
				c.Send <- []byte(`{ "err" : "` + err.Error() + `" }`)
			} else {
				c.Send <- response
			}
		case "Answer":
			fmt.Fprintf(os.Stderr, "%#v\n", message)
			b, _ := json.Marshal(message.Data)
			if _, ok := WebRTCMap[label]; ok {
				c := WebRTCMap[label]
				c.ReceiveAnswer(string(b))
			}
		case "Candidate":
			fmt.Fprintf(os.Stderr, "%#v\n", message)
			b, _ := json.Marshal(message.Data)
			if _, ok := WebRTCMap[label]; ok {
				c := WebRTCMap[label]
				c.AddCandidate(string(b))
			}
		}
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			//fmt.Fprintf(os.Stderr, "Write:\n %#v\n", message)
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) CompareRemoteAddr(src, tgt string) bool {
	srcTokens := strings.Split(src, ":")
	tgtTokens := strings.Split(tgt, ":")
	srcAddr := ""
	for i, e := range srcTokens {
		if i == len(srcTokens)-1 {
			break
		}
		srcAddr += e
	}
	tgtAddr := ""
	for i, e := range tgtTokens {
		if i == len(tgtTokens)-1 {
			break
		}
		tgtAddr += e
	}
	if strings.Compare(srcAddr, tgtAddr) == 0 {
		return true
	} else {
		return false
	}
}
func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			fmt.Fprintf(os.Stderr, "DEBUG: WebSocket Opened(%v)\n", client.Conn.RemoteAddr())
			h.Clients[client] = true
			//go func() {
			//	for {
			//		if _, ok := h.Clients[client.ClientToken]; ok {
			//			rs := WebSocketMessage{
			//				Data: Payload{
			//					Category:  "ws",
			//					Service:   "ReadyToLive",
			//					Account:   h.Clients[client.ClientToken].User.Login.Account,
			//					MateToken: h.Clients[client.ClientToken].ClientToken,
			//				},
			//			}
			//			WebSocketHub.BroadcastToDermaster(h.Clients[client.ClientToken].User.Id, rs)
			//			time.Sleep(10 * time.Second)
			//		} else {
			//			break
			//		}
			//	}
			//}()
		case client := <-h.unregister:
			//logger.Info("Unregistered")
			if _, ok := h.Clients[client]; ok {
				//userid := client.user.UserId
				//remember := client.remember
				//session_uuid := client.session_uuid
				RemoveWebRTC(client.LocalDataChannelLabel)
				RemoveWebRTC(client.RemoteDataChannelLabel)
				delete(h.Clients, client)
				close(client.Send)

				//logger.Debugf("getCount: %v", h.getCount(userid))
				//logger.Infof("remeber: %v", remember)
				//logger.Debugf("session_uuid: %v", session_uuid)

				// if h.getCount(userid) == 0 && remember == false {
				// 	clearSession(session_uuid)
				// }
			}
		case message := <-h.broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					//logger.Info("Closed")
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}

func (h *Hub) Broadcast(message []byte) {
	h.broadcast <- message
}
func ServeWebSocket(w http.ResponseWriter, r *http.Request) {

	//_, err := session(w, r)
	//if err != nil {
	//	//logger.Error("websocket session check fail")
	//	return
	//}

	//cookie, _ := r.Cookie("userid")
	//var userid string
	//if cookie != nil {
	//	userid = cookie.Value
	//} else {
	//	userid = ""
	//}
	//
	//cookie, _ = r.Cookie("sessionUuid")
	//var session_uuid string
	//if cookie != nil {
	//	session_uuid = cookie.Value
	//} else {
	//	session_uuid = ""
	//}
	//
	//cookie, _ = r.Cookie("remember")
	//remember := false
	//if cookie != nil {
	//	if cookie.Value == "true" {
	//		remember = true
	//	} else {
	//		remember = false
	//	}
	//}

	// 인증 성공했으므로 반드시 있다.
	//User, err := data.UserByUserId(userid)
	//if err != nil {
	//	//logger.Errorf("Not found: %v", userid)
	//	w.WriteHeader(http.StatusNotAcceptable)
	//	return
	//}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	//var respBody Payload
	//if err := json.NewDecoder(r.Body).Decode(&respBody); err != nil {
	//	fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	//	return
	//}
	//respBody.Debug("< " + GetFunctionName() + " WebSocket Connection Incoming")

	client := &Client{
		Conn: conn,
		Send: make(chan []byte, 4096),
	}
	WebSocketHub.register <- client

	go client.writePump()
	go client.readPump()
}
