package server

import (
	"context"
	"errors"
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	"github.com/dcache/cache"
	"github.com/dcache/server"
	"google.golang.org/grpc"
	//	pb "github.com/dcache/grpc/server"
)

type GrpcServer struct {
	ct   *cache.CacheTable
	port string
}

func NewGrpcServer(port string, ctable *cache.CacheTable) *GrpcServer {
	return &GrpcServer{port: port, ct: ctable}
}

func (gs *GrpcServer) GetServers(ctx context.Context, in *Client) (out *Servers, err error) {
	nodes := server.GetServerNodes()
	var ss = new(Servers)
	for _, node := range nodes {
		s1 := &Server{Host: node.Host, Port: node.Port}
		ss.S = append(ss.S, s1)
	}
	return ss, nil

}
func (gs *GrpcServer) Add(ctx context.Context, in *Cacheitem) (out *Cacheitem, err error) {
	ci := cache.NewCacheItem(in.Key, in.Value)
	gs.ct.Add(ci)
	return in, nil
}
func (gs *GrpcServer) Update(ctx context.Context, in *Cacheitem) (out *Cacheitem, err error) {
	ci := cache.NewCacheItem(in.Key, in.Value)
	gs.ct.Add(ci)
	return in, nil
}

func (gs *GrpcServer) Delete(ctx context.Context, in *Key) (out *Cacheitem, err error) {
	gs.ct.RemoveCacheItem(in.Key)

	return &Cacheitem{Key: "", Value: ""}, nil
}

func (gs *GrpcServer) Get(ctx context.Context, in *Key) (out *Cacheitem, err error) {
	println("get key:", in.Key)
	ci := gs.ct.Get(in.Key)
	if ci != nil {
		println("found key:", ci.Key.(string), ci.Value.(string))
		ciout := &Cacheitem{Key: ci.Key.(string), Value: ci.Value.(string)}
		return ciout, nil
	}
	return nil, errors.New("key not found")
}

func (gs *GrpcServer) Init() {
	lis, err := net.Listen("tcp", gs.port)
	if err != nil {
		log.Fatalf("error listen:%v", err)
	}

	s := grpc.NewServer()
	RegisterSServiceServer(s, gs)
	reflection.Register(s)

	if err := s.Serve(lis); lis != nil {
		log.Fatalf("error serve: %v", err)
	}
}
