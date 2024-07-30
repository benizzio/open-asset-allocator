CREATE TEMP TABLE sws_summary AS
    SELECT asset, current_value
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

install prql from community;
install postgres;
load prql;
load postgres;

ATTACH '' AS pgsql (TYPE POSTGRES);

BEGIN TRANSACTION;

INSERT INTO pgsql.asset (ticker)
    SELECT swss.asset AS ticker FROM sws_summary swss
    LEFT JOIN pgsql.asset ass ON swss.asset = ass.ticker
    WHERE ass.ticker IS NULL
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
