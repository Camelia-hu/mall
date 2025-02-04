package main

import (
	"context"
	"github.com/Camelia-hu/mall/idl/payment/biz/service"
	payment "github.com/Camelia-hu/mall/rpc_gen/kitex_gen/payment"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// CreateCredit implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) CreateCredit(ctx context.Context, req *payment.CreateCreditReq) (resp *payment.CreateCreditResp, err error) {
	resp, err = service.NewCreateCreditService(ctx).Run(req)

	return resp, err
}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	resp, err = service.NewChargeService(ctx).Run(req)

	return resp, err
}
