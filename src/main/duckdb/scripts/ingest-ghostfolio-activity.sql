.bail on
.changes on

CREATE TEMP TABLE ghostf_activity AS
    SELECT activity_struct.* FROM (
        SELECT unnest(activities) as activity_struct
        FROM read_json(format('{}/ghostfolio-export-*.json', getenv('INTERNAL_DUCKDB_INPUT_PATH')))
    )
;

-- TODO to debug, remove?
-- select * from ghostf_activity;
-- select DISTINCT type from ghostf_activity;
-- select DISTINCT currency from ghostf_activity;

CREATE TEMP TABLE ghostf_datasource AS
    select DISTINCT symbol, dataSource from ghostf_activity;

-- TODO to debug, remove?
select * from ghostf_datasource;

install prql from community;
load prql;

-- TODO PRQL to reduce activities to asset values
CREATE TEMP TABLE ghostf_symbol_aggegation AS
    (|
        from ghostf_activity
        select {
            symbol,
            multiplier = case [ `type` == "BUY" => 1, `type` == "SELL" => -1 ],
            quantity,
            fee
        }
        select {
            symbol,
            quantity_mutation = quantity * multiplier,
            fee
        }
        group { symbol } (
            aggregate {
                total_quantity = sum quantity_mutation,
                total_fee = sum fee
            }
        )
|)
;

-- TODO to debug, remove?
select * from ghostf_symbol_aggegation;

--TODO continue: fetch prices
