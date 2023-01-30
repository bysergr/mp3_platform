#build stage
FROM golang:1.19.4-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY ["go.mod", "go.sum", "./"] 
RUN go mod download
COPY . .
RUN go build -o app -v .

#final stage
FROM alpine

RUN apk update
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app . 

ENV GIN_MODE=release
EXPOSE 5000

ENTRYPOINT ["./app"]
