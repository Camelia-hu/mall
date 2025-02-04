package service

import (
	"context"
	"github.com/Camelia-hu/mall/idl/order/biz/dal/mysql"
	"github.com/Camelia-hu/mall/idl/order/module"
	order "github.com/Camelia-hu/mall/rpc_gen/kitex_gen/order"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"log"
	"strconv"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
	span := trace.SpanFromContext(s.ctx)
	var orders []*module.Order
	err = mysql.DB.Where("uid = ?", req.UserId).Find(&orders).Error
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		log.Println("find one_order list err : ", err)
		return nil, err
	}
	var Orders []*order.Order
	for _, o := range orders {
		oneOrder := &order.Order{
			OrderId:      strconv.Itoa(int(o.ID)),
			UserId:       o.Uid,
			UserCurrency: o.UserCurrency,
			Email:        o.Email,
		}
		//here
		Orders = append(Orders, oneOrder)
	}
	resp = &order.ListOrderResp{}
	resp.Orders = Orders
	return resp, nil
}
