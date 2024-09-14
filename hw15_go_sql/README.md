## Домашнее задание №15 «Go-клиент для БД онлайн магазина»

- Создайте пакет для работы в БД
- Реализуйте объект подключния 
- Реализуйте функции для выполнения запросов из предыдущего ДЗ
- Обеспечьте атомарность выполнения запросов с помозью транзакций
- В качестве транспортного слоя используйте серверную часть из ДЗ №13

### Критерии оценки
- Понятность и чистота кода - до 2 баллов
- Реализован слой взаимодействия с БД - 2 балла
- Используются транзакции - 2 балла
- Сервис отдает данные по HTTP - 2 балла

#### Зачёт от 6 баллов

# Go Web Application

Этот проект представляет собой веб-приложение на Go с использованием PostgreSQL в качестве базы данных. Приложение предоставляет API для управления пользователями, продуктами и заказами.

## Структура проекта

- `main.go`: Основной файл приложения, который инициализирует базу данных и запускает HTTP сервер.
- `config/config.go`: Файл для загрузки конфигурации из `config.yml`.
- `db/db.go`: Файл для взаимодействия с базой данных.
- `handlers/handlers.go`: Файл, содержащий обработчики HTTP запросов.

## Требования

- Go (1.18 и выше)
- PostgreSQL
- Docker (опционально, для запуска PostgreSQL в контейнере)
- Docker Compose (опционально, для удобства запуска контейнеров)

## Установка и запуск

### 1. Установка зависимостей

Сначала убедитесь, что все зависимости установлены:

```bash
go mod tidy
```
### 2. Конфигурация
Создайте файл конфигурации config.yml в корневом каталоге проекта с следующим содержимым:

```yaml
database:
  user: your_db_user
  password: your_db_password
  dbname: your_db_name
  host: localhost
  port: 5432
  sslmode: disable
```
### 3. Запуск базы данных
Если вы используете Docker, запустите PostgreSQL с помощью Docker Compose:

```bash
docker-compose up -d
```
Убедитесь, что база данных запущена и доступна. Если вы не используете Docker, установите PostgreSQL локально и создайте базу данных.

### 4. Инициализация базы данных
Запустите скрипт для создания таблиц и схем в базе данных. Например, вы можете использовать команду psql:

```sql
CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    password VARCHAR(100)
);

CREATE TABLE Products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    price DECIMAL
);

CREATE TABLE Orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id),
    order_date DATE,
    total_amount DECIMAL
);

CREATE TABLE OrderProducts (
    order_id INTEGER REFERENCES Orders(id),
    product_id INTEGER REFERENCES Products(id),
    quantity INTEGER,
    PRIMARY KEY (order_id, product_id)
);
```

### 5. Запуск приложения
Запустите приложение с помощью следующей команды:

```bash
go run main.go
```

### 6. Тестирование API
Используйте curl или Invoke-WebRequest для тестирования API. Примеры команд:

Создание пользователя:

``` bash
curl -X POST http://localhost:8080/users -d '{
    "name": "John",
    "email": "unique_email@example.com",
    "password": "password123"
}' -H "Content-Type: application/json"
```
Получение пользователей:

``` bash
curl http://localhost:8080/users
```
Создание продукта:

```bash
curl -X POST http://localhost:8080/products -d '{
    "name": "Product1",
    "price": 10.99
}' -H "Content-Type: application/json"
```
Получение продуктов:

``` bash
curl http://localhost:8080/products
```

Создание заказа:

```bash
curl -X POST http://localhost:8080/orders -d '{
    "userId": 1,
    "orderDate": "2024-09-14",
    "totalAmount": 20.99,
    "orderProducts": [
        {
            "productId": 1,
            "quantity": 2
        }
    ]
}' -H "Content-Type: application/json"
```
Получение заказов:

```bash
curl http://localhost:8080/orders?user_id=1
```

### Устранение проблем
Ошибка: pq: повторяющееся значение ключа нарушает ограничение уникальности "users_email_key"

Эта ошибка возникает, когда вы пытаетесь вставить запись с уже существующим значением в поле, для которого установлено уникальное ограничение. Проверьте существующие записи в базе данных и используйте уникальные значения в запросах.

Ошибка: Error connecting to the database: pq: password authentication failed for user "your_db_user"

Проверьте правильность указанных учетных данных в файле config.yml и убедитесь, что пользователь и пароль корректны.