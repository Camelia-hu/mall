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
	"log"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// Finish your business logic.
	span := trace.SpanFromContext(s.ctx)
	if !span.SpanContext().IsValid() {
		return nil, errors.New("span invalid")
	}
	myClaims := new(module.MyClaims)
	token, err := jwt.ParseWithClaims(req.Token, myClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.GetConf().Jwt.Key), nil
	})
	if err != nil {
		span.SetStatus(codes.Error, "token parse err")
		log.Println("1 ", err)
		return nil, err
	}
	if !token.Valid {
		span.SetStatus(codes.Error, "token expired")
		return nil, errors.New("token 过期喵～")
	}
	resp = &auth.VerifyResp{}
	resp.Res = true

	return resp, nil
	return
}
