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
## Analytics Dashboard

The URL Shortener publishes events to Kafka whenever a URL is created or visited. These events will later be consumed by a dedicated Analytics Service to generate click statistics and dashboards.

Event Flow
URL Shortener
      |
      ▼
Kafka Topic (url-events)
      |
      ▼
Analytics Service
      |
      ▼
PostgreSQL
      |
      ▼
Analytics Dashboard
Kafka Setup
Create Kafka Topic

Access the Kafka container:

docker exec -it kafka bash

Create the topic:

/opt/kafka/bin/kafka-topics.sh \
--create \
--topic url-events \
--bootstrap-server localhost:9092

Expected output:

Created topic url-events.
List Topics
/opt/kafka/bin/kafka-topics.sh \
--list \
--bootstrap-server localhost:9092

Expected output:

url-events
Listen to Events

Start a console consumer:

/opt/kafka/bin/kafka-console-consumer.sh \
--topic url-events \
--from-beginning \
--bootstrap-server localhost:9092
Sample Events
URL Created Event
{
  "event_type": "URL_CREATED",
  "url_id": 100001,
  "short_code": "Q0v",
  "long_url": "https://onlineksrtcswift.com/",
  "created_at": "2026-06-13T06:59:21.067133Z"
}
URL Visited Event
{
  "event_type": "URL_VISITED",
  "url_id": 100001,
  "short_code": "Q0v",
  "visited_at": "2026-06-13T07:04:57.303508Z"
}
Published Events
Event Type	Description
URL_CREATED	Published whenever a new short URL is created
URL_VISITED	Published whenever a short URL is accessed
Future Enhancements

The Analytics Service will consume events from the url-events topic and provide:

Total Click Count
URL-wise Analytics
Top Performing URLs
Daily Click Statistics
Real-Time Dashboard
Historical Trend Analysis
Geographic Analytics
Device & Browser Analytics

This event-driven architecture ensures that analytics processing does not impact URL redirection latency and can scale independently from the URL Shortener service.


---

## License

This project is intended for learning, and demonstration purposes.

