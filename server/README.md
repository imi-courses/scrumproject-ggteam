# Настройка окружения

Создайте два файла: `.env` и `postgres.env`

Добавьте в них то, что находится в `example.env` (какие строчки кода в какой файл написано в комментариях файла `example.env`)

Далее нужно переименовать файл `config/example.config.yaml` в `config/prod.yaml` (`config/` - это директория)

Чтобы убедиться в том, что всё работает, нужно сначала запустить через `docker` все наши сервисы (база данных и сам сервер), использовав команду `docker-compose up --build -d`

url будет `http://localhost:5000`
API url `http://localhost:5000/api/v1`