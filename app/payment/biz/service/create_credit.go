package service

import (
	"context"
	payment "github.com/Camelia-hu/mall/rpc_gen/kitex_gen/payment"
)

type CreateCreditService struct {
	ctx context.Context
} // NewCreateCreditService new CreateCreditService
func NewCreateCreditService(ctx context.Context) *CreateCreditService {
	return &CreateCreditService{ctx: ctx}
}

// Run create note info
func (s *CreateCreditService) Run(req *payment.CreateCreditReq) (resp *payment.CreateCreditResp, err error) {
	// Finish your business logic.

	return
}
