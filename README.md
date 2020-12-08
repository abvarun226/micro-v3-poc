# comments service

## HTTP Setup
```
# Run micro server
$ MICRO_PROFILE=localnoauth MICRO_API_HANDLER=http ./micro server

# Build comment service
$ go build -tags http

# Run comment service
$ MICRO_REGISTRY_ADDRESS="127.0.0.1:8000" ./comments

# Test endpoint
$ http 'http://localhost:8080/comment/v1/hello'
```

## GRPC setup
```
# Setup proto
$ make init proto

# Run micro server
$ MICRO_PROFILE=localnoauth ./micro server

# Build comment service
$ go build -tags grpc

# Run comment service
$ MICRO_REGISTRY_ADDRESS="127.0.0.1:8000" ./comments

# Test endpoint
$ http 'http://localhost:8080/comment/list'
```