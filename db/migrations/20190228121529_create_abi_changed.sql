-- +goose Up
CREATE TABLE ens.abi_changed (
  id                SERIAL PRIMARY KEY,
  header_id         INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
  node              CHARACTER VARYING(66) NOT NULL,
  content_type      BIGINT NOT NULL,
  tx_idx            INTEGER NOT NUll,
  log_idx           INTEGER NOT NUll,
  raw_log           JSONB,
  UNIQUE (header_id, tx_idx, log_idx)
);

ALTER TABLE public.checked_headers
  ADD COLUMN abi_changed_checked INTEGER NOT NULL DEFAULT 0;


-- +goose Down
DROP TABLE ens.abi_changed;

ALTER TABLE public.checked_headers
  DROP COLUMN abi_changed_checked;
