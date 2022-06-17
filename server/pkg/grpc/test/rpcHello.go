package test

import (
	"context"
	"fmt"
	grpc_etcdv3 "github.com/tmnhs/fginx/server/pkg/grpc/getcdv3"
	pb "github.com/tmnhs/fginx/server/pkg/grpc/test/proto"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"strings"
)

type rpcHello struct {
	rpcPort         int
	rpcRegisterName string
	etcdSchema      string
	etcdAddr        []string
}

func NewRpcChatServer(port int) *rpcHello {
	rc := rpcHello{
		rpcPort:         port,
		rpcRegisterName: "hello",
		etcdSchema:      "hello",
		etcdAddr:        []string{"127.0.0.1:2379"},
	}
	return &rc
}
func (rpc *rpcHello) Run() {
	fmt.Println("", "", "rpc get_token init...")

	address := "127.0.0.1:" + strconv.Itoa(rpc.rpcPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("", "", "listen network failed, err = %s, address = %s", err.Error(), address)
		return
	}
	fmt.Println("", "", "listen network success, address = ", address)

	//grpc server
	srv := grpc.NewServer()
	defer srv.GracefulStop()

	//service registers with etcd

	pb.RegisterUserInfoServiceServer(srv, rpc)
	err = grpc_etcdv3.RegisterEtcd(rpc.etcdSchema, strings.Join(rpc.etcdAddr, ","), "127.0.0.1", rpc.rpcPort, rpc.rpcRegisterName, 10)
	if err != nil {
		fmt.Println("", "", "register rpc get_token to etcd failed, err = %s", err.Error())
		return
	}

	err = srv.Serve(listener)
	if err != nil {
		fmt.Println("", "", "rpc get_token fail, err = %s", err.Error())
		return
	}
	fmt.Println("rpc get_token init success")
}
func (rpc *rpcHello) GetUserInfo(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	fmt.Println("GetUserInfo,", in.Name)
	return nil, nil
}

func (rpc *rpcHello) Hello(ctx context.Context, in *pb.UserRequest) (*pb.HelloResponse, error) {
	fmt.Println("Hello...", in.Name)
	return nil, nil
}
