package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	pb "../protos"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:6001"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)

	fmt.Print("username: ")
	var name string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		name = scanner.Text()
		if len(name) > 0 {
			break
		} else {
			fmt.Print("username: ")
		}
	}

	stream, err := c.Connect(context.Background())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("Failed to recv: %v", err)
			}
			if name != res.GetName() {
				fmt.Println(fmt.Sprintf("%s> %s", res.GetName(), res.GetMessage()))
			}
		}
	}()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		msg := scanner.Text()
		if msg == ":quit" {
			stream.CloseSend()
			return
		}
		stream.Send(&pb.Post{
			Name:    name,
			Message: msg,
		})
	}
}
