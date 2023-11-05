package proto

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/alibaba/sentinel-golang/logging"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
	"log"
	"net"
	"reflect"
	"time"
)

type QProto struct {
	ctx    context.Context
	cancel context.CancelFunc
	reader *bufio.Reader
	writer *bufio.Writer
	conn   net.Conn
	token  string
}

type Header struct {
	Length int    `json:"length"`
	token  string `json:"token"`
	Code   int    `json:"code,omitempty"`
}

type Response struct {
	*Header
	Body []byte
}

func NewQProto(ctx context.Context, conn net.Conn, cancel context.CancelFunc) *QProto {
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	return &QProto{
		ctx:    ctx,
		conn:   conn,
		cancel: cancel,
		reader: r,
		writer: w,
	}
}

func (p *QProto) Verify(token string) error {
	// 发送token验证
	if err := p.Send(0, token); err != nil {
		return err
	}

	// 接收验证结果
	response, err := p.Recv()
	if err != nil {
		return err
	}

	if response.Code != 200 {
		return errorpkg.ErrAuthFailed
	}

	fmt.Println("客户端鉴权成功", string(response.Body))

	go p.HeartBeat()

	return nil
}

// Auth server端鉴权
func (p *QProto) Auth() error {
	// 读取n个字节，但是reader不移动
	//token, err := p.reader.Peek(10)

	// 读取len(token)个字节
	//token := make([]byte, 10)
	//_, err := p.reader.Read(token)

	resp, err := p.Recv()
	if err != nil {
		return errorpkg.ErrAuthFailed
	}

	return p.authentication(resp.Body)
}

// authentication 鉴权
func (p *QProto) authentication(token []byte) error {
	if string(token) == "6666666666" {
		p.token = "6666666666"
		return nil
	}
	return errorpkg.ErrAuthFailed
}

// HeartBeat 客户端发起心跳
func (p *QProto) HeartBeat() {
	// 监听关闭信号
	var isDone bool
	go func() {
		select {
		case <-p.ctx.Done():
			isDone = true
			return
		}
	}()

	ch := time.Tick(2 * time.Second)
	for {
		if isDone {
			fmt.Println("心跳结束")
			return
		}
		select {
		case <-ch:
			err := p.heartBeat()
			if err != nil {
				p.Close("heart beat")
			}
		}
	}
}

// heartBeat 发送心跳数据包
func (p *QProto) heartBeat() error {
	if err := p.Send(0, "心跳"); err != nil {
		return err
	}
	resp, err := p.Recv()
	if err != nil {
		return err
	}
	fmt.Println("接收到心跳回复", string(resp.Body))
	return nil
}

// decode 解码协议
func (p *QProto) decode() (*Response, error) {
	// 读取头部
	header, err := p.reader.ReadSlice('\n')
	if err != nil {
		fmt.Println("读取头部错误", err.Error())
		return nil, errorpkg.ErrSystem
	}

	h := &Header{}
	if err := json.Unmarshal(header, h); err != nil {
		fmt.Println("解析头部错误", err.Error())
		return nil, errorpkg.ErrParam
	}

	// 根据len获取content
	content := make([]byte, h.Length)
	if contentLen, err := p.reader.Read(content); err != nil {
		fmt.Println("读取内容错误", err.Error())
		return nil, errorpkg.ErrParam
	} else {
		log.Println(contentLen)
	}

	response := &Response{
		h,
		content,
	}
	return response, err
}

// encode 编码协议
func (p *QProto) encode(h any, content []byte) error {
	header, err := json.Marshal(h)
	if err != nil {
		return errorpkg.ErrSystem
	}
	body := fmt.Sprintf("%s\n%s", header, content)

	if _, err = p.writer.Write([]byte(body)); err != nil {
		return errorpkg.ErrSystem
	}

	if err = p.writer.Flush(); err != nil {
		return errorpkg.ErrSystem
	}
	return err
}

// Send 发送数据
func (p *QProto) Send(code int, response any) error {
	v := reflect.ValueOf(response)
	var content []byte
	switch v.Kind() {
	case reflect.Struct:
		content, _ = json.Marshal(v)
	case reflect.String:
		content = []byte(response.(string))
	}
	h := Header{
		Length: len(content),
		Code:   code,
	}
	if err := p.encode(h, content); err != nil {
		logging.Error(err, "encode proto fail")
		return err
	}
	return nil
}

// Recv 发送
func (p *QProto) Recv() (*Response, error) {
	return p.decode()
}

func (p *QProto) Close(action string) {
	// 触发双工通道 和 心跳协程关闭
	p.cancel()

	// 等7秒后关闭链接
	ctx, _ := context.WithTimeout(context.Background(), time.Second*7)
	select {
	case <-ctx.Done():
		fmt.Println(action + "close connection")
		p.conn.Close()
	}
}
