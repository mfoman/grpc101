package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	t "time"

	"github.com/mfoman/grpc101/tcp"
	"github.com/mfoman/grpc101/time"

	"google.golang.org/grpc"
)

func main() {
	// Creat a virtual RPC Client Connection on port  9080 WithInsecure (because  of http)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	// Defer means: When this function returns, call this method (meaing, one main is done, close connection)
	defer conn.Close()

	TcpLoop(conn)
}

func TcpLoop(conn *grpc.ClientConn) {
	//  Create new Client from generated gRPC code from proto
	c := tcp.NewTcpMessagingClient(conn)

	s1 := rand.NewSource(t.Now().UnixNano())
    r1 := rand.New(s1)
	seq := int64(r1.Intn(1_000_000) + 1_000_000)

	SendTcpPackage(c, &tcp.Tcp{
		Source:   "127.0.0.1",
		Dest:     "127.0.0.1",
		Seq:      seq,
		Ack:      0,
		Offset:   "",
		Reserved: "",
		Flags:    "SYN",
		Window:   "",
		Checksum: "",
		Urgentp:  "",
		Options:  "",
		Data:     "",
	})
}

func GetTimeLoop(conn *grpc.ClientConn) {
	//  Create new Client from generated gRPC code from proto
	c := time.NewGetCurrentTimeClient(conn)

	for {
		SendGetTimeRequest(c)
		t.Sleep(5 * t.Second)
	}
}

func SendTcpPackage(c tcp.TcpMessagingClient, message *tcp.Tcp) {
	res, err := c.SendMessage(context.Background(), message)

	if err != nil {
		log.Fatalf("Error when calling SendMessage %s", err)
	}

	fmt.Printf("Client: SYN+ACK reseived %s\n", res)

	if res.Flags == "SYN+ACK" {
		res, err = c.SendMessage(context.Background(), &tcp.Tcp{
			Source:   "127.0.0.1",
			Dest:     "127.0.0.1",
			Seq:      res.Ack,
			Ack:      res.Seq + 1,
			Offset:   "",
			Reserved: "",
			Flags:    "ACK",
			Window:   "",
			Checksum: "",
			Urgentp:  "",
			Options:  "",
			Data:     "TCP HANDSHAKE COMPLETE!",
		})

		if err != nil {
			log.Fatalf("Error when calling SendMessage %s", err)
		}

		fmt.Printf("Client: ACK received %s\n", res)
	}
}

func SendGetTimeRequest(c time.GetCurrentTimeClient) {
	// Between the curly brackets are nothing, because the .proto file expects no input.
	message := time.GetTimeRequest{}

	response, err := c.GetTime(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetTime: %s", err)
	}

	fmt.Printf("Current time right now: %s\n", response.Reply)
}
