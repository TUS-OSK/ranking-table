FROM golang:latest

RUN go install github.com/air-verse/air@latest
WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
