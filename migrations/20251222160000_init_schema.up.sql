-- conmesh: initial schema bootstrap
-- Keep this minimal. Add real tables in subsequent migrations.

CREATE TABLE IF NOT EXISTS schema_migrations_meta (
  id              BIGSERIAL PRIMARY KEY,
  created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  note            TEXT NOT NULL
);

INSERT INTO schema_migrations_meta (note)
VALUES ('init schema bootstrap');
