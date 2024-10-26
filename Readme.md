#### Старт сервера

Cоздать конфиг файл формата

```yaml
telegram:
  bot_token: "BOT TOKEN"
  chat_id: "CHAT ID"
  user: "user"
  password: "test-connect"
  port: "8087"
  host: "localhost"

```
Запустить
```
go run cmd/app/main.go --config=config.yaml
```