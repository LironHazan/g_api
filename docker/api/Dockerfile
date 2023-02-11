############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/gapi/api/
# Download Go modules
COPY . .
RUN go mod download

WORKDIR $GOPATH/gapi/api/apps/api

# Build the binary.
RUN go build -o /go/bin/api
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy executable.
COPY --from=builder /go/bin/api /go/bin/api

EXPOSE 3000
ENTRYPOINT ["/go/bin/api"]

#Build image:
# docker build -t go-api:latest -f ./docker/api/Dockerfile .
#Interactively run the container:
#```docker run -it -p 3000:3000 go-api:latest sh```
