package service

import (
	"context"
	auth "github.com/Camelia-hu/mall/rpc_gen/kitex_gen/auth"
	"testing"
)

func TestReFreshTokenByRPC_Run(t *testing.T) {
	ctx := context.Background()
	s := NewReFreshTokenByRPCService(ctx)
	// init req and assert value

	req := &auth.RefreshReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
