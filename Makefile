default:
	@echo 'Usage of make: [ build | linux_build | windows_build | build_web | clean ]'

build: 
	@go build -o ./bin/etcd-manage ./

linux_build: 
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/etcd-manage ./

windows_build: 
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/etcd-manage.exe ./

docker_build: linux_build
	docker build -t shiguanghuxian/etcd-manage:1 .

docker_run: docker_build
	docker-compose up --force-recreate

docker_cluster_run: docker_build
	docker-compose -f docker-compose-cluster.yml up --force-recreate

run: build
	@./bin/etcd-manage

install: build
	@mv ./bin/etcd-manage $(GOPATH)/bin/etcd-manage

build_web:
	cd static && npm run build && cp -r dist ../tpls && cd ../tpls && ./compile.sh

clean: 
	@rm -f ./bin/etcd-manage*
	@rm -f ./bin/logs/*

.PHONY: default build linux_build windows_build clean