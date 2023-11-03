package client

import (
	"fmt"
	"net"
	"testproj/socket/proto"
	"time"

	"github.com/alibaba/sentinel-golang/logging"
)

func Dial() {
	dsn := "127.0.0.1:9797"
	conn, err := net.Dial("tcp", dsn)
	if err != nil {
		panic(err)
	}
	defer func() {
		if connerr := conn.Close(); connerr != nil {
			fmt.Println("close connerr:", connerr)
		}
	}()

	parser := proto.NewQProto(conn)

	err = parser.Verify("6666666666")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("验证成功，进入双工通讯")

	for {
		parser.Send(0, "nihaoya")
		response, err := parser.Recv()
		if err != nil {
			logging.Error(err, "recv fail 1")
			continue
		}
		fmt.Println(string(response.Body))
		time.Sleep(3 * time.Second)
	}

}
