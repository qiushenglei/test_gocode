package proto

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alibaba/sentinel-golang/logging"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
	"log"
	"net"
	"reflect"
)

type QProto struct {
	reader *bufio.Reader
	writer *bufio.Writer
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

func NewQProto(conn net.Conn) *QProto {
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	return &QProto{
		reader: r,
		writer: w,
	}
}

func (p *QProto) Verify(token string) error {
	n, err := p.writer.Write([]byte(token))
	if err != nil {
		return err
	}
	fmt.Println(n)
	//p.Send(0, token)
	response, err := p.Recv()
	if err != nil {
		return err
	}
	if response.Code != 200 {
		return errors.New("verify fail")
	}
	return nil
}

func (p *QProto) Auth() error {
	//token, err := p.reader.Peek(10)
	token := make([]byte, 10)
	_, err := p.reader.Read(token)
	fmt.Println(token, 1111)
	if err != nil {
		fmt.Println(token, 22222)
		return err
	}
	return p.authentication(token)
}

func (p *QProto) authentication(token []byte) error {
	if string(token) == "6666666666" {
		p.token = "6666666666"
	}
	return errorpkg.ErrNoLogin
}

func (p *QProto) Decode() (*Response, error) {
	header, err := p.reader.ReadSlice('\n')
	if err != nil {
		return nil, errorpkg.ErrParam
	}

	h := &Header{}
	if err := json.Unmarshal(header, h); err != nil {
		return nil, errorpkg.ErrParam
	}

	content := make([]byte, h.Length)
	if contentLen, err := p.reader.Read(content); err != nil {
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

func (p *QProto) Encode(h any, content []byte) error {
	header, err := json.Marshal(h)
	if err != nil {
		return err
	}
	body := fmt.Sprintf("%s\n%s", header, content)

	_, err = p.writer.Write([]byte(body))
	return err
}

func (p *QProto) Send(code int, response any) {
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
	if err := p.Encode(h, content); err != nil {
		logging.Error(err, "encode proto fail")
	}
}

func (p *QProto) Recv() (*Response, error) {
	return p.Decode()
}

func (p *QProto) Close() {
	p.writer.Flush()
}

func (p *QProto) GetToken() string {
	return p.token
}

func (p *QProto) SetToken(token string) {
	p.token = token
}
