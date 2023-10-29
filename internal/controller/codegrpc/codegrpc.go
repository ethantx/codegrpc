package codegrpc

import (
	"context"

	"github.com/giteexx/codegrpc/api"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/util/grand"
)

type Controller struct {
	api.UnimplementedGatewayServer
}

func Register(s *grpcx.GrpcServer) {
	api.RegisterGatewayServer(s.Server, &Controller{})
}

func (*Controller) GetBalance(ctx context.Context, req *api.WayRequest) (res *api.BalanceReply, err error) {
	return &api.BalanceReply{
		Data: "",
	}, nil
}

func (*Controller) UsePhone(ctx context.Context, req *api.WayRequest) (res *api.PhoneReply, err error) {
	return &api.PhoneReply{
		OutTradeNo: "",
		Phone:      req.GetApiField().Phone,
	}, nil
}

func (*Controller) GetPhone(ctx context.Context, req *api.WayRequest) (res *api.PhoneReply, err error) {
	return &api.PhoneReply{
		OutTradeNo: "",
		Phone:      grand.S(11),
	}, nil
}

func (*Controller) GetCode(ctx context.Context, req *api.WayRequest) (res *api.CodeReply, err error) {
	return &api.CodeReply{
		Code: "sss",
	}, nil
}

func (*Controller) CodeFeed(ctx context.Context, req *api.WayRequest) (res *api.EmptyReply, err error) {
	return
}

func (*Controller) PhoneFeed(ctx context.Context, req *api.WayRequest) (res *api.EmptyReply, err error) {
	return
}

func (*Controller) Rrelease(ctx context.Context, req *api.WayRequest) (res *api.EmptyReply, err error) {
	return
}
