# Pub/sub message broker
## About
A pub/sub message broker server that communicates via QUIC.

## 3rd party packages:
- [quic-go](https://pkg.go.dev/github.com/lucas-clemente/quic-go) - a QUIC implementation in pure Go
- [testify](https://pkg.go.dev/github.com/stretchr/testify) - a testing framework for Go.

## Server specifications:
- Accepts QUIC connections on 2 ports - publisher port and subscriber port.
- The server notifies publishers when a subscriber has connected.
- If no subscribers are connected, the server informs the publishers.
- The server sends any messages received from publishers to all connected subscribers.

## Usage
### Build image:
```
docker build -t broker .
```

### Run image:
```
docker run -p 6666:6666/udp -p 6667:6667/udp broker
```

Adjust the ports according to the needs.

### Other information
- The TLS generating function was copied from the examples in the quic-go library, which are located [here](https://github.com/quic-go/quic-go/blob/d3c5f389d44797108a1bee7e06d5b92434c26d6d/example/echo/echo.go#L99C39-L99C39).

### TODOs:
- The project's folder structure could be better aligned with Go best practices.
- More unit tests need to be added.
- More consistent error handling is needed, as not all cases might be covered and proper actions taken.
- Graceful shutdown could be added.
- There are some hardcoded values that are left as-is for testing purposes.
- Remove subscribers/publishers from the list and close the stream on timeout/error.

## License
[MIT](https://choosealicense.com/licenses/mit/)