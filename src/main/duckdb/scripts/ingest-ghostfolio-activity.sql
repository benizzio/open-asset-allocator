.bail on
.changes on

.print '=> Loading external modules and databases needed for ingestion'
install prql from community;
install postgres;
-- install httpfs;
SET custom_extension_repository='https://scrooge-duck.s3.us-west-2.amazonaws.com';
SET allow_extensions_metadata_mismatch=true;
install scrooge; -- from community;

load prql;
load scrooge;
load postgres;
-- load httpfs;

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

.print '=> Starting data ingestion transaction'
BEGIN TRANSACTION;

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

.print '=> Inserting asset datasource data'
CREATE TEMP VIEW asset_market_data_source_insertion AS
   SELECT DISTINCT ass.id AS asset_id, ga.dataSource AS data_source FROM ghostf_activity ga
   JOIN pgsql.asset ass ON ga.symbol = ass.ticker
   LEFT JOIN pgsql.asset_market_data_source amds ON amds.asset_id = ass.id AND amds.data_source = ga.dataSource
   WHERE amds.data_source IS NULL
;

INSERT INTO pgsql.asset_market_data_source (asset_id, data_source)
    SELECT asset_id, data_source FROM asset_market_data_source_insertion
    -- Inserting yahoo equivalents to fetch all data from there for now
    UNION
    SELECT asset_id, 'YAHOO' FROM asset_market_data_source_insertion WHERE data_source != 'YAHOO'
;

.print '=> Creating temporary view for obtaining market pice data on Yahoo'
CREATE TEMP VIEW yahoo_asset_list AS
    SELECT ass.id AS asset_id, split_part(ass.ticker, '.', 1) AS yahoo_ticker, amds.id AS asset_data_source_id
    FROM pgsql.asset_market_data_source amds
    JOIN pgsql.asset ass ON amds.asset_id = ass.id
    WHERE amds.data_source = 'YAHOO'
;

-- to debug if needed
-- SELECT * FROM yahoo_asset_list;

.print '=> Creating temporary table and view for market data from Yahoo'
-- function "yahoo_finance" does not accept dynamic values, constant for now
CREATE TEMP TABLE yahoo_finance_data AS
    SELECT symbol, Date, Close
    FROM yahoo_finance(
        [
            "SOL-USD",
            "UNI7083-USD",
            "LINK-USD",
            "MANA-USD",
            "BTC-USD",
            "ETH-USD",
            "XRP-USD",
            "AXS-USD",
            "USDC-USD",
            "LTC-USD",
            "XTZ-USD",
            "XLM-USD",
            "SAND-USD",
            "AUDIO-USD",
            "MATIC-USD",
            "MKR-USD"
        ],
        (current_date() - INTERVAL 1 DAY)::DATE,
        current_date(),
        "1d"
    )
;

CREATE TEMP VIEW yahoo_current_asset_price AS
    SELECT yal.asset_data_source_id, yf.Date AS market_date, yf.Close AS close_price FROM yahoo_asset_list yal
    JOIN yahoo_finance_data yf ON yf.symbol = yal.yahoo_ticker
    WHERE yf.Date = current_date()
;

.print '=> Reading and registering (upserting) market data from Yahoo'
INSERT INTO pgsql.asset_price_market_data (asset_data_source_id, market_date, market_close_price)
    SELECT ycap.asset_data_source_id, ycap.market_date, ycap.close_price FROM yahoo_current_asset_price ycap
    LEFT JOIN pgsql.asset_price_market_data apmd ON apmd.asset_data_source_id = ycap.asset_data_source_id AND apmd.market_date = ycap.market_date
    WHERE apmd.asset_data_source_id IS NULL
;

UPDATE pgsql.asset_price_market_data apmd SET market_close_price = ycap.close_price
    FROM yahoo_current_asset_price ycap
    WHERE apmd.asset_data_source_id = ycap.asset_data_source_id AND apmd.market_date = ycap.market_date
;

SELECT * FROM pgsql.asset_price_last_market_data;

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

COMMIT;
