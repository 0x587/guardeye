all:
	goctl api plugin -plugin goctl-swagger="swagger -filename apiclient/api.json" -api api.api -dir . && \
	for l in go javascript php; do \
      docker run --rm -v "$(pwd):/go-work" swaggerapi/swagger-codegen-cli generate \
        -i "/go-work/apiclient/api" \
        -l "$l" \
        -o "/go-work/apiclient/$l"
    done