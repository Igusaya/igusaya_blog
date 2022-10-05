# デプロイ用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.19.1-alpine3.16 as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main

# -------------------------------------------

# デプロイ用のコンテナ
FROM alpine:3.15 as deploy

WORKDIR /app
COPY --from=deploy-builder /app/main .

EXPOSE 8080
CMD ["/app/main"]

# -------------------------------------------

# ローカル開発環境で利用するホットリロード環境
FROM golang:1.19 as dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]