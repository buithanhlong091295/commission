SRC_PATH:= ${GOPATH}/src/xtek/exchange/commission

proto-gen:
	@echo "Protobuf generate"
	${SRC_PATH}/deployment/scripts/protobuf-gen.sh
build:
	@echo "Building to binary execute server"
	go build -o commission_srv ${SRC_PATH}/cmd/srv/...

build-alpine:
	CGO_ENABLED=0 GOOS=linux go build -o commission_srv ${SRC_PATH}/cmd/srv/...
	
docker-dev:
	@docker-compose -p citicoins_dev -f ${SRC_PATH}/deployment/docker/docker-compose.dev.yml up -d --build
docker-prod:
	@docker-compose -p owifi_prod -f ./devops/docker/docker-compose.yml up -d --build
run:
	go run ${SRC_PATH}/cmd/srv/...
install:
	@echo "install to binary execute server"
	go install ${SRC_PATH}/cmd/srv/...
start:
	./commission_srv start
build-image:
	docker build -t xtekexchange/commission-service .
push-image:
	docker push xtekexchange/commission-service
docker-deploy:
	make build-image
	make push-image

deploy-image:
	${SRC_PATH}/deployment/scripts/deploy-image.sh