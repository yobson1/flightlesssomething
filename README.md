# FlightlessSomething

[flightlessmango.com](https://flightlessmango.com/) website clone, written in Go.

Yes, there is a lot of crappy copypasta html/css/js code. As long as it works! 🤷

# Features

* Written in Go:
  * Fast performance
  * Multithreaded
  * Single, statically linked binary
* Uses `gin` web framework
* Uses `gorm` ORM (Can be easily ported to other databases)

## Features that will NOT be included

* TLS/SSL/ACME - use reverse proxy (I suggest [Caddy](https://github.com/caddyserver/caddy))

# Development

To run this code locally, install `go`, setup your .env file following the [.env.example](./.env.example) file and then run:

```bash
go run cmd/flightlesssomething/main.go
```

Then open in browser: http://localhost:8080/
