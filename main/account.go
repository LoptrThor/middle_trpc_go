package main

import (
	"context"

	pb "github.com/LoptrThor/middle_trpc_go"
)

type accountImpl struct {
	pb.UnimplementedAccount
}

func (s *accountImpl) OpenCheck(
	ctx context.Context,
	req *pb.Request,
) (*pb.Reply, error) {
	rsp := &pb.Reply{}
	return rsp, nil
}
