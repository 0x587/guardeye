FROM golang:1.22.7-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY report-gateway/etc /app/etc
RUN go build -ldflags="-s -w" -tags no_k8s -o /app/gateway ./report-gateway


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
COPY --from=ghcr.io/tarampampam/curl:8.6.0 /bin/curl /bin/curl
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/gateway /app/gateway
COPY --from=builder /app/etc /app/etc

CMD ["./gateway", "-f", "etc/gateway.yaml"]