 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build

all: run
run: 
	$(GOCMD) run github.com/nickdenardis/go-webservice
build: 
	$(GOBUILD) github.com/nickdenardis/go-webservice
test:
	$(GOCMD) test ./...