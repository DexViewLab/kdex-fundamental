package server

import (
	"fmt"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

const (
	NETWORK  string = "tcp"
	recvSize int    = 12 * 1024 * 1024
	sendSize int    = 12 * 1024 * 1024
)

func StartGRPCServer(port string, cert string, keyFile string, f func(grpcServer *grpc.Server)) error {
	address := fmt.Sprintf("0.0.0.0%s", port)
	listener, err := net.Listen(NETWORK, address)
	if err != nil {
		glog.Fatalf("net.Listen err: %v", err)
	}

	var options = []grpc.ServerOption{
		grpc.MaxRecvMsgSize(recvSize),
		grpc.MaxSendMsgSize(sendSize),
		grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionIdle: 2 * time.Minute}),
	}

	var s *grpc.Server
	if cert != "" && keyFile != "" {
		creds, err := credentials.NewServerTLSFromFile(cert, keyFile)
		if err != nil {
			glog.Fatalf("fail to credentials, %+v", err)
		}

		glog.Infoln(address + " net.Listing...")

		options = append(options, grpc.Creds(creds))
	} else {
		glog.Infoln(address + " net.Listing...")
	}

	s = grpc.NewServer(options...)
	f(s)

	err = s.Serve(listener)
	if err != nil {
		glog.Fatalf("grpcServer.Serve err: %v", err)
	}

	return nil
}
