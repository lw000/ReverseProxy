package ws

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"demo/ReverseProxy/cmd/ReverseProxy/network/ws/control"
	log "github.com/alecthomas/log4go"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  10240,
	WriteBufferSize: 10240,
	CheckOrigin:     func(r *http.Request) bool { return true },
} // use default options

func handleConnect(c *websocket.Conn) {
	defer func() {
		log.Info("%s exit", c.RemoteAddr().String())
		err := c.Close()
		if err != nil {
			log.Error("%v", err)
		}
	}()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Error(err)
			break
		}

		if mt == websocket.BinaryMessage {
			err = control.GetHub().DispatchMessage(c, message)
			if err != nil {
				break
			}
		} else {
			err = c.WriteMessage(mt, message)
			if err != nil {
				break
			}
		}
	}
}

func wsEcho(w http.ResponseWriter, r *http.Request) {
	c, er := upgrader.Upgrade(w, r, nil)
	if er != nil {
		log.Error("upgrade:", er)
		return
	}
	go handleConnect(c)
}

func RunEchoMain() {
	log.Info(fmt.Sprintf("listening [%d]", 8830))
	http.HandleFunc("/", wsEcho)
	log.Error(http.ListenAndServe(fmt.Sprintf(":%d", 8830), nil))
}
