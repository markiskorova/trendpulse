
# 🧠 Trend Pulse – Analytics Simulation Backend (Auth MVP)

Trend Pulse is a simulated SaaS analytics platform designed to showcase secure backend architecture, authentication, and API design for immersive brand interaction tracking. This phase focuses on the MVP core — a fully functional user registration, login, and authenticated API using Go, PostgreSQL, Docker, and JWT.

---

## ✨ Features (Implemented in MVP)

### ✅ User Registration
- Users can create accounts via `/register`.
- Passwords are securely hashed with bcrypt before storing.
- Duplicate emails are rejected using database uniqueness constraint.

### ✅ User Login
- Users log in via `/login` with email and password.
- Upon success, a JWT is issued with encoded `user_id` and expiration.

### ✅ JWT-Based Auth Middleware
- Middleware parses and validates the `Authorization: Bearer <token>` header.
- Extracted user ID is added to the request context.
- Protected routes deny unauthenticated access.

### ✅ Protected API Endpoint
- `POST /api/events` is restricted to logged-in users.
- Returns a placeholder message for now: `"Event received (stub)"`
- Prepares the path for future expansion into event persistence and analytics.

### ✅ Health Check
- Simple `GET /health` endpoint returns `OK` for readiness checks.

---

## ⚙️ Tech Stack

- **Go** – HTTP server, routing, middleware, JWT
- **PostgreSQL** – Relational database for user storage
- **Docker + Compose** – Containerized environment for backend and DB
- **JWT** – Stateless user authentication
- **bcrypt** – Password hashing
- **PowerShell/cURL/Postman** – Local testing tools

---

## 🚀 How to Run It Locally

### 1. Clone the Repository
```bash
git clone https://github.com/yourusername/trendpulse.git
cd trendpulse
```

### 2. Build and Start Services
```bash
docker-compose down -v
docker-compose up --build
```

### 3. Health Check
```bash
curl http://localhost:8080/health  # should return OK
```

### 4. Register a New User
```powershell
Invoke-RestMethod -Method POST http://localhost:8080/register `
  -ContentType "application/json" `
  -Body '{"email":"test@example.com", "password":"secret123"}'
```

### 5. Log In and Save JWT Token
```powershell
$token = (Invoke-RestMethod -Method POST http://localhost:8080/login `
  -ContentType "application/json" `
  -Body '{"email":"test@example.com", "password":"secret123"}').token
```

### 6. Call Protected Endpoint
```powershell
Invoke-RestMethod -Method POST http://localhost:8080/api/events `
  -Headers @{ Authorization = "Bearer $token" }
```

---

## 📦 Folder Structure
```
trendpulse/
├── docker-compose.yml
├── docker-entrypoint-initdb.d/0001_create_users.sql
├── trendpulse-backend/
│   ├── cmd/api/main.go
│   ├── internal/
│   │   ├── db/db.go
│   │   ├── handlers/
│   │   │   ├── auth.go
│   │   │   ├── event.go
│   │   │   └── health.go
│   │   ├── middleware/jwt.go
│   │   └── models/user.go
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
```

---

## 🛣 Roadmap (Next Phases)

1. **Event Persistence**
   - Store impressions, clicks, purchases in PostgreSQL
2. **Background Processing**
   - Use NATS or RabbitMQ + worker service
3. **GraphQL API**
   - Query aggregated analytics via Strawberry or gqlgen
4. **React + TypeScript Dashboard**
   - View insights via charts and tables
5. **AWS Deployment**
   - S3 + EC2 or ECS for full-stack deployment

---

## 📄 License
MIT – Free to use and adapt. Mention Trend Pulse if shared.

---

## 🙋‍♂️ Author
Built by [Your Name] – Software Engineer focused on secure backend systems and platform-scale architecture.
