# Album Listing api

## Overview

This document outlines the design and usage of the Album Lisitng api in Go using the Gin framework. 

## Structs

### Album

- `ID` (string): The unique identifier for the album.
- `Title` (string): The title of the album.
- `Artist` (string): The artist of the album.
- `Price` (float64): The price of the album.

## File Structure

```plaintext
api/
├── main.go
├── controllers/
│   └── album_controller.go
├── routes/
│   └── album_routes.go
├── models/
│   └── album.go
├── data/
│   └── albums.go
├── docs/
│   └── documentation.md
├── go.mod
└── go.sum
```
## Usage
To run the application, use the following command:

```
go run main.go
```
Follow the interactive prompts to:

#### Add a New Album
Endpoint: POST /albums/create_album
Request Body:
```json
{
    "id": "string",
    "title": "string",
    "artist": "string",
    "price": float64
}
```

#### Get All Albums
Endpoint: GET /albums
Response:

```json
[
    {
        "id": "1",
        "title": "Blue Train",
        "artist": "John Coltrane",
        "price": 56.99
    },
    {
        "id": "2",
        "title": "Jeru",
        "artist": "Gerry Mulligan",
        "price": 17.99
    },
    {
        "id": "3",
        "title": "Sarah Vaughan and Clifford Brown",
        "artist": "Sarah Vaughan",
        "price": 39.99
    },
    {
        "id": "4",
        "title": "The Magic Flute",
        "artist": "Graham Norton",
        "price": 39.99
    },
    {
        "id": "10",
        "title": "Tamirat",
        "artist": "John Coltrane",
        "price": 56.99
    }
]

```
### Get Album by ID
Endpoint: GET /albums/{id}
Response:
```json

{
    "id": "1",
    "title": "Blue Train",
    "artist": "John Coltrane",
    "price": 56.99
}
```







