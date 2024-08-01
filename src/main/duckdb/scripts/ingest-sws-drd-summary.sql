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
install prql from community;
install postgres;
load prql;
load postgres;

ATTACH '' AS pgsql (TYPE POSTGRES);

.print '=> Starting data ingestion transaction'
BEGIN TRANSACTION;

.print '=> Inserting asset data into the asset table the WHEN it does not exist'
INSERT INTO pgsql.asset (ticker)
    SELECT swss.asset AS ticker FROM sws_summary swss
    LEFT JOIN pgsql.asset ass ON swss.asset = ass.ticker
    WHERE ass.ticker IS NULL
;

.print '=> Inserting portfolio data in the asset fact table, joining data classification from asset dimension mapping classifier file'
-- Has to fail when the asset dimension data is missing to indicate that file is broken and needs more data
-- INSERT INTO pgsql.asset_value_fact
    -- (asset_id, class, cash_reserve, asset_quantity, asset_market_price, total_market_value)

    -- TODO fix this query
    SELECT
        ass.id,
        adm.class,
        adm.cash_reserve,
        nullif(adm.asset_quantity, -1),
        constant_or_null(swss.current_price, adm.asset_quantity),
        swss.current_value
    FROM sws_summary swss
    LEFT JOIN asset_dimension_mapping adm ON adm.ticker = swss.asset
    LEFT JOIN pgsql.asset ass ON ass.ticker = swss.asset
;

-- (|
--     from sws_summary
--     select {
--         asset_ticker = asset,
--         total_market_value = current_value
--     }
-- |);

COMMIT;

SELECT * FROM pgsql.asset;
