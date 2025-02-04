.PHONY: kitex-auth

kitex-auth:
	@cp idl/auth.proto app/auth && cd app/auth && cwgo server  --type RPC  --idl auth.proto  --server_name auth --module github.com/ -I . && rm -rf auth.proto