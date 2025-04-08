all: clean
	cp -r ../pb ./ && \
	pip install guardeye_python_sdk grpcio grpcio-tools && \
	cd pb && \
	python3 -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. ./*.proto && \
	cp -r ./*.py ../

clean:
	rm -rf pb