FROM golang:alpine AS builder

WORKDIR /project/go-jwt

COPY src/go.* ./

RUN go mod download

COPY src/ .
RUN go build -o /project/go-jwt/build/main .

FROM alpine:latest
COPY --from=builder /project/go-jwt/build/main /app/build/main

COPY src/migrations /app/migrations

RUN apk update && apk add postgresql-client
COPY launch.sh /app/launch.sh
RUN chmod +x /app/launch.sh

EXPOSE 8080
CMD [ "/app/launch.sh" ]