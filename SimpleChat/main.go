package main

import (
	"fmt"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/sharif0rifat/GoChat/SimpleChat/pb"
	"github.com/sharif0rifat/GoChat/SimpleChat/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Server")
	log := hclog.Default()
	gs := grpc.NewServer()
	cs := server.NewCurrency(log)
	pb.RegisterCurrencyServer(gs, cs)
	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}
	gs.Serve(l)
}
