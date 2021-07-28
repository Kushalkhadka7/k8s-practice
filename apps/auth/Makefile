SHELL :=/bin/bash
CWD := $(PWD)

build-dev:
	@docker build \
	-t kushal/auth \
    --target dev .
build:
	@docker build \
	-t kushal/auth .
run:
	@docker run \
	-p 4000:4000 \
  	--env-file=.env \
  	kushal/auth
build-app:
	@exec go build

.PHONY: all
