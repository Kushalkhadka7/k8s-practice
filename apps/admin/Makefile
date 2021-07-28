SHELL :=/bin/bash
CWD := $(PWD)

build-dev:
	@docker build \
	-t kushal/admin \
    --target dev .
build:
	@docker build \
	-t kushal/admin .
run:
	@docker run \
	-p 3000:3000 \
  	--env-file=.env \
  	kushal/admin

.PHONY: all
