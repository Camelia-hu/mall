.PHONY: kitex-auth

kitex-auth-server:
	@cp idl/auth.proto app/auth && \
	cd app/auth && \
	cwgo server  --type RPC  --idl auth.proto  --server_name auth --pass "-use github.com/Camelia-hu/mall/rpc_gen" --module github.com/Camelia-hu/mall/idl/auth -I . \
	&& rm -rf auth.proto