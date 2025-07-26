# AI-bot

**AI-bot** is a microservice-based application designed to interact with large language models (LLMs) via Telegram.

It consists of two gRPC-connected services:

- **chat-bot** – starts Telegram bot
- **llm-service** – communicates with the OpenRouter API

## 🚀 Usage

### 1. Clone the repository

```bash
git clone https://github.com/your-username/AI-bot.git  
cd AI-bot
```

### 2. Configure environment variables

Create `.env` files inside both `chat-bot/` and `llm-service/` directories based on the provided `.env.example` files.

### 3. Run the services with Docker Compose

```bash
docker compose up --build
```

## 📡 Protocol Buffers

All proto definitions are located in the `proto/` directory:

- `llm.proto` – gRPC service definition  
- `llm.pb.go`, `llm_grpc.pb.go` – auto-generated files

To regenerate:

```bash
protoc --go_out=. --go_opt=paths=source_relative \  
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \  
       llm.proto
```

## 🧠 Technologies Used

- Go 1.24.4  
- gRPC  
- Docker + Docker Compose  
- Telegram Bot API  
- OpenRouter API  
- Redis

---

## ⚙️ Redis Configuration

To provide Redis password, create a `redis.conf` file in the project root.  
Use `redis.conf.example` as a reference.  
The `docker-compose.yml` mounts this configuration automatically if present.

---

## 🔨 TODO:
Interaction with images, processing and transferring them to the LLM

## 📬 License

This project is licensed under the MIT License.
