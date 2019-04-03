package main

import (
	"context"
	"log"
	"time"

	"strconv"

	pb "github.com/dcache/grpc/server"
	"google.golang.org/grpc"
)

func testAdd(s pb.SServiceClient, i int) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	kv := &pb.Cacheitem{Key: "key" + strconv.Itoa(i), Value: "value" + strconv.Itoa(i)}
	s.Add(ctx, kv)

}
func testGet(s pb.SServiceClient, i int) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	key := &pb.Key{Key: "key" + strconv.Itoa(i)}
	ci, err := s.Get(ctx, key)
	if err == nil {
		//println("err is nil, no error found")
		println("find key-value pairs: ", ci.Key, ci.Value)
	} else {
		log.Fatalf("error get: %v", err)
	}

}

func main() {

	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connect: %v", err)
		return
	}

	defer conn.Close()

	s := pb.NewSServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	r, err := s.GetServers(ctx, &pb.Client{Host: "a client access through rpc"})
	if err != nil {
		log.Fatalf("error getservers: %v", err)

	}
	for _, server := range r.S {
		println(server.Host, server.Port)
	}

	for i := 0; i < 1000; i++ {
		go testAdd(s, i)
		go testGet(s, i)
	}

	time.Sleep(10 * time.Second)
}
