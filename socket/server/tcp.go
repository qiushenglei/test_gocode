package server

import (
	"log"
	"net"
	"testproj/socket/proto"
)

type Server struct {
	parser proto.QProto
}

func (s *Server) Start() {
	// StartServer()
}

func (s *Server) Stop() {
}

func StartServer() {
	dsn := "127.0.0.1:9797"
	lis, err := net.Listen("tcp", dsn)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {

	p := proto.NewQProto(conn)
	if err := p.Auth(); err != nil {
		p.Send(400, err)
		return
	}

	for {
		response, err := p.Recv()
		if err != nil {
			p.Send(500, err.Error())
			return
		}

		// 业务处理

		// 输出
		p.Send(200, "hello world, content: "+string(response.Body))
	}
}
