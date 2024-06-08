FROM golang

WORKDIR /back

COPY . ./

RUN go mod download

ENV DB_HOST="DB"
ENV DB_PORT="27017"
ENV DB_NAME="DB"

EXPOSE 8000

CMD ["go", "run", "Main.go"]