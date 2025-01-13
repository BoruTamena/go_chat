
# syntax=docker/dockerfile:1 

# stage -1 build stage 
FROM golang:1.23.2 AS buildstage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# copy the entire source tree
COPY  . .

# Generate swagger docs
RUN make swagger
 
RUN ls -R /app



# build the app 
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/app_bin ./cmd

# optionally list files in /build for debugging
RUN ls -R /build


# stage -2 running stage
FROM alpine:latest AS runstage

WORKDIR /
RUN apk --no-cache add ca-certificates

# copy binary file from build stage
COPY --from=buildstage /build/app_bin .
COPY --from=buildstage app/config/config.yaml  config/config.yaml
COPY --from=buildstage /app/docs  /docs

EXPOSE 8000

# Run the binary file as entrypoint
ENTRYPOINT ["./app_bin"]
