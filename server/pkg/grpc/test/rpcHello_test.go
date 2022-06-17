package test

import (
	"context"
	grpc_etcdv3 "github.com/tmnhs/fginx/server/pkg/grpc/getcdv3"
	pb "github.com/tmnhs/fginx/server/pkg/grpc/test/proto"
	"strings"
	"testing"
	"time"
)

func TestRpcHello_Hello(t *testing.T) {
	rpcServer := NewRpcChatServer(9000)
	go rpcServer.Run()
	time.Sleep(time.Second * 5)

	etcdConn := grpc_etcdv3.GetConn(rpcServer.etcdSchema, strings.Join(rpcServer.etcdAddr, ","), rpcServer.rpcRegisterName)
	client := pb.NewUserInfoServiceClient(etcdConn)
	func() {
		//time.Sleep(time.Second*5)
		_, _ = client.GetUserInfo(context.Background(), &pb.UserRequest{
			Name: "tmnhs",
		})
		_, _ = client.Hello(context.Background(), &pb.UserRequest{
			Name: "tmnhs",
		})
	}()

}
