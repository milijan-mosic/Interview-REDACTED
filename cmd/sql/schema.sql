-- schema.sql
-- Minimal schema for the take-home exercise (SQLite)

PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS parts (
  id         TEXT PRIMARY KEY,
  name       TEXT NOT NULL,
  status     TEXT NOT NULL,
  supplier   TEXT NOT NULL,
  material   TEXT NOT NULL,
  weight     REAL NOT NULL,
  critical   INTEGER NOT NULL DEFAULT 0,
  updated_at TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS part_status_audit (
  id         INTEGER PRIMARY KEY AUTOINCREMENT,
  part_id    TEXT NOT NULL,
  old_status TEXT NOT NULL,
  new_status TEXT NOT NULL,
  changed_at TEXT NOT NULL,
  request_id TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_part_status_audit_part_id_changed_at
  ON part_status_audit(part_id, changed_at);

-- Optional: uncomment if implementing idempotency (optional feature)
CREATE TABLE IF NOT EXISTS idempotency_keys (
  key        TEXT PRIMARY KEY,
  part_id    TEXT NOT NULL,
  created_at TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_parts_status
  ON parts(status);
