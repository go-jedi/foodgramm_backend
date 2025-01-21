FROM golang:1.23.3-alpine AS builder

WORKDIR /github.com/go-jedi/foodgrammm-backend/app
COPY . /github.com/go-jedi/foodgrammm-backend/app

RUN go mod download
RUN go build -o .bin/app cmd/app/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/go-jedi/foodgrammm-backend/app/.bin/app .
COPY config.yaml /root/

CMD ["./app", "--config", "config.yaml"]