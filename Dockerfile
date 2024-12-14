FROM golang:1.23.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY --chown=app:app . /app

RUN go build -o /serve

EXPOSE 3000

CMD ["/serve"]
