package main

import (
	"context"
	"fmt"
	"log"
	"net"
	t "time"

	"github.com/mfoman/grpc101/tcp"
	"github.com/mfoman/grpc101/time"

	"google.golang.org/grpc"
)

type Server struct {
	time.UnimplementedGetCurrentTimeServer
	tcp.UnimplementedTcpMessagingServer
}

func (s *Server) GetTime(ctx context.Context, in *time.GetTimeRequest) (*time.GetTimeReply, error) {
	fmt.Printf("Received GetTime request\n")
	return &time.GetTimeReply{Reply: t.Now().String()}, nil
}

func (s *Server) SendMessage(ctx context.Context, in *tcp.Tcp) (*tcp.Tcp, error) {
	message := tcp.Tcp {
		Source: "you",
		Dest: "me",
		Seq: "random",
		Ack: "",
		Offset: "",
		Reserved: "",
		Flags: "",
		Window: "",
		Checksum: "",
		Urgentp: "",
		Options: "",
		Data: "",
	}

	switch in.Flags {
	case "SYN":
		fmt.Printf("Server: SYN received %s\n", in)
		message.Flags = "SYN+ACK"
	case "AWK":
		fmt.Printf("Server: AWK received %s\n", in)
		message.Flags = "ACK"
	case "FIN":
		fmt.Printf("Server: FIN received %s\n", in)
		message.Flags = "FIN"
	}

	return &message, nil
}

func main() {
	// Create listener tcp on port 9080
	list, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}
	grpcServer := grpc.NewServer()

	time.RegisterGetCurrentTimeServer(grpcServer, &Server{})
	tcp.RegisterTcpMessagingServer(grpcServer, &Server{})

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
