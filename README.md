# Twitch-Chat-Downloader
A multithreaded Twitch chat downloader written in Go

## Testing

Standard mode
```
go test
```

Verbose mode
```
go test -v
```

Test coverage
```
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```