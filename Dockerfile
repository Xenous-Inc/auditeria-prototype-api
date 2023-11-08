FROM golang:1.21.0-alpine as builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

RUN go build -o ./bin/app cmd/api/main.go

FROM alpine as runner

COPY --from=builder /usr/local/src/bin/app /

CMD [ "/app" ]