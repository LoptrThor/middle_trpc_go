// Package main is originally generated by trpc-cmdline v1.0.7.
// It is located at `project/cmd/client`.
// Run this file by executing `go run cmd/client/main.go` in the project directory.
package main

import (
	pb "github.com/LoptrThor/middle_trpc_go"
	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/log"
)

func callAccountOpenCheck() {
	proxy := pb.NewAccountClientProxy(
		client.WithTarget("ip://127.0.0.1:8000"),
		client.WithProtocol("trpc"),
	)
	ctx := trpc.BackgroundContext()
	// Example usage of unary client.
	reply, err := proxy.OpenCheck(ctx, &pb.Request{})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	log.Debugf("simple  rpc   receive: %+v", reply)
}

func main() {
	// Load configuration following the logic in trpc.NewServer.
	cfg, err := trpc.LoadConfig(trpc.ServerConfigPath)
	if err != nil {
		panic("load config fail: " + err.Error())
	}
	trpc.SetGlobalConfig(cfg)
	if err := trpc.Setup(cfg); err != nil {
		panic("setup plugin fail: " + err.Error())
	}
	callAccountOpenCheck()
}
