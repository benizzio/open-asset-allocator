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
    WHERE current_value > 0;

install prql FROM community;
load prql;

(|
    from sws_summary
    select {
        asset_ticker = asset,
        total_market_value = current_value
    }
|);
