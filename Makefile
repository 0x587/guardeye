docker-build:
	docker build -f report/Dockerfile -t registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:report_$(version) . && \
	docker build -f report-gateway/Dockerfile -t registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:report_gateway_$(version) . && \
	docker build -f wsapi/Dockerfile -t registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:wsapi_$(version) . && \
	docker build -f api/Dockerfile -t registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:api_$(version) .

docker-push:
	docker push registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:report_$(version) && \
	docker push registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:report_gateway_$(version) && \
	docker push registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:wsapi_$(version) && \
	docker push registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:api_$(version)

docker: docker-build docker-push