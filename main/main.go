package main

import (
	pb "github.com/LoptrThor/middle_trpc_go"
	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	_ "trpc.group/trpc-go/trpc-filter/recovery"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	s := trpc.NewServer()
	pb.RegisterAccountService(s.Service("trpc.app.server.Account"), &accountImpl{})
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
