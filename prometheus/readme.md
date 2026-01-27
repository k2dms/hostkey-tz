sudo dnf install golang -y


docker compose up -d


cd updater


go run main.go


Скрипт:
читает список хостов из XML-файла
извлекает IP-адреса из тега <address>
генерирует файл targets.yml в формате file_sd_configs
обеспечивает автоматическое добавление и удаление нод в Prometheus без 

______________________________________________________
hosts.xml  →  Go-скрипт  →  targets.yml  →  Prometheus
______________________________________________________