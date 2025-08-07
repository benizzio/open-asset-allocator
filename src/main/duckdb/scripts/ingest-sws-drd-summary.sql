.bail on
.changes on

.print '=> Loading external modules and databases needed for ingestion'
install postgres;
load postgres;

.print '=> Reading the Simply Wall Street Summary file'
CREATE TEMP TABLE sws_summary AS
    SELECT asset, total_shares, current_price, current_value
    FROM read_csv(
        format('{}/*-us-complete-portfolio-summary.csv', getenv('INTERNAL_DUCKDB_INPUT_PATH')),
        columns = {
            'asset': 'TEXT',
            'total_bought': 'NUMERIC(18,8)',
            'total_shares': 'SMALLINT',
            'current_price': 'NUMERIC(18,8)',
            'current_value': 'NUMERIC(18,8)',
            'capital_gains': 'NUMERIC(18,8)',
            'dividends': 'NUMERIC(18,8)',
            'total_gain_currency': 'NUMERIC(18,8)',
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
        format('{}/*-asset-dimension-mapping.csv', getenv('INTERNAL_DUCKDB_INPUT_PATH')),
        columns = {
            'ticker': 'TEXT',
            'class': 'TEXT',
            'asset_quantity': 'NUMERIC(18,8)',
            'cash_reserve': 'BOOLEAN'
        }
    )
;

ATTACH '' AS pgsql (TYPE POSTGRES);

.print '=> Starting data ingestion transaction'
BEGIN TRANSACTION;

.print '=> Asset data to be inserted into the asset table (WHEN it does not exist)'
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

.print '=> Creating portfolio observation time dimension record, if it does not exist'
INSERT INTO pgsql.portfolio_allocation_obs_time (observation_time_tag, observation_timestamp)
SELECT getenv('PORTFOLIO_ALLOCATION_OBS_TIME_TAG') AS observation_time_tag, current_timestamp AS observation_timestamp
    WHERE NOT EXISTS (
    SELECT 1 FROM pgsql.portfolio_allocation_obs_time
    WHERE observation_time_tag = getenv('PORTFOLIO_ALLOCATION_OBS_TIME_TAG')
)
;

.print '=> Portfolio data to be inserted, joining data classification from asset dimension mapping classifier file'
CREATE TEMP VIEW portfolio_allocation_fact_insertion AS
    SELECT
        ass.id as asset_id,
        adm.class as class,
        adm.cash_reserve as cash_reserve,
        if(adm.asset_quantity > 0, adm.asset_quantity, swss.total_shares) as asset_quantity,
        swss.current_price as asset_market_price,
        if(
            adm.asset_quantity > 0,
            adm.asset_quantity::DECIMAL(30,8) * swss.current_price::DECIMAL(30,8),
            swss.current_value
        ) as total_market_value,
        paot.id AS observation_time_id,
        getenv('PORTFOLIO_ID')::INTEGER AS portfolio_id
    FROM sws_summary swss
    LEFT JOIN asset_dimension_mapping adm ON adm.ticker = swss.asset
    LEFT JOIN pgsql.asset ass ON ass.ticker = swss.asset
    JOIN pgsql.portfolio_allocation_obs_time paot ON paot.observation_time_tag = getenv('PORTFOLIO_ALLOCATION_OBS_TIME_TAG')
;

SELECT * FROM portfolio_allocation_fact_insertion;

.print '=> Inserting portfolio data in the asset fact table'
-- Has to fail when the asset dimension data is missing to indicate that file is broken and needs more data
INSERT INTO pgsql.portfolio_allocation_fact (
        asset_id,
        class,
        cash_reserve,
        asset_quantity,
        asset_market_price,
        total_market_value,
        observation_time_id,
        portfolio_id
    )
    SELECT
        asset_id,
        class,
        cash_reserve,
        asset_quantity,
        asset_market_price,
        total_market_value,
        observation_time_id,
        portfolio_id
    FROM portfolio_allocation_fact_insertion
;

COMMIT;

-- .print '===> DEBUGGING'
-- .print 'asset table'
-- SELECT * FROM pgsql.asset;
-- .print 'portfolio_allocation_fact table'
-- SELECT * FROM pgsql.portfolio_allocation_fact;
