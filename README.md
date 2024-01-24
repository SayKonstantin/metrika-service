[![CI/CD](https://github.com/SayKonstantin/metrika-service/actions/workflows/ci_cd.yaml/badge.svg)](https://github.com/SayKonstantin/metrika-service/actions/workflows/ci_cd.yaml)

### Metrika Service

gRPC cервер для сбора данных от API Yandex Metrika и записи в BiqQuery через Cloud Storage

#### Методы


* `PushHitsToBQ` – загрузить просмотры

* `PushVisitsToBQ` – загрузить визиты

* `GetCounters` – получить список счетчиков по кабинету



## Алгоритм работы

* Получение данных по API по одному дню
* Создание `tsv` файлов
* Проверка наличия / создание таблицы в BigQuery
* Сохранение файла в CloudStorage Bucket (Bucket создается сам)
* Запись данных в таблицу BigQuery из CloudStorage

### Примечания

* BigQuery Dataset должен быть уже создан
* BigQuery Table создается автоматически если отсутствует
* `tsv` файлы удаляются автоматически


### Необходимые переменные окружения

Для использования переменных окружения используйте флаг  `--env`

| Переменная         | Описание                                         |
|--------------------|--------------------------------------------------|
| `GRPC_IP`          | Host                                             |
| `GRPC_PORT`        | Порт, который будет прослушивать сервис          | 
| `TG_TOKEN`         | Токен для telegram бота                          |
| `TG_CHAT`          | ID чата в который будут отправляться уведомления |
| `TG_ENABLED`       | Статус уведомлений                               |
| `KEYS_DIR `        | Путь к папке с сервисными ключами                |
| `PROMETHEUS_ADDR`  | Адрес сервера Prometheus                         |
| `ATTACHMENTS_DIR`  | Директория для `tsv` файлов                      |


