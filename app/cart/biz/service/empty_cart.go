package service

import (
	"context"
	"errors"
	"github.com/Camelia-hu/mall/idl/cart/biz/dal/mysql"
	"github.com/Camelia-hu/mall/idl/cart/biz/dal/redis"
	"github.com/Camelia-hu/mall/idl/cart/module"
	cart "github.com/Camelia-hu/mall/rpc_gen/kitex_gen/cart"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"strconv"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	// Finish your business logic.
	span := trace.SpanFromContext(s.ctx)
	uid := strconv.Itoa(int(req.UserId))
	result, err := redis.RedisClient.Del(s.ctx, uid).Result()
	if err != nil || result <= 0 {
		span.SetStatus(codes.Error, "redis delete err")
		return nil, errors.New("redis delete err")
	}
	err = mysql.DB.Where("uid = ?", req.UserId).Delete(&module.Cart{}).Error
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	resp = &cart.EmptyCartResp{}
	return resp, nil
}
