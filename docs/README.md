# Backend Take-Home: Parts Service (Go)

## Goal

Build a small HTTP service in Go that exposes a minimal Parts API. The focus is on sound backend engineering practices (correctness, clarity, and testability), not on building a large feature set.

**Timebox:** please keep this to **~2 hours**.

We will use your submission as the basis for a follow-up technical interview.

---

## What to build (required)

### 1) HTTP server (net/http)

Implement a JSON REST API using Go's `net/http` (a router library is fine, but the service must be an `http.Server`-based app).

#### Run contract (important)

Your service **must**:

- Start via: `go run ./cmd/server`
- Run tests via: `go test ./...`
- Listen on:
  - `PORT` env var if set (e.g., `PORT=18080`)
  - otherwise default to **8080**
- Use SQLite for persistence:
  - `DATABASE_PATH` env var if set
  - otherwise default to `./data.db`
- Use a static bearer token for auth:
  - `AUTH_TOKEN` env var if set
  - otherwise default to `dev-token`

> We will evaluate your service using an automated harness that relies on the run contract above.

### 2) Data model (SQLite)

Store parts in SQLite. You may use any SQLite driver.

You may use the provided `schema.sql` and `seed.sql`, or write your own schema, as long as you meet the API and behavioral requirements below.

**Statuses:** `Draft`, `In Review`, `Approved`, `Rejected`

### 3) Required endpoints

#### Health

- `GET /healthz`
  - **No auth**
  - Response: `200` with JSON like `{ "status": "ok" }`

#### Debug endpoints (exercise-only)

These endpoints are required to make evaluation deterministic. They are not “product” endpoints.

- `POST /debug/reset`
  - **No auth**
  - Resets the database to a deterministic state (schema + seed).
  - Response: `200` with JSON (any shape is fine).

- `GET /debug/stats`
  - **No auth**
  - Response: `200` with JSON containing at least:
    - `parts_count` (int)
    - `audit_count` (int)

#### Parts API (versioned)

All `/v1/*` endpoints **must** require auth: `Authorization: Bearer <token>`.

- `GET /v1/parts`
  - Returns a JSON array of parts (order is up to you).
  - Must return the seeded parts after `/debug/reset`.

- `GET /v1/parts/{id}`
  - Returns a single part.
  - If not found: `404`.

- `PATCH /v1/parts/{id}/status`
  - Body JSON: `{ "status": "Approved" }` (string; must be one of the allowed statuses)
  - Behavior:
    - Update the part's status.
    - Insert an audit row capturing old/new status.
  - If status is invalid: `400`.
  - If part not found: `404`.

### 4) Transactions (important)

The status update must be transactional:
- If the update succeeds, the audit row must exist.
- If the update fails, there must be no partial writes.

### 5) Timeouts, logging, and context

Keep this lightweight:
- Set reasonable `http.Server` timeouts (e.g., `ReadHeaderTimeout`, `ReadTimeout`, `WriteTimeout`, `IdleTimeout`).
- Use structured-ish logging (even a simple JSON line per request is fine).
- Use request contexts (`r.Context()`) in handlers and middleware.

We may discuss these choices in the interview.

### 6) Testing (required)

Include tests that can be run with:

- `go test ./...`

Minimum:
- one handler test using `net/http/httptest`
- one test that covers transactional behavior (e.g., rollback on error)

---

## Optional (pick at most one)

If you have time remaining, pick **at most one** of these (do not exceed the timebox):

1) **Async event processing** for status changes
   - On successful status update, enqueue an in-memory event
   - Process events in a background goroutine via buffered channel
   - Simulate downstream action (e.g., sleep + counter)
   - Add `events_processed` (int) to `/debug/stats`

2) **Idempotency** for `PATCH /v1/parts/{id}/status` using `Idempotency-Key` header
   - Prevent duplicate audit rows for replayed requests
   - Store key + part_id in an `idempotency_keys` table

3) **Rate limiting** middleware for `/v1/*` endpoints
   - Token-bucket or sliding-window approach
   - Per auth token
   - Return `429` when limited with `Retry-After` header
   - Must be concurrency-safe

4) **Cursor pagination** for `GET /v1/parts` using `limit` + `cursor`

5) **Debug metrics endpoint** `GET /debug/metrics` with simple counters/latencies as JSON

Optional items are not required for a passing submission; they are a way to show additional initiative.

---

## What we will evaluate

We are primarily looking for:

- Correctness (especially transactions)
- Code clarity and structure
- Pragmatic engineering judgment
- Tests that provide confidence
- Thoughtful tradeoffs documented in your README (brief is fine)

---

## Submission

Please submit:

- A link to a git repository OR a zipped directory
- A short README that includes:
  - how to run (`go run ...`)
  - how to test
  - any tradeoffs / next steps (bullet points are fine)

---

## Notes / Out of scope

To keep this small, please do **not** spend time on:

- Docker / docker-compose
- gRPC
- Redis / Kafka / NATS / RabbitMQ
- Full OAuth/JWT flows
- Full observability stacks (Prometheus, OpenTelemetry exporters, etc.)

Use standard libraries where possible; small helper libs are fine.

