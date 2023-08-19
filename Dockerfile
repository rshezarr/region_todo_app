FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download && go build -o todo ./cmd/app/main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/todo .
COPY --from=builder /app/configs ./configs
EXPOSE 9090
CMD ["/app/todo"]