# LMS System – Main Service

Основной микросервис системы **LMS System** — pet-проект, разработанный в рамках стажировки для управления курсами, главами и уроками студентов **BITLAB Academy**.

## 🧩 Архитектура проекта

Проект задуман как система из нескольких микросервисов:

- 🧑‍🎓 **User Service (KeyCloak)** — управление пользователями и авторизацией
- 📘 **Main Service (Golang)** — управление курсами, доступами, группами *(реализуется сейчас)*
- 📁 **File Service (S3, MinIO)** — хранение файлов (загрузка, скачивание)

---

## ✅ Реализовано в `main_service`:

- [x] CRUD курсов
- [x] CRUD глав
- [x] CRUD уроков
- [x] Покупка курса пользователем
- [x] Авторизация / доступ по ролям
- [x] REST API (через Gin)
- [x] Документация через Swagger
- [x] Миграции базы данных (Goose)
- [x] Docker + Docker Compose
- [x] Unit тесты (Testify)
- [x] Моки репозиториев (Mockery)

---

## ⚙️ Технологии

- Язык: **Golang**
- Веб-фреймворк: **Gin**
- ORM: **GORM**
- БД: **PostgreSQL**
- Документация: **Swagger**
- Миграции: **Goose**
- Тестирование: **Testify**, **Mockery**
- Контейнеризация: **Docker**, **Docker Compose**
- Переменные окружения: `.env`

---

## 🚀 Запуск проекта

> Убедитесь, что установлен Docker и Docker Compose.

1. Скопируйте `.env.example`:
```bash
cp .env.example .env
```

2. Запустите проект:
```bash
docker-compose up --build
```

3. Swagger документация будет доступна по адресу:
```
http://localhost:8080/swagger/index.html
```

---

## 🧪 Тесты

```bash
go test ./...
```

---

## 📦 Миграции базы данных

```bash
go run cmd/migrate/main.go up
```

> Или используйте `make migrate-up`, если есть Makefile.

---

## 📝 Примечание

Это pet-проект, разрабатываемый в рамках стажировки, и не предназначен для использования в продакшене без доработок.

---

## 📫 Обратная связь

Разработчик: [Alexandra Flicker](https://github.com/Alexandra-Flicker)

---

