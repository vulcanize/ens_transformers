-- +goose Up
CREATE TABLE ens.new_bid (
  id                SERIAL PRIMARY KEY,
  header_id         INTEGER NOT NULL REFERENCES headers (id) ON DELETE CASCADE,
  hash              CHARACTER VARYING(66) NOT NULL,
  bidder            CHARACTER VARYING(66) NOT NULL,
  deposit           NUMERIC NOT NULL,
  tx_idx            INTEGER NOT NUll,
  log_idx           INTEGER NOT NUll,
  raw_log           JSONB,
  UNIQUE (header_id, tx_idx, log_idx)
);

ALTER TABLE public.checked_headers
  ADD COLUMN new_bid_checked INTEGER NOT NULL DEFAULT 0;


-- +goose Down
DROP TABLE ens.new_bid;

ALTER TABLE public.checked_headers
  DROP COLUMN new_bid_checked;
