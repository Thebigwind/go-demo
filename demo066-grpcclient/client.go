package main

import (
	"context"
	"google.golang.org/grpc"
	"lch-tests/protobufs/workload"
	"log"
	"time"
)

var (
	WorkloadClient workload.WorkloadServiceClient
)

type Conn struct {
	clientConn  *grpc.ClientConn
	callTimeOut time.Duration
}

func InitConnections(addr string) {
	cc := GetConn(addr)
	WorkloadClient = workload.NewWorkloadServiceClient(cc)
}

func GetConn(addr string) *grpc.ClientConn {
	c := &Conn{
		callTimeOut: 3 * time.Second,
	}
	c.Dial(addr)
	return c.clientConn
}

func (c *Conn) Dial(addr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.callTimeOut)
	defer cancel()

	var err error
	c.clientConn, err = grpc.DialContext(ctx, addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("dial rpc error:%s", err.Error())
	}
	return nil
}
