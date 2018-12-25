default:
	@echo 'Usage of make: [ build | linux_build | windows_build | clean ]'

build: 
	@go build -ldflags "-X main.VERSION=1.0.0 -X main.GIT_HASH=`git rev-parse HEAD` -s" -o ./bin/etcd-manage ./

linux_build: 
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.VERSION=1.0.0 -X main.GIT_HASH=`git rev-parse HEAD` -s" -o ./bin/etcd-manage ./

windows_build: 
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-X main.VERSION=1.0.0 -X main.GIT_HASH=`git rev-parse HEAD` -s" -o ./bin/etcd-manage.exe ./

run: build
	@./bin/etcd-manage

install: build
	@mv ./bin/etcd-manage $(GOPATH)/bin/etcd-manage

build_web:
	cd static && npm run build && cp -r dist ../tpls && cd ../tpls && ./compile.sh

clean: 
	@rm -f ./bin/etcd-manage*

.PHONY: default build linux_build windows_build clean