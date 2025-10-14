# 📝 Todo API — Simple CRUD App in Go

A simple **Todo List REST API** built with **Go**, using **GORM** and **PostgreSQL**.  
This project demonstrates clean architecture principles (Repository & Service layers), dependency injection (DI), and idiomatic Go project structure.

---

## ⚙️ Tech Stack

| Component                 | Description                              |
| ------------------------- | ---------------------------------------- |
| **Language**              | Go 1.22+                                 |
| **Framework / Router**    | Gin                                      |
| **ORM**                   | [GORM](https://gorm.io)                  |
| **Database**              | PostgreSQL                               |
| **Dependency Management** | Go Modules                               |
| **Architecture**          | Layered (Handler → Service → Repository) |

---

## 📂 Project Structure

```
todo-api/
├── cmd/
│   └── main.go                  # Application entry point
│
├── internal/
│   ├── controllers/             # HTTP controllers (handle requests, call services)
│   │   └── todo_controller.go
│   │
│   ├── models/                  # Database models (e.g. Todo struct)
│   │   └── todo.go
│   │
│   ├── repositories/            # Database access layer (GORM operations)
│   │   ├── todo_repository.go   # Repository interface
│   │   └── todo_repository_impl.go
│   │
│   ├── services/                # Business logic layer
│   │   ├── todo_service.go      # Service interface
│   │   └── todo_service_impl.go
│   │
│   └──                         # (other internal packages if needed)
│
├── pkg/
│   └── config/                  # App configuration & DB connection
│       └── config.go
│
├── .env                         # Environment variables (local)
├── .env.example                 # Example environment config for setup
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

---

### 🧩 Layer Responsibilities

| Layer                      | Description                                                                           |
| -------------------------- | ------------------------------------------------------------------------------------- |
| **cmd/**                   | Entry point for the app (`main.go`) — initializes dependencies and starts the server. |
| **internal/controllers/**  | Handles incoming HTTP requests and responses.                                         |
| **internal/services/**     | Contains business logic — uses repositories to manipulate data.                       |
| **internal/repositories/** | Handles data persistence via GORM and PostgreSQL.                                     |
| **internal/models/**       | Defines database models / entities.                                                   |
| **pkg/config/**            | Loads environment variables, sets up DB connection.                                   |
| **.env**                   | Local environment configuration (ignored by Git).                                     |
| **.env.example**           | Template of required environment variables.                                           |

---

This structure follows the **Clean Architecture / Layered Design** principle, which provides:

- Clear separation of concerns between layers.
- Easy maintenance and scalability.
- Simple dependency injection (DB, logger, cache, etc.).
- Better testability with independent unit tests for each layer.

---

## 🚀 Getting Started

### 1️⃣ Clone the repository

```bash
git clone git@github.com:KaitoHasei/go-todo-api.git
cd go-todo-api
```

### 2️⃣ Initialize and download dependencies

```bash
go mod tidy
```

### 3️⃣ Set up environment variables

Create a `.env` file in the root directory:

```bash
APP_PORT=8080

DATABASE_URL=postgres://postgres:postgres@localhost:5432/todo?sslmode=disable
```

---

### 4️⃣ Run PostgreSQL locally (Docker)

If you have Docker installed, you can quickly start a Postgres container:

```bash
docker run --name todo-db -e POSTGRES_PASSWORD=your_password -p 5432:5432 -d postgres
```

---

### 5️⃣ Run the application

```bash
go run cmd/main.go
```

The API will be available at:
👉 `http://localhost:8080`

---

## 🧠 API Endpoints

| Method   | Endpoint     | Description     |
| -------- | ------------ | --------------- |
| `GET`    | `/todos`     | Get all todos   |
| `GET`    | `/todos/:id` | Get todo by ID  |
| `POST`   | `/todos`     | Create new todo |
| `PUT`    | `/todos/:id` | Update todo     |
| `DELETE` | `/todos/:id` | Delete todo     |

---

## 🧩 Example JSON

### ✅ Create Todo

**Request**

```json
POST /todos
{
  "title": "Learn Go",
  "completed": false
}
```

**Response**

```json
{
  "id": 1,
  "title": "Learn Go",
  "completed": false,
  "created_at": "2025-10-13T15:04:05Z"
}
```

---

## 🧪 Testing

Run tests (if you’ve written them):

```bash
go test ./...
```

---

## 🧰 Useful Commands

| Command          | Description                 |
| ---------------- | --------------------------- |
| `go mod tidy`    | Sync dependencies           |
| `go build ./cmd` | Build the binary            |
| `go install`     | Install executable globally |
| `go fmt ./...`   | Format all Go files         |
| `go vet ./...`   | Analyze code for issues     |

---

## 🧱 Architecture Overview

```text
[ Handler (HTTP) ]
       ↓
[ Service (Business Logic) ]
       ↓
[ Repository (Database Access) ]
       ↓
[ PostgreSQL ]
```

- **Handler**: Defines API endpoints (maps routes to service calls).
- **Service**: Contains core business logic (validation, transformation).
- **Repository**: Handles persistence using GORM.
- **Database**: PostgreSQL as storage layer.

---

## 🧩 Dependency Injection (DI) Example

```go
func main() {
    db := database.Connect()
    repo := repositories.NewTodoRepository(db)
    service := services.NewTodoService(repo)
    handler := handlers.NewTodoHandler(service)

    router := routes.SetupRouter(handler)
    router.Run(":8080")
}
```

---

## 📜 License

This project is licensed under the MIT License — feel free to use and modify.

---

## 🧑‍💻 Author

**Kaito Hasei**
🔗 [github.com/KaitoHasei](https://github.com/KaitoHasei)
