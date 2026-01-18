-- Migration: Add field size constraints to text fields
-- This migration adds VARCHAR constraints to text fields to enforce size limits
-- that will be synchronized with front-end and back-end validation

-- Drop views that depend on asset.ticker before altering column types
DROP VIEW IF EXISTS asset_price_last_market_data;
DROP VIEW IF EXISTS asset_ticker_market_data_source;

-- Asset table: ticker and name constraints
ALTER TABLE asset
    ALTER COLUMN ticker SET DATA TYPE varchar(20),
    ALTER COLUMN name SET DATA TYPE varchar(100);

-- Recreate views that were dropped
CREATE VIEW asset_ticker_market_data_source AS
    SELECT ass.ticker, amds.data_source
    FROM asset_market_data_source amds
    JOIN asset ass ON amds.asset_id = ass.id
;

CREATE VIEW asset_price_last_market_data AS
    SELECT DISTINCT ass.id AS asset_id, ass.ticker, amds.id AS asset_data_source_id, amds.data_source, apmd.market_date, apmd.market_close_price
    FROM asset_market_data_source amds
    JOIN asset ass ON amds.asset_id = ass.id
    LEFT JOIN (
        SELECT asset_data_source_id, max(market_date) as max_marked_date FROM asset_price_market_data
        GROUP BY asset_data_source_id
    ) apmdmd ON amds.id = apmdmd.asset_data_source_id
    LEFT JOIN asset_price_market_data apmd ON amds.id = apmd.asset_data_source_id AND apmd.market_date = apmdmd.max_marked_date
;

-- Portfolio table: name constraint
ALTER TABLE portfolio
    ALTER COLUMN name SET DATA TYPE varchar(100);

-- Allocation plan table: name and type constraints
ALTER TABLE allocation_plan
    ALTER COLUMN name SET DATA TYPE varchar(100),
    ALTER COLUMN type SET DATA TYPE varchar(50);

-- Portfolio allocation fact table: class constraint
ALTER TABLE portfolio_allocation_fact
    ALTER COLUMN class SET DATA TYPE varchar(100);

-- Portfolio allocation observation time table: observation_time_tag constraint
ALTER TABLE portfolio_allocation_obs_time
    ALTER COLUMN observation_time_tag SET DATA TYPE varchar(100);
