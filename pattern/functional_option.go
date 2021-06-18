package pattern

import (
	"crypto/tls"
	"time"
)

// https://coolshell.cn/articles/21146.html#Functional_Options
// 强烈推荐使用Functional Options这种方式，这种方式至少带来了如下的好处：
//     直觉式的编程
//     高度的可配置化
//     很容易维护和扩展
//     自文档
//     对于新来的人很容易上手
//     没有什么令人困惑的事（是nil 还是空）

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

type Option func(*Server)

func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}
func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}
func MaxConns(maxconns int) Option {
	return func(s *Server) {
		s.MaxConns = maxconns
	}
}
func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServer(addr string, port int, options ...Option) (*Server, error) {

	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 1000,
		TLS:      nil,
	}
	for _, option := range options {
		option(&srv)
	}
	//...
	return &srv, nil
}

func show() {
	if srv, err := NewServer(
		"1.1.1.1", 3330,
		Protocol("tcp4"),
		Timeout(time.Duration(10)*time.Second),
		TLS(&tls.Config{}),
	); err == nil {
		Proc(srv)
	} else {
		return
	}
}

func Proc(srv *Server) {

}
