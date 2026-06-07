# URL Shortener

A production-oriented URL Shortener service built in Go, following clean architecture principles, dependency injection, interface-driven design, and industry-standard project structure.

## Features

* Create short URLs from long URLs
* Resolve short URLs and redirect to original URLs
* In-memory repository implementation
* Base62 short code generation
* Thread-safe ID generation using atomic counters
* Clean separation of concerns
* Swagger/OpenAPI documentation
* Dependency injection
* Interface-driven design for testability

---

## Architecture

```text
HTTP Request
      |
      v
+-------------+
|   Handler   |
+-------------+
      |
      v
+-------------+
|   Service   |
+-------------+
      |
      +-------------------+
      |                   |
      v                   v
+-------------+   +---------------+
| Repository  |   | Generators    |
+-------------+   +---------------+
```

### Layers

#### Handler Layer

Responsible for:

* Request parsing
* Response serialization
* HTTP status codes
* Routing

#### Service Layer

Contains all business logic:

* URL validation
* Duplicate URL detection
* Short code generation
* URL resolution

#### Repository Layer

Abstracts data storage.

Current implementation:

* In-memory repository

Future implementations:

* MySQL
* PostgreSQL
* MongoDB
* Redis

#### Generator Layer

Responsible for:

* Unique ID generation
* Base62 encoding

---

## Project Structure

```text
url-shortener/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── docs/
│
├── internal/
│   │
│   ├── domain/
│   │   └── url.go
│   │
│   ├── dto/
│   │   ├── shorten_request.go
│   │   ├── shorten_response.go
│   │   └── error_response.go
│   │
│   ├── errors/
│   │   └── errors.go
│   │
│   ├── generator/
│   │   ├── id_generator.go
│   │   ├── atomic_id_generator.go
│   │   ├── code_generator.go
│   │   └── base62_generator.go
│   │
│   ├── handler/
│   │   └── url_handler.go
│   │
│   ├── repository/
│   │   ├── url_repository.go
│   │   └── inmemory_url_repository.go
│   │
│   └── service/
│       └── url_service.go
│
├── go.mod
├── go.sum
└── README.md
```

---

## API Endpoints

### Create Short URL

```http
POST /shorten
```

Request:

```json
{
  "url": "https://www.google.com"
}
```

Response:

```json
{
  "short_code": "1"
}
```

---

### Redirect

```http
GET /{shortCode}
```

Example:

```http
GET /1
```

Response:

```http
302 Found
Location: https://www.google.com
```

---

## Running the Application

### Clone Repository

```bash
git clone <repository-url>
cd url-shortener
```

### Install Dependencies

```bash
go mod tidy
```

### Start Server

```bash
go run cmd/server/main.go
```

Server starts on:

```text
http://localhost:8080
```

---

## Swagger Documentation

Generate Swagger documentation:

```bash
swag init -g cmd/server/main.go
```

Run the application:

```bash
go run cmd/server/main.go
```

Open:

```text
http://localhost:8080/swagger/index.html
```

---

## Example Usage

Create a short URL:

```bash
curl -X POST http://localhost:8080/shorten \
-H "Content-Type: application/json" \
-d '{"url":"https://golang.org"}'
```

Response:

```json
{
  "short_code": "1"
}
```

Resolve URL:

```bash
curl -v http://localhost:8080/1
```

---

## Design Decisions

### Why Interfaces?

Interfaces enable:

* Loose coupling
* Dependency injection
* Easier testing
* Future extensibility

### Why Base62?

Base62 uses:

```text
0-9
A-Z
a-z
```

This generates compact, URL-safe identifiers.

Example:

```text
12345 -> 3D7
```

### Why Atomic ID Generator?

Atomic operations provide:

* Thread safety
* Better performance than mutexes
* Simpler implementation

### Why Context Propagation?

All repository and service methods accept `context.Context` to support:

* Request cancellation
* Timeouts
* Tracing
* Observability

---

## Future Enhancements

* MySQL/PostgreSQL persistence
* Redis caching
* URL expiration
* Custom aliases
* Click analytics
* Rate limiting
* Distributed ID generation (Snowflake)
* Prometheus metrics
* Structured logging
* Graceful shutdown
* Unit and integration tests
* Docker support
* Kubernetes deployment

---

## Tech Stack

* Go
* net/http
* Swagger/OpenAPI
* Clean Architecture Principles
* Dependency Injection
* Interface-Based Design

---

## License

This project is intended for learning, and demonstration purposes.
