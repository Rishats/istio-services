# Используйте образ Go 1.22.1 как базовый
FROM golang:1.22.0 as builder

# Установите рабочую директорию в контейнере
WORKDIR /app

# Копируйте go.mod и go.sum для управления зависимостями
COPY go.mod go.sum ./

RUN go mod download

# Копируйте исходный код в контейнер
COPY . ./

# Соберите приложение
RUN CGO_ENABLED=0 GOOS=linux go build -v -o service cmd/product/main.go

# Используйте образ alpine для окончательного образа
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируйте исполняемый файл из builder-стадии
COPY --from=builder /app/service .

# Укажите порт, на котором будет работать сервис
EXPOSE 8080

# Запустите приложение
CMD ["./service"]

