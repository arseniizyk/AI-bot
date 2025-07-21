FROM golang:alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.work go.work.sum ./

COPY proto/go.mod proto/go.sum ./proto/
COPY services/chat-bot/go.mod services/chat-bot/go.sum ./services/chat-bot/
COPY services/llm-service/go.mod services/llm-service/go.sum ./services/llm-service/

RUN go mod download

COPY proto ./proto
COPY services/chat-bot ./services/chat-bot
COPY services/llm-service ./services/llm-service

RUN go build -o /app/chat-bot ./services/chat-bot/cmd

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/chat-bot .

RUN chmod +x ./chat-bot

EXPOSE 8081

CMD ["./chat-bot"]