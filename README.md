# 📝 To-Do List API in Go (Gin Framework)

A simple **CRUD REST API** written in **Go** using the **Gin framework**. This project demonstrates routing, JSON handling, request validation, and basic in-memory data storage.

---

## 🚀 Features
- CRUD operations for a simple To-Do list
- JSON responses for all endpoints
- Request validation (e.g., title required)
- In-memory storage (no database required)
- Simple, lightweight, suitable for demo or small client backends

---

## 📂 Project Structure
```
todo-api/
├── main.go      # Main Go program
├── go.mod       # Go modules file
└── README.md
```

---

## 🔧 Installation & Run

```bash
# Clone repo
git clone https://github.com/your-username/todo-api.git
cd todo-api

# Initialize Go modules
go mod tidy

# Run the API
go run main.go
```

The server runs on `http://localhost:8080`.

---

## 🛠 API Endpoints

| Method | Endpoint         | Description                     |
|--------|----------------|---------------------------------|
| GET    | /todos          | List all todos                  |
| GET    | /todos/:id      | Get a todo by ID                |
| POST   | /todos          | Create a new todo               |
| PUT    | /todos/:id      | Update an existing todo         |
| DELETE | /todos/:id      | Delete a todo                   |

### Example Requests

**Create a todo**
```bash
curl -X POST http://localhost:8080/todos \ 
-H "Content-Type: application/json" -d '{"title":"Learn Go"}'
```

**Update a todo**
```bash
curl -X PUT http://localhost:8080/todos/1 -H "Content-Type: application/json" -d '{"title":"Learn Go and Gin","done":true}'
```

**Delete a todo**
```bash
curl -X DELETE http://localhost:8080/todos/1
```

**Get all todos**
```bash
curl http://localhost:8080/todos
```

**Get a todo by ID**
```bash
curl http://localhost:8080/todos/1
```

---

## 🧑‍💻 Next Steps / Enhancements
- Add unit tests using `httptest`
- Persist data using a database (SQLite/Postgres)
- Add authentication (JWT/API key)
- Pagination and filtering for large lists
- Dockerize for easy deployment

---

## 📜 License
MIT License – free to use and modify.
