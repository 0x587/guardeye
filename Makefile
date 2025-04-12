docker-build-one:
	docker build -f $(name)/Dockerfile -t registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:$(name)_$(version) .

docker-build:
	docker build -f report/Dockerfile -t registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:report_$(version) . && \
	docker build -f report-gateway/Dockerfile -t registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:report_gateway_$(version) . && \
	docker build -f wsapi/Dockerfile -t registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:wsapi_$(version) . && \
	docker build -f api/Dockerfile -t registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:api_$(version) . && \
	docker build -f link/Dockerfile -t registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:link_$(version) . && \
	docker build -f foxglove_cdrservice/Dockerfile -t registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:foxglove_cdrservice_$(version) ./foxglove_cdrservice

docker-push-one:
	docker push registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:$(name)_$(version)

docker-push:
	docker push registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:report_$(version) && \
	docker push registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:report_gateway_$(version) && \
	docker push registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:wsapi_$(version) && \
	docker push registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:api_$(version) && \
	docker push registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:link_$(version) && \
	docker push registry.cn-shenzhen.aliyuncs.com/shawnsiu/guardeye:foxglove_cdrservice_$(version)

docker-one: docker-build-one docker-push-one

docker: docker-build docker-push