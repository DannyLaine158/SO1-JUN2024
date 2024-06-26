package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"serverGRPC/kafka"
	"serverGRPC/model"
	pb "serverGRPC/server"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	tweet := model.Data{
		Texto: in.GetTexto(),
		Pais:  in.GetPais(),
	}

	fmt.Println(tweet)

	kafka.Produce(tweet)

	return &pb.ReplyInfo{Info: "Hola cliente, recibí el album"}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":3001")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
