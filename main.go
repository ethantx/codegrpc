package main

import (
	"github.com/giteexx/codegrpc/internal/controller/codegrpc"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

func main() {
	grpcx.Resolver.Register(etcd.New("127.0.0.1:2379"))
	s := grpcx.Server.New(&grpcx.GrpcServerConfig{
		Address: ":850",//设置端口
		Name: "test",//设置通道名称(不要与现有的重名 必须用英文名。系统会根据这个名进行调用的)
		LogStdout:        true,//输出日志
		AccessLogEnabled: true,//开启访问日志  快捷配置方法不读取配置文件
		
	})
	codegrpc.Register(s)
	s.Run()
}
