### Development

```
go run main.go
```

Runs the api in the development mode.
Request with POST method on http://localhost:3001/healthcheck with a file.

### Build Docker Image

```
docker build -f Dockerfile -t healthcheck-api .
```

### Run Docker Image

```
docker run --publish 3001:3001 --name healthcheck-api healthcheck-api
```
