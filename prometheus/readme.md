sudo dnf install golang -y
docker compose up -d
cd updater
go run main.go
