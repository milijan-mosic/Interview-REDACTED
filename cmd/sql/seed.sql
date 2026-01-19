-- seed.sql
-- Deterministic seed dataset for the take-home exercise (SQLite)

DELETE FROM idempotency_keys;
DELETE FROM part_status_audit;
DELETE FROM parts;

INSERT INTO parts (id, name, status, supplier, material, weight, critical, updated_at) VALUES
  ('part-001', 'Mounting Bracket', 'Approved',  'Internal',          'Aluminum',  1.2, 1, '2026-01-01T00:00:00Z'),
  ('part-002', 'Gear Housing',     'In Review', 'Acme Manufacturing','Steel',     3.8, 0, '2026-01-01T00:00:00Z'),
  ('part-003', 'Drive Shaft',      'Draft',     'Internal',         'Steel',     2.5, 0, '2026-01-01T00:00:00Z'),
  ('part-004', 'Retaining Clip',   'Rejected',  'Beta Metals',      'Titanium',  0.1, 0, '2026-01-01T00:00:00Z'),
  ('part-005', 'Spacer Ring',      'Approved',  'Gamma Works',      'Nylon',     0.2, 0, '2026-01-01T00:00:00Z'),
  ('part-006', 'Sensor Mount',     'Draft',     'Internal',         'Aluminum',  0.7, 0, '2026-01-01T00:00:00Z');
