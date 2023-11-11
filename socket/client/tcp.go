package client

import (
	"context"
	"fmt"
	"net"
	"testproj/socket/proto"
	"time"
)

func Dial() {
	dsn := "127.0.0.1:9797"
	conn, err := net.Dial("tcp", dsn)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	parser := proto.NewQProto(ctx, conn, cancel)

	err = parser.Verify("6666666666")
	if err != nil {
		fmt.Println(err.Error())
	}

	for {
		parser.Send(0, "nihaoya")
		response, err := parser.Recv()
		if err != nil {
			//logging.Error(err, "recv fail 1")
			parser.Close("client")
			return
		}
		fmt.Println(string(response.Body))
		time.Sleep(3 * time.Second)
	}

}
