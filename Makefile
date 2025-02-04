.PHONY: kitex-auth

kitex-auth-client:
	@cd rpc_gen && \
	cwgo client  --type RPC  --idl ../idl/auth.proto  --server_name auth --module github.com/Camelia-hu/mall/rpc_gen -I ../idl


kitex-auth-server:
	@cp idl/auth.proto app/auth && \
	cd app/auth && \
	cwgo server  --type RPC  --idl auth.proto  --server_name auth --pass "-use github.com/Camelia-hu/mall/rpc_gen" --module github.com/Camelia-hu/mall/idl/auth -I . \
	&& rm -rf auth.proto

kitex-cart-client:
	@cd rpc_gen && \
	cwgo client  --type RPC  --idl ../idl/cart.proto  --server_name cart --module github.com/Camelia-hu/mall/rpc_gen -I ../idl


kitex-cart-server:
	@cp idl/cart.proto app/cart && \
	cd app/cart && \
	cwgo server  --type RPC  --idl cart.proto  --server_name cart --pass "-use github.com/Camelia-hu/mall/rpc_gen" --module github.com/Camelia-hu/mall/idl/cart -I . \
	&& rm -rf cart.proto

kitex-checkout-client:
	@cd rpc_gen && \
	cwgo client  --type RPC  --idl ../idl/checkout.proto  --server_name checkout --module github.com/Camelia-hu/mall/rpc_gen -I ../idl


kitex-checkout-server:
	@cp idl/checkout.proto app/checkout && \
	cd app/checkout && \
	cwgo server  --type RPC  --idl checkout.proto  --server_name checkout --pass "-use github.com/Camelia-hu/mall/rpc_gen" --module github.com/Camelia-hu/mall/idl/checkout -I . \
	&& rm -rf checkout.proto

kitex-order-client:
	@cd rpc_gen && \
	cwgo client  --type RPC  --idl ../idl/order.proto  --server_name order --module github.com/Camelia-hu/mall/rpc_gen -I ../idl


kitex-order-server:
	@cp idl/order.proto app/order && \
	cd app/order && \
	cwgo server  --type RPC  --idl order.proto  --server_name order --pass "-use github.com/Camelia-hu/mall/rpc_gen" --module github.com/Camelia-hu/mall/idl/order -I . \
	&& rm -rf order.proto

kitex-payment-client:
	@cd rpc_gen && \
	cwgo client  --type RPC  --idl ../idl/payment.proto  --server_name payment --module github.com/Camelia-hu/mall/rpc_gen -I ../idl


kitex-payment-server:
	@cp idl/payment.proto app/payment && \
	cd app/payment && \
	cwgo server  --type RPC  --idl payment.proto  --server_name payment --pass "-use github.com/Camelia-hu/mall/rpc_gen" --module github.com/Camelia-hu/mall/idl/payment -I . \
	&& rm -rf payment.proto

kitex-product-client:
	@cd rpc_gen && \
	cwgo client  --type RPC  --idl ../idl/product.proto  --server_name product --module github.com/Camelia-hu/mall/rpc_gen -I ../idl


kitex-product-server:
	@cp idl/product.proto app/product && \
	cd app/product && \
	cwgo server  --type RPC  --idl product.proto  --server_name product --pass "-use github.com/Camelia-hu/mall/rpc_gen" --module github.com/Camelia-hu/mall/idl/product -I . \
	&& rm -rf product.proto

kitex-user-client:
	@cd rpc_gen && \
	cwgo client  --type RPC  --idl ../idl/user.proto  --server_name user --module github.com/Camelia-hu/mall/rpc_gen -I ../idl


kitex-user-server:
	@cp idl/user.proto app/user && \
	cd app/user && \
	cwgo server  --type RPC  --idl user.proto  --server_name user --pass "-use github.com/Camelia-hu/mall/rpc_gen" --module github.com/Camelia-hu/mall/idl/user -I . \
	&& rm -rf user.proto