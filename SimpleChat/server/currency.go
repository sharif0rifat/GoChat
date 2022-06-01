package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/sharif0rifat/GoChat/SimpleChat/pb"
)

type CurrencyServer struct {
	log hclog.Logger
	pb.UnimplementedCurrencyServer
}

func NewCurrency(l hclog.Logger) *CurrencyServer {
	return &CurrencyServer{l, pb.UnimplementedCurrencyServer{}}
}

func (c *CurrencyServer) GetRate(ctx context.Context, rr *pb.RateRequest) (*pb.RateResponse, error) {
	c.log.Info("Handle get rate", "base", rr.GetBase(), "destination", rr.GetDestination())
	return &pb.RateResponse{Rate: 0.5}, nil
}

//func (c *UnimplementedCurrencyServer) mustEmbedUnimplementedCurrencyServer() {}

// func (c *CurrencyServer) mustEmbedUnimplementedCurrencyServer() {
// 	c.log.Info("UnImplemented")
// }
