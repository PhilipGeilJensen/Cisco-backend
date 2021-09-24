package api

import (
	"Backend/server"
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func (api *API) Websocket(r *server.APIRequest) (err error) {
	host, ok := r.Request.Vars["host"]
	if !ok {
		return
	}
	fmt.Println("Connection has been established")
	fmt.Println(host)
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(r.Request.W, r.Request.R, nil)
	if err != nil {
		log.Printf("There was an error upgrading to WS: %s", err)
	}
	socketConn := server.SocketClient{
		Conn: ws,
		Host: host,
	}
	defer func() {
		api.srv.Websockets = server.RemoveConn(api.srv.Websockets, socketConn)
		ws.Close()
	}()
	api.srv.Websockets = append(api.srv.Websockets, socketConn)
	//TODO Upgrade to websocket here!

	for {
		_, message, err := socketConn.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		fmt.Println(message)
	}
	return nil
}
