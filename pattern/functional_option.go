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
	// 彻底封装成员变量
	addr     string
	port     int
	protocol string
	timeout  time.Duration
	maxConns int
	cTLS     *tls.Config
}

type Option func(*Server)

func Protocol(p string) Option {
	return func(s *Server) {
		s.protocol = p
	}
}
func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}
func MaxConns(maxconns int) Option {
	return func(s *Server) {
		s.maxConns = maxconns
	}
}
func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.cTLS = tls
	}
}

func NewServer(addr string, port int, options ...Option) (*Server, error) {
	srv := Server{
		addr:     addr,
		port:     port,
		protocol: "tcp",
		timeout:  30 * time.Second,
		maxConns: 1000,
		cTLS:     nil,
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
	}
}

func Proc(srv *Server) {

}
