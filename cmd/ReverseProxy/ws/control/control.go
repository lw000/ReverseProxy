package control

import (
	"demo/ReverseProxy/constant"
	"demo/ReverseProxy/protocol"
	"tuyue/tuyue_common/ws/packet"
	"tuyue/tuyue_common/ws/srv/hub"

	log "github.com/alecthomas/log4go"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

var (
	hub *tyhub.Hub
)

func init() {
	registerHub()
}

func AckMessage(c *websocket.Conn, data []byte) {
	if err := c.WriteMessage(websocket.BinaryMessage, data); err != nil {
		log.Error(err)
	}
}

func onRegisgerService(c *websocket.Conn, pk *typacket.Packet) {
	var req proxy.ReqRegService
	if err := proto.Unmarshal(pk.Data(), &req); err != nil {
		log.Error(err)
		return
	}
	log.Info("[%s] %+v req:%+v", c.RemoteAddr().String(), pk, req)

	ack := typacket.NewPacket(pk.Mid(), pk.Sid(), pk.RequestId())
	if err := ack.EncodeProto(&proxy.AckRegService{Code: 0, Msg: ""}); err == nil {
		AckMessage(c, ack.Data())
	}
}

func onHeartBeat(c *websocket.Conn, pk *typacket.Packet) {
	log.Info("[%s] %+v", c.RemoteAddr().String(), pk)

	ack := typacket.NewPacket(pk.Mid(), pk.Sid(), pk.RequestId())
	if err := ack.EncodeProto(nil); err == nil {
		AckMessage(c, ack.Data())
	}
}

func onUpdateConfig(c *websocket.Conn, pk *typacket.Packet) {
	var req proxy.ReqUpdateConfig
	if err := proto.Unmarshal(pk.Data(), &req); err != nil {
		log.Error(err)
		return
	}
	log.Info("[%s] %+v req:%+v", c.RemoteAddr().String(), pk, req)

	ack := typacket.NewPacket(pk.Mid(), pk.Sid(), pk.RequestId())
	if err := ack.EncodeProto(&proxy.AckUpdateConfig{Code: 1, Msg: ""}); err == nil {
		AckMessage(c, ack.Data())
	}
}

func registerHub() {
	hub = tyhub.NewWSHub()
	hub.Handle(constant.MDM_SERVICE, constant.SUB_SERVICE_REGISTER, onRegisgerService)
	hub.Handle(constant.MDM_HEARTBEAT, constant.SUB_HEARTBEAT, onHeartBeat)
	hub.Handle(constant.MDM_CMD, constant.SUB_CMD_UPDATECONFIG, onUpdateConfig)
}

func GetHub() *tyhub.Hub {
	return hub
}
