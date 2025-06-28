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
COPY ./config/credentials.yml.enc config/credentials.yml.enc
COPY ./start.sh /app/start.sh

COPY --from=builder /app/main .

ENV GIN_MODE=release
ENV TZ=Asia/Bangkok
EXPOSE 3000
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]