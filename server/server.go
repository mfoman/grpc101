package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
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
	s1 := rand.NewSource(t.Now().UnixNano())
    r1 := rand.New(s1)
	seq := int64(r1.Intn(1_000_000) + 1_000_000)

	message := tcp.Tcp {
		Source: "127.0.0.1",
		Dest: "127.0.0.1",
		Seq: 0,
		Ack: 0,
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
		message.Seq = seq
		message.Ack = in.Seq + 1
		message.Data = "PIGGYBACKING"
	case "ACK":
		fmt.Printf("Server: ACK received %s\n", in)
		message.Flags = "ACK"
		message.Seq = in.Ack
		message.Ack = in.Seq + 1
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
