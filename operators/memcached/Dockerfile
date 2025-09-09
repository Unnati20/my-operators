FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o my-operator .

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/my-operator .
CMD ["./my-operator"]