FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main cmd/main/main.go

FROM alpine:latest AS runtime

WORKDIR /app

RUN apk add --no-cache curl

COPY --from=builder /app/main .

CMD [ "/app/main" ]