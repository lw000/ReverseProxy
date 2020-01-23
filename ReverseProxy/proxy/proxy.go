package proxy

import (
	"fmt"
	log "github.com/alecthomas/log4go"
	"net"
)

type TCPProxy struct {
}

func (proxy *TCPProxy) Start(listenPort, targetPort int) {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", listenPort))
	if err != nil {
		log.Error("Unable to listen on: %d, error: %s", listenPort, err.Error())
		return
	}

	defer func() {
		_ = listen.Close()
	}()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Error("unable to accept a request, error: %s", err.Error())
			continue
		}
		go proxy.runForward(conn, targetPort)
	}
}

func (proxy *TCPProxy) runForward(conn net.Conn, targetPort int) {
	defer func() {
		if x := recover(); x != nil {
			log.Error(x)
		}
	}()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Error("unable to read from input, error: %s", err)
		return
	}

	var targetConn net.Conn
	targetConn, err = net.Dial("tcp", fmt.Sprintf(":%d", targetPort))
	if err != nil {
		log.Error("unable to connect to: %d, error: %s", targetPort, err)
		_ = conn.Close()
		return
	}

	n, err = targetConn.Write(buf[:n])
	if err != nil {
		log.Error("unable to write to output, error: %s", err)
		_ = conn.Close()
		_ = targetConn.Close()
		return
	}

	go proxy.forwardRequest(conn, targetConn)
	go proxy.forwardRequest(targetConn, conn)
}

// Forward all requests from r to w
func (proxy *TCPProxy) forwardRequest(r net.Conn, w net.Conn) {
	defer func() {
		if x := recover(); x != nil {
			log.Error(x)
		}
		_ = r.Close()
		_ = w.Close()
	}()

	var buf = make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil {
			log.Error("unable to read from input, error: %s", err)
			break
		}

		n, err = w.Write(buf[:n])
		if err != nil {
			log.Error("unable to write to output, error: %s", err)
			break
		}
	}
}
