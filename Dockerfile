# ตัวอย่าง Dockerfile สำหรับ Go
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o moon-telegrambot ./cmd/main.go

FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=builder /app/moon-telegrambot /moon-telegrambot
COPY .env /app/.env        # ถ้าใช้ .env จริง ต้อง mount หรือ ENV ผ่าน Cloud Run settings
CMD ["/moon-telegrambot"]
