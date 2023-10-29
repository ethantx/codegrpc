package main

import (
	"github.com/giteexx/codegrpc/internal/controller/codegrpc"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

func main() {
	grpcx.Resolver.Register(etcd.New("127.0.0.1:2379"))
	s := grpcx.Server.New()
	codegrpc.Register(s)
	s.Run()
}
