# GO Clean Architecture

A structured and modular clean architecture for building web applications using the Gin framework in Go. This project follows clean code principles, ensuring maintainability, scalability, and ease of testing.

## 🚀 Features

- **Modular structure** with a clear separation of concerns
- **Authentication module** with user registration, login, refresh token, and profile (protected) functionality
- **Scalable and maintainable** architecture

## 🏗️ Project Structure

```sh
📁 go-clean-architecture
├── 📁 cmd  
│   └── 📁 api          # Entry point for the application (e.g., main.go inside api/)
├── 📁 config          # Configuration files (e.g., database, environment variables)
├── 📁 internal        # Core business logic, domain models, and use cases
│   ├── 📁 domain       # Entities and core business rules
│   ├── 📁 infrastructure # External dependencies (database, repositories)
│   │   └── 📁 entities  # ORM models or database-specific representations
│   ├── 📁 interfaces   # Application interfaces (HTTP handlers, routes)
│   │   ├── 📁 http      # Controllers and middleware
│   │   └── 📁 routes    # Route definitions
│   └── 📁 usecases     # Business logic (application services)
├── 📁 pkg             # Shared utility packages
│   └── 📁 security    # Authentication (JWT, encryption, hashing)
```

## 🛠️ Getting Started

### ✅ Prerequisites

- **Go 1.23+**

### 📥 Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/mehedicode-lab/go-clean-architecture.git
   cd go-clean-architecture
   ```

2. **Initialize the Go module:**
   ```sh
   go mod init github.com/mehedicode-lab/go-clean-architecture
   go mod tidy
   ```

3. **Configure the application:**
   - Copy `.env.sample` to `.env`
   - Update `.env` with your settings

4. **Run the application:**
   ```sh
   go run cmd/api/main.go
   ```

## 📌 API Endpoints

### 🔐 Authentication

#### 1️⃣ Register
- **Endpoint:** `POST /auth/register`
- **Request Body:**
  ```json
  {
    "fullName": "John Doe",
    "email": "example@example.com",
    "password": "password123"
  }
  ```

#### 2️⃣ Login
- **Endpoint:** `POST /auth/login`
- **Request Body:**
  ```json
  {
    "email": "example@example.com",
    "password": "password123"
  }
  ```

#### 3️⃣ Refresh Token
- **Endpoint:** `POST /auth/refresh`
- **Request Body:**
  ```json
  {
    "refresh_token": "***********"
  }
  ```

#### 4️⃣ Get Profile (Protected)
- **Endpoint:** `GET /auth/profile`
- **Request Header:**
  ```sh
  Authorization: Bearer ********
  ```

---

This architecture provides a solid foundation for building scalable and well-structured applications in Go using the Gin framework. 🚀

