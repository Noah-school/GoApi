# Example Go API

A tiny example HTTP API written in Go. It demonstrates:

- Authentication middleware (token via `Authorization` or `token` header)
- Routing with `chi`
- Consistent JSON error handling
- Mocked data store for users and coin balances

## Prerequisites

- Go installed (any recent 1.x version)

## Run

```bash
go mod tidy
go run cmd/api/main.go
```

Server starts on `http://localhost:8000`

## API

- Base URL: `http://localhost:8000`

### GET /account/coins

Returns the coin balance for a user.

- Query:
	- `username`
- Headers:
	- `Authorization: <token>`

#### Example

```bash
curl -s "http://localhost:8000/account/coins?username=noah" \
	-H "Authorization: Bearer root"
```

Response

```json
{
	"code": 200,
	"balance": 100
}
```

#### Errors

Missing/invalid credentials:

```json
{
	"code": 400,
	"message": "invalid username or token"
}
```

## Test Users (mock DB)

- `noah` / `root` outputs the balance `100`
- `sander` / `crazylongpasswordthateveniwouldforgetit` outputs the balance `727`

## Acknowledgements

This was made possible by [@mr_mux408](https://www.youtube.com/@mr_mux408) with his Golang introduction
[video](https://www.youtube.com/watch?v=8uiZC0l4Ajw&t=2406s).