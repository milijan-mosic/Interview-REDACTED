# Interview-REDACTED

## Run the app

```sh

go run ./cmd/server/main.go
```

---

## Notes

### No library

Pros: Lightweight, no dependencies
Cons: Harder to extend features and/or slower to write features

### SQL only (no ORM)

Pros: Lightweight, no dependencies
Cons: Manual scan required for every struct

### No filtering / sorting in GET /v1/parts

Pros: Keeps the API minimal
Cons: Clients cannot query by status, supplier, or updated date

---

## Next steps

- Start using Chi/Gin/Echo
- Switch to ORM lib (GORM)
- Use framework's middlewares
- Add filters & sorting for GET /v1/parts
- Implement better auth
- Start using better database (Postgres)
