ALTER TABLE asset
ALTER COLUMN name DROP NOT NULL;

ALTER TABLE asset
ADD CONSTRAINT asset_ticker_uk UNIQUE (ticker);