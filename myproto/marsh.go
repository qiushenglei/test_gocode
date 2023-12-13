package myproto

import (
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

func StructCopy() {
	str1 := "MemberSingleChat_5849593942124239"
	str := `
{"mid":"5971273533215482042","seqId":19,"sender":5901959374578073,"receiver":5901959374578097,"forwarder":0,"replyId":"5971273533198704779","groupId":5849593942124239,"msgType":1,"msgFlag":1,"data":"G08YPNeV8eaBDOj9HVttxKhEgYjdZG1IxWuwV8EGJuaZ/RrbpBqemREWTF8EEjVS","timestamp":1702456727869,"version":"","uuid":"f9bcccdcda4747fa88b8413f5117b36c","clientType":"PressureTest","address":"127.0.0.1","deviceId":"2D4EB99F-2B6E-5A43-A76A-902211C9C1DD","delIds":[],"pinTopIds":[],"qid":"","globalId":322342,"backup1":"encryptVersion1","userIds":[],"backup2":"","backup3":"","backup4":0,"createTime":"2023-12-13T16:38:47+08:00","updateTime":"2023-12-13T16:38:47+08:00","isEdit":0,"isDelete":0,"isQueue":0}
`

	str2 := "MemberReadSingleMessage_5849593942124239_5901959374578097"
	str3 := 19

	MemberGroupChat_5849593942124241
	{"mid":"5966858835753707894","seqId":10,"sender":5901959374578073,"receiver":0,"forwarder":0,"replyId":"","groupId":5849593942124241,"msgType":2,"msgFlag":1,"data":"mypLGAqtBuzy2SDzujKeSg==","timestamp":1702300627518,"version":"","uuid":"1d9e4b2c30034c55bf45a7f3451c8544","clientType":"mac","address":"18.166.165.255","deviceId":"2D4EB99F-2B6E-5A43-A76A-902211C9C1DD","delIds":[],"pinTopIds":[],"qid":"","globalId":317542,"backup1":"encryptVersion1","userIds":[],"backup2":"","backup3":"","backup4":0,"createTime":"2023-12-11T21:17:07+08:00","updateTime":"2023-12-11T21:17:07+08:00","isEdit":0,"isDelete":0,"isQueue":0}
	MemberReadGroupMessage_5849593942124241_5901959374578097
	10

}
