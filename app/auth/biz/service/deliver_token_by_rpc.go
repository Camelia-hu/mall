package service

import (
	"context"
	"errors"
	"github.com/Camelia-hu/mall/idl/auth/conf"
	"github.com/Camelia-hu/mall/idl/auth/module"
	auth "github.com/Camelia-hu/mall/rpc_gen/kitex_gen/auth"
	"github.com/golang-jwt/jwt/v5"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"time"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// Finish your business logic.
	span := trace.SpanFromContext(s.ctx)
	if !span.SpanContext().IsValid() {
		return nil, errors.New("span invalid")
	}
	myAccessClaims := module.MyClaims{
		Id: int(req.UserId),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * 60 * time.Minute)),
		},
	}
	myRefreshClaims := module.MyClaims{
		Id: int(req.UserId),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
		},
	}
	resp = &auth.DeliveryResp{}
	resp.AccessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, myAccessClaims).SignedString([]byte(conf.GetConf().Jwt.Key))
	resp.RefreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, myRefreshClaims).SignedString([]byte(conf.GetConf().Jwt.Key))
	if err != nil {
		span.SetStatus(codes.Error, "token deliver err")
		return nil, err
	}
	return resp, nil
	return
}
