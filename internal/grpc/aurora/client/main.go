package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"os"
	"time"
)

const (
	address     = "localhost:50051"
	defaultHash = "0x20cc6269dd49cbc796720c56f69ea042f307cc13b865577138239ab5341e7f17"
)

var kp = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := aoa.NewAOAClient(conn)
	hash := defaultHash
	if len(os.Args) > 1 {
		hash = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetTransactionReceipt(ctx, &aoa.ReceiptRequest{Hash: hash})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Print(r)
}
