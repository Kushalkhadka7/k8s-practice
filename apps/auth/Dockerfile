# Stage: Development.
FROM golang:alpine as dev

WORKDIR /app 

RUN CGO_ENABLED=0 
RUN GOOS=linux 

# Enable go modules
ENV GO111MODULE=on

COPY . /app/

ENV SHELL=/bin/bash

EXPOSE 4000

CMD [ "go", "run", "main.go"]

# Stage: Builder.
FROM golang:alpine as builder

WORKDIR /app

COPY --from=dev /app /app

CMD [ "make", "build-app"]


# Stage: prod.
FROM golang:alpine as prod

WORKDIR /app

COPY --from=builder /app/auth /app/

EXPOSE 4000

CMD [ "./auth"]