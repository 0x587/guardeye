docker-build:
	docker build -f report/Dockerfile -t registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:report_$(version) . && \
	docker build -f report-gateway/Dockerfile -t registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:report_gateway_$(version) .

docker-push:
	docker push registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:report_$(version) && \
	docker push registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:report_gateway_$(version)

docker: docker-build docker-push