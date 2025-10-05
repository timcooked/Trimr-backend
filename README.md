# Trimr Backend - URL Shortener API

RESTful API backend for Trimr URL shortener built with Go and MongoDB.

## Tech Stack

- **Language:** Go 1.21+
- **Database:** MongoDB
- **Architecture:** RESTful API
- **CORS:** Enabled for frontend integration
- **Package Management:** Go Modules

**Frontend Repository:** [Link to frontend repo](https://github.com/timcooked/Trimr)

## Features

- Generate short codes for long URLs
- Store and retrieve URL mappings
- Handle redirects to original URLs
- CORS support for web frontend
- Fast and lightweight

## Getting Started

### Prerequisites

- Go 1.21 or higher
- MongoDB (local or cloud instance)
- Git

### Installation

1. Clone the repository:
```bash
git clone https://github.com/timcooked/trimr-backend.git
cd trimr-backend
```

2. Initialize Go modules:
```bash
go mod init trimr-backend
go mod tidy
```

3. Install dependencies:
```bash
go get go.mongodb.org/mongo-driver/mongo
go get github.com/rs/cors
```

4. Set up environment variables (create `.env` file):
```env
MONGODB_URI=mongodb://localhost:27017
DATABASE_NAME=trimr
COLLECTION_NAME=urls
PORT=8080
```

5. Run the server:
```bash
go run main.go
```

Server will start on `http://localhost:8080`

## API Endpoints

### POST /shorten
Create a shortened URL

**Request:**
```json
{
  "url": "https://example.com/very-long-url"
}
```

**Response:**
```json
{
  "shortCode": "abc123",
  "shortUrl": "http://localhost:8080/redirect/abc123"
}
```

### GET /url/:shortCode
Get URL details by short code

**Response:**
```json
{
  "shortCode": "abc123",
  "originalUrl": "https://example.com/very-long-url",
  "shortUrl": "http://localhost:8080/redirect/abc123"
}
```

### GET /redirect/:shortCode
Redirect to the original URL

**Response:** HTTP 302 redirect to original URL

## Project Structure

```
trimr-backend/
├── main.go              # Server entry point
├── handlers/            # HTTP handlers
│   ├── shorten.go      # URL shortening logic
│   ├── redirect.go     # Redirect handler
│   └── url.go          # URL details handler
├── models/             # Database models
│   └── url.go          # URL model and database operations
├── go.mod              # Go modules
├── go.sum              # Dependencies
└── README.md
```

## Database Schema

**Collection: urls**
```json
{
  "_id": "ObjectId",
  "shortCode": "string",
  "originalURL": "string",
  "createdAt": "datetime"
}
```

## Usage with Frontend

1. Start MongoDB
2. Run the backend server (`go run main.go`)
3. Ensure server is running on port 8080
4. Start the frontend application
5. Frontend will connect to `http://localhost:8080`

## CORS Configuration

CORS is enabled for:
- Origin: `http://localhost:3000` (frontend)
- Methods: GET, POST, OPTIONS
- Headers: Content-Type, Authorization

## Testing

Test the API using curl:

```bash
# Shorten a URL
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://www.youtube.com"}'

# Get URL details
curl http://localhost:8080/url/abc123

# Test redirect
curl -I http://localhost:8080/redirect/abc123
```

## Dependencies

- `go.mongodb.org/mongo-driver/mongo` - MongoDB driver
- `github.com/rs/cors` - CORS middleware

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| MONGODB_URI | MongoDB connection string | mongodb://localhost:27017 |
| DATABASE_NAME | Database name | urls |
| COLLECTION_NAME | Collection name | urls |
| PORT | Server port | 8080 |

## Author

[timcooked](https://github.com/timcooked)
