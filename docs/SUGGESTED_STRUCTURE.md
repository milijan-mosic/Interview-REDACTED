# Suggested project structure (optional)

You may organize your code however you like. One common layout:

- cmd/server/main.go          # main package
- internal/httpapi/           # handlers, middleware
- internal/store/             # SQLite persistence
- internal/async/             # (optional) event processing
- internal/ratelimit/         # (optional) rate limiter
- schema.sql / seed.sql       # used by /debug/reset
