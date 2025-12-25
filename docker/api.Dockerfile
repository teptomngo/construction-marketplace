FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the binary outside the volume-mounted path
RUN go build -o /usr/local/bin/conmesh-api ./cmd/api

EXPOSE 8080

CMD ["conmesh-api"]
