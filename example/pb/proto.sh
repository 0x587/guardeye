mkdir -p go js python java ts openapi && \
docker run --rm --user $(id -u) -v $(pwd):$(pwd) -w $(pwd) rvolosatovs/protoc \
  --go_out=go --go-grpc_out=go \
  --js_out=import_style=commonjs:js --grpc-web_out=import_style=commonjs,mode=grpcwebtext:js \
  --ts_out=service=grpc-web:ts \
  --python_out=python \
  --java_out=java \
  --openapiv2_out=openapi \
  -I=. ./*.proto