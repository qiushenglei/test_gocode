package server

import (
	"context"
	"net"
	"reflect"
	"testing"
	"testproj/socket/proto"
)

func TestListenSignal(t *testing.T) {
	type args struct {
		listener net.Listener
		cancel   context.CancelFunc
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListenSignal(tt.args.listener, tt.args.cancel)
		})
	}
}

func TestRecv(t *testing.T) {
	type args struct {
		c context.Context
		p *proto.QProto
	}
	tests := []struct {
		name string
		args args
		want func() error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Recv(tt.args.c, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Recv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSend(t *testing.T) {
	type args struct {
		c context.Context
		p *proto.QProto
	}
	tests := []struct {
		name string
		args args
		want func() error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Send(tt.args.c, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Send() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Start(t *testing.T) {
	type fields struct {
		parser proto.QProto
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				parser: tt.fields.parser,
			}
			s.Start()
		})
	}
}

func TestServer_Stop(t *testing.T) {
	type fields struct {
		parser proto.QProto
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				parser: tt.fields.parser,
			}
			s.Stop()
		})
	}
}

func TestStartServer(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StartServer()
		})
	}
}

func Test_process(t *testing.T) {
	type args struct {
		ctx  context.Context
		conn net.Conn
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			process(tt.args.ctx, tt.args.conn)
		})
	}
}
