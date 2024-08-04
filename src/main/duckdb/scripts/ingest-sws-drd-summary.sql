.bail on
.changes on

.print '=> Reading the Simply Wall Street Summary file'
CREATE TEMP TABLE sws_summary AS
    SELECT asset, total_shares, current_price, current_value
    FROM read_csv(
        '/duckdb/input/us-complete-portfolio-summary.csv',
        columns = {
            'asset': 'TEXT',
            'total_bought': 'NUMERIC',
            'total_shares': 'SMALLINT',
            'current_price': 'NUMERIC',
            'current_value': 'NUMERIC',
            'capital_gains': 'NUMERIC',
            'dividends': 'NUMERIC',
            'total_gain_currency': 'NUMERIC',
            'average_years': 'SMALLINT',
            'total_return': 'TEXT'
        }
    )
    WHERE current_value > 0
;

.print '=> Reading the asset dimension mapping classifier file'
CREATE TEMP TABLE asset_dimension_mapping AS
    SELECT *
    FROM read_csv(
        '/duckdb/input/asset-dimension-mapping.csv',
        columns = {
            'ticker': 'TEXT',
            'class': 'TEXT',
            'asset_quantity': 'NUMERIC',
            'cash_reserve': 'BOOLEAN'
        }
    )
;

.print '=> Loading external modules and databases needed for ingestion'
-- install prql from community;
install postgres;
-- load prql;
load postgres;

ATTACH '' AS pgsql (TYPE POSTGRES);

.print '=> Starting data ingestion transaction'
BEGIN TRANSACTION;

.print '=> Asset data to be inserted into the asset table the (WHEN it does not exist)'
CREATE TEMP VIEW asset_insertion AS
    SELECT swss.asset AS ticker FROM sws_summary swss
    LEFT JOIN pgsql.asset ass ON swss.asset = ass.ticker
    WHERE ass.ticker IS NULL
;

SELECT * FROM asset_insertion;

.print '=> Inserting asset data'
INSERT INTO pgsql.asset (ticker)
    SELECT ticker FROM asset_insertion
;

.print '=> Portfolio data to be inserted, joining data classification from asset dimension mapping classifier file'
CREATE TEMP VIEW asset_value_fact_insertion AS
    SELECT
        ass.id as asset_id,
        adm.class as class,
        adm.cash_reserve as cash_reserve,
        if(adm.asset_quantity > 0, adm.asset_quantity, swss.total_shares) as asset_quantity,
        swss.current_price as asset_market_price,
        if(adm.asset_quantity > 0, adm.asset_quantity * swss.current_price, swss.current_value) as total_market_value
    FROM sws_summary swss
    LEFT JOIN asset_dimension_mapping adm ON adm.ticker = swss.asset
    LEFT JOIN pgsql.asset ass ON ass.ticker = swss.asset
;

SELECT * FROM asset_value_fact_insertion;

.print '=> Inserting portfolio data in the asset fact table'
-- Has to fail when the asset dimension data is missing to indicate that file is broken and needs more data
 INSERT INTO pgsql.asset_value_fact
    SELECT * FROM asset_value_fact_insertion
;

-- (|
--     from sws_summary
--     select {
--         asset_ticker = asset,
--         total_market_value = current_value
--     }
-- |);

COMMIT;

.print '===> DEBUGGING'
.print 'asset table'
SELECT * FROM pgsql.asset;
.print 'asset_value_fact table'
SELECT * FROM pgsql.asset_value_fact;
