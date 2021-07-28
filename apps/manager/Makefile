SHELL :=/bin/bash
CWD := $(PWD)

build-dev:
	@docker build \
	-t kushal/manager \
    --target dev .
build:
	@docker build \
	-t kushal/manager .
run:
	@docker run \
	-p 5000:5000 \
  	--env-file=.env \
  	kushal/manager

.PHONY: all
