package myproto

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	"time"
)

// protoc --proto_path=./ --go_out=./  --go_opt=paths=source_relative  --go-grpc_out=./  --go-grpc_opt=paths=source_relative  chat.proto

func Marshal() {

	// msgFlat : 8
	//str := "{\"mid\":\"5971273533181928919\",\"seqid\":0,\"sender\":5901959374578073,\"receiver\":0,\"forwarder\":0,\"replyid\":\"\",\"groupid\":5849593942124228,\"msgtype\":2,\"msgflag\":8,\"data\":\"\",\"timestamp\":1702450724814,\"version\":\"\",\"uuid\":\"f9bcccdcda4747fa88b8413f5117b36c\",\"nickname\":\"yyds1123 usopp\",\"avatar\":\"\",\"qid\":\"\",\"globalid\":0,\"clienttype\":\"mac\",\"address\":\"18.166.165.255\",\"deviceid\":\"2d4eb99f-2b6e-5a43-a76a-902211c9c1dd\"}"

	// msgFlag ： 13 已读信息
	//str := "{\"mid\":\"5971273533181929143\",\"seqId\":2,\"sender\":5901959374578073,\"receiver\":5901959374578073,\"forwarder\":0,\"replyId\":\"\",\"groupId\":5849593942124121,\"msgType\":2,\"msgFlag\":13,\"data\":\"\",\"timestamp\":1702451278400,\"version\":\"\",\"uuid\":\"4f8dbe8f69e34853a8906e5bffc9c6fb\",\"nickName\":\"yyds1123 usopp\",\"avatar\":\"\",\"qid\":\"5939333722414261202\",\"globalId\":0,\"clientType\":\"mac\",\"address\":\"18.166.165.255\",\"deviceId\":\"2D4EB99F-2B6E-5A43-A76A-902211C9C1DD\"}"

	// msgFlag : 1  usopp 发给 usopp1 文字消息
	str := "{\"mid\":\"\",\"seqid\":0,\"sender\":5901959374578073,\"receiver\":0,\"forwarder\":0,\"replyid\":\"\",\"groupid\":5849593942124239,\"msgtype\":1,\"msgflag\":1," +
		"\"data\":\"\",\"timestamp\":1702450724814,\"version\":\"\",\"uuid\":\"f9bcccdcda4747fa88b8413f5117b36c\",\"nickname\":\"yyds1123 usopp\",\"avatar\":\"\",\"qid\":\"\",\"globalid\":0," +
		"\"clienttype\":\"mac\",\"address\":\"18.166.165.255\",\"deviceid\":\"2d4eb99f-2b6e-5a43-a76a-902211c9c1dd\"}"

	readmsg := &Message{
		Sender:    5901959374578073,
		Receiver:  5901959374578097,
		Forwarder: 0,
		GroupId:   5849593942124239,
		MsgType:   1,
		MsgFlag:   1,
		Data:      "{\"text\":\"hello usopp1,  my  name is usopp\"}",
		Timestamp: time.Now().Unix(),
		Uuid:      "f9bcccdcda4747fa88b8413f5117b36c",
		Nickname:  "yyds1123 usopp",
	}
	replymsg := &Message{
		Sender:    5901959374578073,
		Receiver:  5901959374578097,
		Forwarder: 0,
		ReplyId:   "5971273533198704779",
		GroupId:   5849593942124239,
		MsgType:   1,
		MsgFlag:   1,
		Data:      "{\"text\":\"hello usopp1,  my  name is usopp\"}",
		Timestamp: time.Now().Unix(),
		Uuid:      "f9bcccdcda4747fa88b8413f5117b36c",
		Nickname:  "yyds1123 usopp",
	}
	o := new(Message)
	err := json.Unmarshal([]byte(str), o)

	b, err := proto.Marshal(readmsg)
	b, err = proto.Marshal(replymsg)
	if err != nil {
		panic(err)
	}
	hexStr := fmt.Sprintf("%x", b)
	fmt.Println(string(hexStr))
}

func Base646Decode() {
	str := "2d2397dee6a4093bde5f60593166df44f503892ad94cb929e32decc1e3e10e7ced79c77aaf35c673ffabfdb197b58a6d4f3db8a595948460f0618085939751e70e288b7a50167fb8af481294561483c17b9ab793ec75da436075074ef2fd539ece64a4e183a0280f4829771acfd9di"
	str = "2d2397dee6a4093bde5f60593166df44f503892ad94cb929e32decc1e3e10e7ced79c77aaf35c673ffabfdb197b58a6d4f3db8a595948460f0618085939751e70e288b7a50167fb8af481294561483c17b9ab793ec75da436075074ef2fd539ece64a4e183a0280f4829771acfd9d"
	str = "d4e3dc41cbb4c6171f07dd29f579010782a0c1d3cd544764b50d3238795fadfdcb26039042082aa1c2fe33ebc51f953cfa1b932c39bdd03b96f67923e58de87192acd137a86077ff869610df38b25c392976962313c6fb0ae93b38100330223db0eca08aa61d42a588ffad3eba157a77dea80d6748ea3be487efcefff0f0ce6e18c68680741d9e0ddc1ffa0658a111533c0fe78d69dc8e29dc3508b4ea90081a"

	var b []byte
	//_, err := hex.Decode([]byte(str), b)
	b, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	readmsg := new(Message)
	err = proto.Unmarshal(b, readmsg)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

// d4e3dc41cbb4c6171f07dd29f579010782a0c1d3cd544764b50d3238795fadfdcb26039042082aa1c2fe33ebc51f953cfa1b932c39bdd03b96f67923e58de87192acd137a86077ff869610df38b25c392976962313c6fb0ae93b38100330223db0eca08aa61d42a588ffad3eba157a77dea80d6748ea3be487efcefff0f0ce6e18c68680741d9e0ddc1ffa0658a111533c0fe78d69dc8e29dc3508b4ea90081a
