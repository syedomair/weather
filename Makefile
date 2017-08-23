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
	@$(GOBUILD) main.go
	@mv main $(BUILDPATH)/build 

build: clean_build copy_build 

run:
	@$(GORUN) main.go -config="config/config_test.yml"

test: 
	rm -f ./tmp/* 
	go test ./... -cover

check:
	go vet ./...
	go fmt ./...
	golint ./...

race: 
	rm -f ./tmp/* 
	go test ./... -race

#test:
#	go test -v unit_test.go
