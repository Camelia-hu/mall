// Code generated by Kitex v0.9.1. DO NOT EDIT.

package paymentservice

import (
	"context"
	"errors"
	payment "github.com/Camelia-hu/mall/rpc_gen/kitex_gen/payment"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreateCredit": kitex.NewMethodInfo(
		createCreditHandler,
		newCreateCreditArgs,
		newCreateCreditResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Charge": kitex.NewMethodInfo(
		chargeHandler,
		newChargeArgs,
		newChargeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	paymentServiceServiceInfo                = NewServiceInfo()
	paymentServiceServiceInfoForClient       = NewServiceInfoForClient()
	paymentServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return paymentServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return paymentServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return paymentServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "PaymentService"
	handlerType := (*payment.PaymentService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "payment",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func createCreditHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(payment.CreateCreditReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(payment.PaymentService).CreateCredit(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CreateCreditArgs:
		success, err := handler.(payment.PaymentService).CreateCredit(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateCreditResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCreateCreditArgs() interface{} {
	return &CreateCreditArgs{}
}

func newCreateCreditResult() interface{} {
	return &CreateCreditResult{}
}

type CreateCreditArgs struct {
	Req *payment.CreateCreditReq
}

func (p *CreateCreditArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(payment.CreateCreditReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateCreditArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateCreditArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateCreditArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CreateCreditArgs) Unmarshal(in []byte) error {
	msg := new(payment.CreateCreditReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateCreditArgs_Req_DEFAULT *payment.CreateCreditReq

func (p *CreateCreditArgs) GetReq() *payment.CreateCreditReq {
	if !p.IsSetReq() {
		return CreateCreditArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateCreditArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateCreditArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateCreditResult struct {
	Success *payment.CreateCreditResp
}

var CreateCreditResult_Success_DEFAULT *payment.CreateCreditResp

func (p *CreateCreditResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(payment.CreateCreditResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateCreditResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateCreditResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateCreditResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CreateCreditResult) Unmarshal(in []byte) error {
	msg := new(payment.CreateCreditResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateCreditResult) GetSuccess() *payment.CreateCreditResp {
	if !p.IsSetSuccess() {
		return CreateCreditResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateCreditResult) SetSuccess(x interface{}) {
	p.Success = x.(*payment.CreateCreditResp)
}

func (p *CreateCreditResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateCreditResult) GetResult() interface{} {
	return p.Success
}

func chargeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(payment.ChargeReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(payment.PaymentService).Charge(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ChargeArgs:
		success, err := handler.(payment.PaymentService).Charge(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ChargeResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newChargeArgs() interface{} {
	return &ChargeArgs{}
}

func newChargeResult() interface{} {
	return &ChargeResult{}
}

type ChargeArgs struct {
	Req *payment.ChargeReq
}

func (p *ChargeArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(payment.ChargeReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ChargeArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ChargeArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ChargeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ChargeArgs) Unmarshal(in []byte) error {
	msg := new(payment.ChargeReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ChargeArgs_Req_DEFAULT *payment.ChargeReq

func (p *ChargeArgs) GetReq() *payment.ChargeReq {
	if !p.IsSetReq() {
		return ChargeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChargeArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ChargeArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ChargeResult struct {
	Success *payment.ChargeResp
}

var ChargeResult_Success_DEFAULT *payment.ChargeResp

func (p *ChargeResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(payment.ChargeResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ChargeResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ChargeResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ChargeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ChargeResult) Unmarshal(in []byte) error {
	msg := new(payment.ChargeResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChargeResult) GetSuccess() *payment.ChargeResp {
	if !p.IsSetSuccess() {
		return ChargeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChargeResult) SetSuccess(x interface{}) {
	p.Success = x.(*payment.ChargeResp)
}

func (p *ChargeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChargeResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateCredit(ctx context.Context, Req *payment.CreateCreditReq) (r *payment.CreateCreditResp, err error) {
	var _args CreateCreditArgs
	_args.Req = Req
	var _result CreateCreditResult
	if err = p.c.Call(ctx, "CreateCredit", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Charge(ctx context.Context, Req *payment.ChargeReq) (r *payment.ChargeResp, err error) {
	var _args ChargeArgs
	_args.Req = Req
	var _result ChargeResult
	if err = p.c.Call(ctx, "Charge", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
