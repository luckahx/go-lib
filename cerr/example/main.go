package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/luckahx/go-lib/cerr"
	"github.com/luckahx/go-lib/cerr/example/models/ping"
	"github.com/luckahx/go-lib/cerr/proto/protocerr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type pServe struct{}

func (ps *pServe) Ping(ctx context.Context, req *ping.PingReq) (*ping.PingResponse, error) {
	return &ping.PingResponse{
		Err:  protocerr.ToProto(l3()),
		Data: "Pong:" + req.Data,
	}, nil
}

func l1() error {
	err := errors.New("file not found")

	return cerr.Wrap(err, "l1_err", "l1 error occured")
}

func l2() error {
	return cerr.Wrap(l1(), "l2_err", "l2 error occured")
}

func l3() error {
	return cerr.Wrap(l2(), "l3_err", "l3 error occured")
}

func main() {

	e1 := cerr.Wrap(nil, "test", "test")
	if e1 == nil {
		log.Println("e1=nil")
	}
	log.Println(e1)

	switch code, theErr := e1.HasFirst("l2_err", "l1_err"); code {
	case "l1_err":
		log.Println(theErr.Error())
	case "l2_err":
		log.Println("has l2 err")
	}

	if err := l3(); err != nil {
		if fe, ok := err.(cerr.CError); ok {
			log.Println(fe.FullMessage())
			// pretty.Println(fe.FullErrorStack())
		}

		j, _ := json.Marshal(err)
		log.Println(string(j))

		outMsg := cerr.CErrorMessage{}
		err := json.Unmarshal(j, &outMsg)
		if err != nil {
			panic(err)
		}

		// pretty.Println(outMsg)
		log.Println(outMsg.FullMessage())
	}
	return

	lis, err := net.Listen("tcp", ":3002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	ping.RegisterPingServer(grpcServer, &pServe{})

	go testClient()

	grpcServer.Serve(lis)

	waitInterupt()
}

func testClient() {
	time.Sleep(2 * time.Second)
	conn, err := grpc.Dial("127.0.0.1:3002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := ping.NewPingClient(conn)

	// add authentication token to meta data
	md := metadata.New(map[string]string{"Authorization": "token test123"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := client.Ping(ctx, &ping.PingReq{Data: "abc"})
	// pretty.Println(res)
	log.Println(res.Err.FullMessage(), err)
	check(err, "ping call")
}

func check(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %+v", msg, err)
	}
}

func waitInterupt() {
	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan
	fmt.Println()
}
