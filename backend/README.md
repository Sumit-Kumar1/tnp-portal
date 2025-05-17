# TnP Portal Backend

This is the backend service for the TnP (Training and Placement) Portal, built with Go and PostgreSQL.

## Features

- User authentication and authorization
- Student profile management
- Job portal functionality
- Company management
- Application management
- Interview scheduling
- Document management
- Notifications
- Analytics and reporting

## Prerequisites

- Go 1.21 or later
- PostgreSQL 12 or later
- Make (optional)

## Setup

1. Clone the repository:
```bash
git clone https://github.com/skumar9664akroger.com/tnp-portal.git
cd tnp-portal/backend
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file in the root directory with the following variables:
```env
# Server Configuration
PORT=8080
ENV=development

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=tnp_portal

# JWT Configuration
JWT_SECRET=your-secret-key
JWT_EXPIRATION=24h

# CORS Configuration
ALLOWED_ORIGINS=http://localhost:3000
```

4. Set up the database:
```bash
# Create database
createdb tnp_portal

# Run migrations
psql -d tnp_portal -f migrations/001_create_users_table.sql
```

## Running the Server

1. Start the server:
```bash
go run cmd/server/main.go
```

2. The server will start on `http://localhost:8080`

## API Documentation

The API documentation is available in the `static/openapi.json` file. You can view it using Swagger UI or any other OpenAPI viewer.

## Project Structure

```
backend/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handler/
│   │   └── user_handler.go
│   ├── middleware/
│   │   └── auth.go
│   ├── models/
│   │   └── user.go
│   ├── repository/
│   │   ├── user_repository.go
│   │   └── postgres/
│   │       ├── db.go
│   │       └── user_repository.go
│   └── service/
│       ├── user_service.go
│       ├── user_service_impl.go
│       ├── jwt_service.go
│       └── jwt_service_impl.go
├── migrations/
│   └── 001_create_users_table.sql
├── static/
│   └── openapi.json
├── .env
├── go.mod
└── README.md
```

## Testing

To run tests:
```bash
go test ./...
```

## Contributing

1. Fork the repository
2. Create a new branch
3. Make your changes
4. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
