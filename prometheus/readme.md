## Требования

- Linux (RHEL / CentOS / Rocky / Alma)
- Docker
- Docker Compose
- Go

## Установка и запуск

### Установить Go
sudo dnf install golang -y

### Запустить Prometheus
docker compose up -d

### Запустить updater
cd updater
go run main.go

## Формат входного файла

hosts.xml:
<hosts>
  <host>
    <address>192.168.1.10</address>
  </host>
  <host>
    <address>192.168.1.11</address>
  </host>
</hosts>

## Как это работает

1. Обновляется `hosts.xml`
2. Go-скрипт пересобирает `targets.yml`
3. Prometheus автоматически применяет изменения
4. Перезапуск Prometheus не требуется

______________________________________________________
hosts.xml  →  Go-скрипт  →  targets.yml  →  Prometheus
______________________________________________________