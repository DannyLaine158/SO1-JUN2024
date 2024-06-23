package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	pb "serverGRPC/server"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

type Data struct {
	Texto string
	Pais  string
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	tweet := Data{
		Texto: in.GetTexto(),
		Pais:  in.GetPais(),
	}

	fmt.Println(tweet)

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
