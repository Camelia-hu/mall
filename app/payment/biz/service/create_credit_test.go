package service

import (
	"context"
	payment "github.com/Camelia-hu/mall/rpc_gen/kitex_gen/payment"
	"testing"
)

func TestCreateCredit_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateCreditService(ctx)
	// init req and assert value

	req := &payment.CreateCreditReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
