package main

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/LoptrThor/middle_trpc_go"
	"go.uber.org/mock/gomock"
	_ "trpc.group/trpc-go/trpc-go/http"
)

//go:generate go mod tidy
//go:generate mockgen -destination=stub/github.com/LoptrThor/middle_trpc_go/middle_proto_mock.go -package=middle_trpc_go -self_package=github.com/LoptrThor/middle_trpc_go --source=stub/github.com/LoptrThor/middle_trpc_go/middle_proto.trpc.go

func Test_accountImpl_OpenCheck(t *testing.T) {
	// Start writing mock logic.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	accountService := pb.NewMockAccountService(ctrl)
	var inorderClient []*gomock.Call
	// Expected behavior.
	m := accountService.EXPECT().OpenCheck(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
		s := &accountImpl{}
		return s.OpenCheck(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// Start writing unit test logic.
	type args struct {
		ctx context.Context
		req *pb.Request
		rsp *pb.Reply
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rsp *pb.Reply
			var err error
			if rsp, err = accountService.OpenCheck(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("accountImpl.OpenCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("accountImpl.OpenCheck() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}
