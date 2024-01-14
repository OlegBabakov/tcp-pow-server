# POW-protected TCP server

## Small improvements

- "workers" amount limiter for incoming connections [internal/server/server.go](internal/server/server.go#53)
- "throttling" mechanism when the server closes connections if there are no available "workers" for certain time to handle a connection [internal/server/server.go](internal/server/server.go#99)

## Getting started

Requirements:

- Go 1.19+ installed (to run tests, start server or client without Docker)
- Docker
- Docker-Compose
- ENV file `.env` (rename [env.example](env.example) as example)

```
# Run server and client by docker-compose
make run-compose

# Run only server
make run-server

# Run only client
make run-client

# other command - call help
make help
```