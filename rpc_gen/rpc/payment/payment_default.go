package payment

import (
	"context"
	payment "github.com/Camelia-hu/mall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func CreateCredit(ctx context.Context, req *payment.CreateCreditReq, callOptions ...callopt.Option) (resp *payment.CreateCreditResp, err error) {
	resp, err = defaultClient.CreateCredit(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateCredit call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Charge(ctx context.Context, req *payment.ChargeReq, callOptions ...callopt.Option) (resp *payment.ChargeResp, err error) {
	resp, err = defaultClient.Charge(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Charge call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
