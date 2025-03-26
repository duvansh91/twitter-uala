FROM golang:1.18 as go-builder

WORKDIR /twitter-uala

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY cmd cmd
COPY pkg pkg
COPY config config

RUN go build -ldflags="-s -w" -o /main cmd/main.go

EXPOSE 8080

CMD ["/main"]
