# eth-balance-proxy

eth-balance-proxy is a server that proxies eth_getbalance calls to external provider

## How to Run the Go App

```sh
go build -o eth-balance-proxy
./eth-balance-proxy
```

Or, to run directly:

```sh
go run main.go
```

## How to Launch Stress Tests

Make sure the Go app is running on `localhost:8080` before starting the stress test.

```sh
./stress.sh
```

## Improvements

- Use LRU cache with expiration of 10 seconds for balance responses.
- Add algorithm to automatically shut down providers that are failing.
- Add Docker setup with Grafana for monitoring and visualization.
- Implement a complex strategy for provider selection
