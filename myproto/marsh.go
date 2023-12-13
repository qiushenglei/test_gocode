package myproto

import (
	"fmt"
	"google.golang.org/protobuf/proto"
)

// protoc --proto_path=./ --go_out=./  --go_opt=paths=source_relative  --go-grpc_out=./  --go-grpc_opt=paths=source_relative  chat.proto

func Marshal() {
	o := &Message{
		SeqId:  1233,
		Sender: 1231412,
	}
	b, err := proto.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
