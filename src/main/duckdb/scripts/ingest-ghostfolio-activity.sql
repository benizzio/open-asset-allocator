.bail on
.changes on

.print '=> Loading external modules and databases needed for ingestion'
install prql from community;
install scrooge from community;
install postgres;
load prql;
load scrooge;
load postgres;

.print '=> Reading the Ghostfolio activity file'
CREATE TEMP TABLE ghostf_activity AS
    SELECT activity_struct.* FROM (
        SELECT unnest(activities) AS activity_struct
        FROM read_json(format('{}/ghostfolio-export-*.json', getenv('INTERNAL_DUCKDB_INPUT_PATH')))
    )
;

-- to debug if needed
-- select * from ghostf_activity;
-- select DISTINCT type from ghostf_activity;
-- select DISTINCT currency from ghostf_activity;

ATTACH '' AS pgsql (TYPE POSTGRES);

.print '=> Asset data to be inserted into the asset table (WHEN it does not exist)'
CREATE TEMP VIEW asset_insertion AS
   SELECT DISTINCT ga.symbol AS ticker FROM ghostf_activity ga
   LEFT JOIN pgsql.asset ass ON ga.symbol = ass.ticker
   WHERE ass.ticker IS NULL
;

SELECT * FROM asset_insertion;

.print '=> Inserting asset data'
INSERT INTO pgsql.asset (ticker)
    SELECT ticker FROM asset_insertion
;

.print '=> Asset datasource data to be inserted into the ghostf_datasource table (WHEN it does not exist)'
CREATE TEMP VIEW asset_market_data_source_insertion AS
   SELECT DISTINCT ass.id AS asset_id, ga.dataSource AS data_source FROM ghostf_activity ga
   JOIN pgsql.asset ass ON ga.symbol = ass.ticker
   LEFT JOIN pgsql.asset_market_data_source amds ON amds.asset_id = ass.id AND amds.data_source = ga.dataSource
   WHERE amds.data_source IS NULL
;

SELECT ass.ticker, amdsi.* FROM asset_market_data_source_insertion amdsi
JOIN pgsql.asset ass ON amdsi.asset_id = ass.id
;

.print '=> Inserting asset datasource data'
INSERT INTO pgsql.asset_market_data_source (asset_id, data_source)
    SELECT asset_id, data_source FROM asset_market_data_source_insertion
;

--TODO continue: fetch prices

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

-- TODO insert current asset position based on last prices