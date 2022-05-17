FROM golang:1.18

WORKDIR /go/src

COPY go.mod .
COPY go.sum .

COPY . .

ENV CGO_ENABLED=0
RUN go build -o /go/src/app ./cmd/server

CMD ["/go/src/app"]