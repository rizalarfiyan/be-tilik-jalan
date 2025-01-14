FROM golang:1.23.2-alpine AS builder
WORKDIR /app
COPY . .

RUN apk add --no-cache upx

RUN mkdir -p /app/database/migrations
RUN go mod download
RUN go mod verify
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build \
    -o backend -a -ldflags '-s -w' -installsuffix cgo

RUN upx -9 backend

FROM scratch
WORKDIR /app

COPY --from=builder /app/backend /app/
COPY --from=builder /app/database/migrations /app/database/migrations

CMD ["/app/backend"]
