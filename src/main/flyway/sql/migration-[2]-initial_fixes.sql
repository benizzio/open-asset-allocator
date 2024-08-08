ALTER TABLE asset
ALTER COLUMN name DROP NOT NULL;

ALTER TABLE asset
ADD CONSTRAINT asset_ticker_uk UNIQUE (ticker);

ALTER TABLE asset_value_fact
ADD COLUMN time_frame_tag text NOT NULL;

COMMENT ON COLUMN asset_value_fact.time_frame_tag
IS E'Classifies a set of facts (portfolio snapshot) to a specified time frame';

ALTER TABLE asset_value_fact
ADD COLUMN create_timestamp timestamp NOT NULL DEFAULT now();

ALTER TABLE asset_value_fact DROP CONSTRAINT asset_value_fact_pk ;

ALTER TABLE asset_value_fact
ADD CONSTRAINT asset_value_fact_pk PRIMARY KEY (class, cash_reserve, asset_id, time_frame_tag);
