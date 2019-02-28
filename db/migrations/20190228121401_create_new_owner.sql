-- +goose Up
CREATE TABLE ens.new_owner (
  id                SERIAL PRIMARY KEY,
  header_id         INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
  node              CHARACTER VARYING(66) NOT NULL,
  label             CHARACTER VARYING(66) NOT NULL,
  owner             CHARACTER VARYING(66) NOT NULL,
  tx_idx            INTEGER NOT NUll,
  log_idx           INTEGER NOT NUll,
  raw_log           JSONB,
  UNIQUE (header_id, tx_idx, log_idx)
);

ALTER TABLE public.checked_headers
  ADD COLUMN new_owner_checked INTEGER NOT NULL DEFAULT 0;


-- +goose Down
DROP TABLE ens.new_owner;

ALTER TABLE public.checked_headers
  DROP COLUMN new_owner_checked;
