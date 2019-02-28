-- +goose Up
CREATE TABLE ens.auction_started (
  id                SERIAL PRIMARY KEY,
  header_id         INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
  hash              CHARACTER VARYING(66) NOT NULL,
  registration_date INTEGER NOT NULL,
  tx_idx            INTEGER NOT NUll,
  log_idx           INTEGER NOT NUll,
  raw_log           JSONB,
  UNIQUE (header_id, tx_idx, log_idx)
);

ALTER TABLE public.checked_headers
  ADD COLUMN auction_started_checked INTEGER NOT NULL DEFAULT 0;


-- +goose Down
DROP TABLE ens.auction_started;

ALTER TABLE public.checked_headers
  DROP COLUMN auction_started_checked;
