FROM golang:1.24.4-alpine3.22 AS builder

# Install build dependencies libraries
RUN apk add --no-cache build-base pkgconfig

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main main.go

# Run stage
FROM golang:1.24.4-alpine3.22

WORKDIR /app

# Copy required files
COPY --from=builder /app/main .
COPY ./start.sh /app/start.sh

ENV GIN_MODE=release
ENV TZ=Asia/Bangkok
EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]