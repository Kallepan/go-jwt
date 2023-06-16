FROM golang:alpine AS builder

WORKDIR /project/go-jwt

COPY src/go.* ./

RUN go mod download

COPY src/ .
RUN go build -o /project/go-jwt/build/main .

FROM alpine:latest
COPY --from=builder /project/go-jwt/build/main /app/build/main

EXPOSE 8080
ENTRYPOINT [ "/app/build/main" ]