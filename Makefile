BUILDPATH=${CURDIR}
GO=$(shell which go)
GOBUILD=${GO} build
GORUN=${GO} run

clean_build:
	@echo "Cleaning ...."
	@rm -rf $(BUILDPATH)/build

copy_build:
	@echo "Building ...."
	@if [ ! -d "$(BUILDPATH)/build" ] ; then mkdir -p $(BUILDPATH)/build ; fi 
	@$(GOBUILD) .
	@mv weather $(BUILDPATH)/build 

build: clean_build copy_build 

run:    build
	build/weather -config="config/config_test.yml"

docker: build
	docker build -t app .

test: 
	rm -f ./tmp/* 
	go test -v ./... -cover

check:
	go vet ./...
	go fmt ./...
	golint ./...

race: 
	rm -f ./tmp/* 
	go test ./... -race

#test:
#	go test -v unit_test.go
