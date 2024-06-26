FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY . ./

RUN go get -d -v
RUN go build -o /go/bin/app

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /go/bin/app /app

ENTRYPOINT ["/app/app"]

EXPOSE 3001
