FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY . ./

RUN go get -d -v
RUN go build -o /go/bin/app

#ENV DB_HOST="localhost"
#ENV DB_PORT="27017"
#ENV DB_NAME="DB"

#EXPOSE 8000

#CMD ["go", "run", "Main.go"]
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /go/bin/app /app
COPY .env /app/.env

ENTRYPOINT ["/app/app"]

EXPOSE 8000