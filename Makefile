NAME			:= hello
VERSION		:= v0.1.0
REVISION	:= $(shell git rev-parse --short HEAD)

SRCS		:= $(shell find . -type f -name '*.go')
LDFLAGS	:= -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""

.DEFAULT_GOAL := bin/$(NAME)

.PHONY: dep-install
dep-install:
ifeq ($(shell command -v dep 2> /dev/null),)
	go get -u github.com/golang/dep/...
endif

.PHONY: init
init:
	dep init

.PHONY: deps
deps: dep
	dep ensure -update

bin/$(NAME): $(SRCS)
	go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$(NAME)

.PHONY: install
install:
	go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o ${GOPATH}/bin/$(NAME)
	#go install $(LDFLAGS) -o hello
