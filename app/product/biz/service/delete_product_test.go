package service

import (
	"context"
	product "github.com/Camelia-hu/mall/rpc_gen/kitex_gen/product"
	"testing"
)

func TestDeleteProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteProductService(ctx)
	// init req and assert value

	req := &product.DeleteReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
