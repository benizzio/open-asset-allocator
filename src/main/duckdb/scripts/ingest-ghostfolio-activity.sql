.bail on
.changes on

CREATE TEMP TABLE ghostf_activity AS
    SELECT activity_struct.* FROM (
        SELECT unnest(activities) as activity_struct
        FROM read_json(format('{}/ghostfolio-export-*.json', getenv('INTERNAL_DUCKDB_INPUT_PATH')))
    )
;

-- TODO to debug, remove
select * from ghostf_activity;

install prql from community;
load prql;

-- TODO PRQL to reduce activities to asset values
-- EX:
-- (|
--     from sws_summary
--     select {
--         asset_ticker = asset,
--         total_market_value = current_value
--     }
-- |);

