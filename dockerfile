FROM golang:1.23-alpine

WORKDIR /app

RUN apk add --no-cache gcc musl-dev git

ENV GOPROXY=proxy.golang.org

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o receipt_service

RUN chmod +x receipt_service

EXPOSE 8080

CMD ["./receipt_service"]
