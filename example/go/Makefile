all: clean
	cp -r ../pb ./ && \
	cd pb && protoc --proto_path=. --go_out=. --go-grpc_out=. ./*.proto

clean:
	rm -rf pb