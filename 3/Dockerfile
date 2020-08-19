FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o binary

ENTRYPOINT ["/app/binary"]