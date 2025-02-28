goctl api plugin -plugin goctl-swagger="swagger -filename apiclient/api.json" -api api.api -dir . && \
for l in go javascript typescript-node nodejs-es5 python; do
    docker run --rm -v "$(pwd)/apiclient:/workdir" swaggerapi/swagger-codegen-cli generate \
      -i "/workdir/api.json" \
      -l "$l" \
      -o "/workdir/$l"
  done