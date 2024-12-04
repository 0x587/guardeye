FROM golang:1.22.7 as builder

ENV GOPROXY=https://goproxy.cn,direct

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -tags no_k8s -o /app/client ./test-client

FROM scratch

COPY --from=builder /app/client /app/client

# 设置默认命令
CMD ["/app/client"]