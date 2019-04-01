-- +goose Up
CREATE TABLE ens.name_changed (
  id                SERIAL PRIMARY KEY,
  header_id         INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
  resolver          CHARACTER VARYING(66) NOT NULL,
  node              CHARACTER VARYING(66) NOT NULL,
  name              TEXT NOT NULL,
  tx_idx            INTEGER NOT NUll,
  log_idx           INTEGER NOT NUll,
  raw_log           JSONB,
  UNIQUE (header_id, tx_idx, log_idx)
);

ALTER TABLE public.checked_headers
  ADD COLUMN name_changed_checked INTEGER NOT NULL DEFAULT 0;


-- +goose Down
DROP TABLE ens.name_changed;

ALTER TABLE public.checked_headers
  DROP COLUMN name_changed_checked;
