![Build](https://github.com/DennisPing/Twitch-Chat-Downloader/actions/workflows/go.yml/badge.svg)
![Coverage](https://img.shields.io/badge/Coverage-22.6%25-red)

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