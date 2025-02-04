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

type ReFreshTokenByRPCService struct {
	ctx context.Context
} // NewReFreshTokenByRPCService new ReFreshTokenByRPCService
func NewReFreshTokenByRPCService(ctx context.Context) *ReFreshTokenByRPCService {
	return &ReFreshTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *ReFreshTokenByRPCService) Run(req *auth.RefreshReq) (resp *auth.RefreshResp, err error) {
	// Finish your business logic.
	span := trace.SpanFromContext(s.ctx)
	var myClaims module.MyClaims
	refreshToken, err := jwt.ParseWithClaims(req.RefreshToken, &myClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.GetConf().Jwt.Key), nil
	})
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	if !refreshToken.Valid {
		span.SetStatus(codes.Error, "refreshToken 过期 too 喵～")
		return nil, errors.New("refreshToken 过期 too 喵～")
	}
	newrefreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, module.MyClaims{
		Id: myClaims.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
		},
	}).SignedString([]byte(conf.GetConf().Jwt.Key))
	newaccessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, module.MyClaims{
		Id: myClaims.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
		},
	}).SignedString([]byte(conf.GetConf().Jwt.Key))
	resp = &auth.RefreshResp{
		AccessToken:  newaccessToken,
		RefreshToken: newrefreshToken,
	}
	return resp, nil
	return
}
