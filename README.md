![Build](https://github.com/DennisPing/Twitch-Chat-Downloader/actions/workflows/go.yml/badge.svg)

# Twitch-Chat-Downloader

A multithreaded Twitch chat downloader written in Go

Inspired by: https://github.com/PetterKraabol/Twitch-Chat-Downloader

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