proto:
	goctl rpc protoc report.proto --go_out=.  --go-grpc_out=.  --zrpc_out=. && \
	protoc --descriptor_set_out=../report-gateway/etc/report.pb report.proto