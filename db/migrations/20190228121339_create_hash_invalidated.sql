-- +goose Up
CREATE TABLE ens.hash_invalidated (
  id                SERIAL PRIMARY KEY,
  header_id         INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
  hash              CHARACTER VARYING(66) NOT NULL,
  name              TEXT NOT NULL,
  value             NUMERIC NOT NULL,
  registration_date NUMERIC NOT NULL,
  tx_idx            INTEGER NOT NUll,
  log_idx           INTEGER NOT NUll,
  raw_log           JSONB,
  UNIQUE (header_id, tx_idx, log_idx)
);

ALTER TABLE public.checked_headers
  ADD COLUMN hash_invalidated_checked INTEGER NOT NULL DEFAULT 0;


-- +goose Down
DROP TABLE ens.hash_invalidated;

ALTER TABLE public.checked_headers
  DROP COLUMN hash_invalidated_checked;
