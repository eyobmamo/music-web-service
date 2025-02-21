# Music Web Service

This is a simple web service built with Go and the Gin framework. It allows you to manage a collection of music albums. The API exposes endpoints to list, add, and retrieve albums by ID.

## Features

- Get a list of all albums.
- Add a new album to the collection.
- Retrieve an album by its ID.

## Requirements

- Go 1.18 or higher.
- Gin Gonic package.

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/music-web-service.git
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Run the application:
    ```bash
    go run main.go
    ```

The server will run on `http://localhost:8080`.

## API Endpoints

### GET `/albums`

Fetches all albums in the collection.

**Response:**
- HTTP Status: 200 OK
- JSON Array of albums

Example:
```json
[
  {
    "id": "1",
    "title": "Blue Train",
    "artist": "John Coltrane",
    "price": 56.32
  },
  {
    "id": "2",
    "title": "Jeru",
    "artist": "Gerry Mulligan",
    "price": 17.99
  }
]

