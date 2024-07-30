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

install prql FROM community;
install postgres;
load prql;
load postgres;

ATTACH 'dbname=postgres user=postgres password=local host=172.17.0.1' AS postgres (TYPE POSTGRES);

BEGIN TRANSACTION;

INSERT INTO postgres.asset (ticker)
    SELECT swss.asset AS ticker FROM sws_summary swss
    LEFT JOIN postgres.asset ass ON swss.asset = ass.ticker
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

SELECT * FROM postgres.asset;
