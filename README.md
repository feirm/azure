<h1 align="center">Azure</h1>
<p align="center">💰 Cryptocurrency data microservice for the Feirm Platform.</p>

<p align="center">
    <img src="https://img.shields.io/github/go-mod/go-version/feirm/azure?style=for-the-badge" />
    <img src="https://goreportcard.com/badge/github.com/feirm/azure?style=for-the-badge" />
</p>

Azure is a microservice used to fetch network parameters for a specified cryptocurrency. This data is subsequently used in the Feirm App to apply the right parameters for cryptography.

## Getting Started
```bash
$ go run cmd/azure/main.go
```

## API
```
GET /v1/coins

POST /v1/coin
{
    "ticker": "xfe"
}
```