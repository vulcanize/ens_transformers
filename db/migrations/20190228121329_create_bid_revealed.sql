-- +goose Up
CREATE TABLE ens.bid_revealed (
  id                SERIAL PRIMARY KEY,
  header_id         INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
  hash              CHARACTER VARYING(66) NOT NULL,
  owner             CHARACTER VARYING(66) NOT NULL,
  value             BIGINT NOT NULL,
  status            INTEGER NOT NULL,
  tx_idx            INTEGER NOT NUll,
  log_idx           INTEGER NOT NUll,
  raw_log           JSONB,
  UNIQUE (header_id, tx_idx, log_idx)
);

ALTER TABLE public.checked_headers
  ADD COLUMN bid_revealed_checked INTEGER NOT NULL DEFAULT 0;


-- +goose Down
DROP TABLE ens.bid_revealed;

ALTER TABLE public.checked_headers
  DROP COLUMN bid_revealed_checked;
