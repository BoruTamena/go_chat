# syntax=docker/dockerfile:1 

# stage -1 build stage 
FROM golang:1.23.2 AS buildStatge

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# copy src tree to wildcard  
COPY . ./

# build the app 
RUN  CGO_ENABLED=0 GOOS=linux go build -o /build ./cmd



# stage -2 running stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

# copy binary file from build stage
COPY --from=buildStatge /build /build

EXPOSE 8000

ENTRYPOINT ["build" ]