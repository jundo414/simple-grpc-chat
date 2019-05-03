package main

import (
	"io"
	"log"
	"net"
	"sync"

	pb "../protos"

	"google.golang.org/grpc"
)

var streams sync.Map

type chatService struct{}

func (chatService) Connect(
	stream pb.ChatService_ConnectServer) error {
	log.Println(" connect", &stream)
	streams.Store(stream, struct{}{})
	defer func() {
		log.Println(" disconnect", &stream)
		streams.Delete(stream)
	}()
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		streams.Range(func(key,
			value interface{}) bool {
			stream := key.(pb.ChatService_ConnectServer)
			stream.Send(&pb.Post{
				Name:    req.GetName(),
				Message: req.GetMessage(),
			})
			return true
		})
	}
}

func main() {
	port := ":6001"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	srv := grpc.NewServer()
	pb.RegisterChatServiceServer(srv, &chatService{})
	log.Printf("start server on port% s\n", port)
	if err := srv.Serve(lis); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
