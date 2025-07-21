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

RUN go build -o /app/llm-service ./services/llm-service/cmd

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/llm-service .

RUN chmod +x ./llm-service

EXPOSE 50051

CMD ["./llm-service"]