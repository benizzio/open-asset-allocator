.bail on
.changes on

.print '=> Loading external modules and databases needed for ingestion'
install prql from community;
install postgres;
install scrooge from community;

load prql;
load scrooge;
load postgres;

.print '=> Reading the Ghostfolio activity file'
CREATE TEMP TABLE ghostf_activity (
    fee NUMERIC(18,8),
    quantity NUMERIC(18,8),
    type TEXT,
    unitPrice NUMERIC(18,8),
    datasource TEXT,
    date TIMESTAMP,
    symbol TEXT
);

INSERT INTO ghostf_activity
    SELECT
        activity_struct.fee,
        activity_struct.quantity,
        activity_struct.type,
        activity_struct.unitPrice,
        activity_struct.datasource,
        activity_struct.date,
        activity_struct.symbol
    FROM (
        SELECT unnest(activities) AS activity_struct
        FROM read_json(format('{}/ghostfolio-export-*.json', getenv('INTERNAL_DUCKDB_INPUT_PATH')))
    )
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

-- to debug if needed
-- select * from ghostf_activity;
-- select DISTINCT type from ghostf_activity;
-- select DISTINCT currency from ghostf_activity;
-- select * from asset_dimension_mapping;

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
    SELECT amdsi.asset_id, 'YAHOO'
    FROM asset_market_data_source_insertion amdsi
    LEFT JOIN pgsql.asset_market_data_source amds ON amds.asset_id = amdsi.asset_id AND amds.data_source = 'YAHOO'
    WHERE amdsi.data_source != 'YAHOO' AND amds.id IS NULL
;

-- to debug if needed
-- select * from pgsql.asset_market_data_source;

.print '=> Creating temporary view for obtaining market price data on Yahoo'
CREATE TEMP VIEW yahoo_asset_list AS
    SELECT ass.id AS asset_id, split_part(ass.ticker, '.', 1) AS yahoo_ticker, amds.id AS asset_data_source_id
    FROM pgsql.asset_market_data_source amds
    JOIN pgsql.asset ass ON amds.asset_id = ass.id
    WHERE amds.data_source = 'YAHOO'
;

-- to debug if needed
-- SELECT * FROM yahoo_asset_list;

.print '=> Creating temporary table and view for market data from Yahoo'
SET VARIABLE yahoo_ticker_list = (SELECT LIST(yahoo_ticker) FROM yahoo_asset_list);

CREATE TEMP TABLE yahoo_finance_data AS
    SELECT symbol, Date[2] as last_date, Close[2] as last_close
    FROM yahoo_finance(
        getvariable('yahoo_ticker_list'),
        (current_date() - INTERVAL 1 DAY)::DATE,
        current_date(),
        "1d"
    )
;

CREATE TEMP VIEW yahoo_current_asset_price AS
    SELECT yal.asset_data_source_id, yf.last_date AS market_date, yf.last_close AS close_price
    FROM yahoo_asset_list yal
    JOIN yahoo_finance_data yf ON yf.symbol = yal.yahoo_ticker
    WHERE yf.last_date = current_date()
;

.print '=> Reading and registering (upserting) market data from Yahoo'
INSERT INTO pgsql.asset_price_market_data (asset_data_source_id, market_date, market_close_price)
    SELECT ycap.asset_data_source_id, ycap.market_date, ycap.close_price FROM yahoo_current_asset_price ycap
    LEFT JOIN pgsql.asset_price_market_data apmd
        ON apmd.asset_data_source_id = ycap.asset_data_source_id AND apmd.market_date = ycap.market_date
    WHERE apmd.asset_data_source_id IS NULL
;

UPDATE pgsql.asset_price_market_data apmd SET market_close_price = ycap.close_price
    FROM yahoo_current_asset_price ycap
    WHERE apmd.asset_data_source_id = ycap.asset_data_source_id AND apmd.market_date = ycap.market_date
;

SELECT * FROM pgsql.asset_price_last_market_data;

.print '=> Mapping and reducing ghostfolio data to calculate asset values, joining with asset dimension mapping'
CREATE TEMP TABLE ghostf_symbol_aggegation (
    symbol TEXT,
    total_quantity NUMERIC(18,8),
    total_fee NUMERIC(18,8),
    class TEXT,
    cash_reserve BOOLEAN
    )
;

INSERT INTO ghostf_symbol_aggegation
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
        filter total_quantity > 0
        join side:left asset_dimension_mapping (ghostf_activity.symbol == asset_dimension_mapping.ticker)
        select {
            symbol,
            total_quantity = case [ `asset_quantity` > 0 => asset_quantity, `asset_quantity` <= 0 => total_quantity ],
            total_fee,
            class,
            cash_reserve
        }
    |)
;

SELECT * FROM ghostf_symbol_aggegation;

.print '=> Creating temporary view for portfolio allocation fact insertion'
CREATE TEMP VIEW portfolio_allocation_fact_insertion AS
    SELECT
        aplmd.asset_id,
        gsa.class,
        gsa.cash_reserve,
        gsa.total_quantity AS asset_quantity,
        aplmd.market_close_price AS asset_market_price,
        gsa.total_quantity::DECIMAL(30,8) * aplmd.market_close_price::DECIMAL(30,8) AS total_market_value,
        extract('year' FROM current_date) || lpad(extract('month' FROM current_date)::text, 2, '0') as time_frame_tag,
        getenv('PORTFOLIO_ID')::INTEGER AS portfolio_id
    FROM ghostf_symbol_aggegation gsa
    LEFT JOIN pgsql.asset_price_last_market_data aplmd ON gsa.symbol = aplmd.ticker
    JOIN pgsql.asset ass ON aplmd.asset_id = ass.id
    WHERE aplmd.data_source = 'YAHOO'
;

SELECT * FROM portfolio_allocation_fact_insertion;

-- TODO insert current asset position based on last prices
.print '=> Inserting portfolio data in the asset fact table'
-- Has to fail when the asset dimension data is missing to indicate that file is broken and needs more data
INSERT INTO pgsql.portfolio_allocation_fact (
        asset_id,
        class,
        cash_reserve,
        asset_quantity,
        asset_market_price,
        total_market_value,
        time_frame_tag,
        portfolio_id
    )
    SELECT
        asset_id,
        class,
        cash_reserve,
        asset_quantity,
        asset_market_price,
        total_market_value,
        time_frame_tag,
        portfolio_id
    FROM portfolio_allocation_fact_insertion
;

COMMIT;

-- .print '===> DEBUGGING'
-- .print 'asset table'
-- SELECT * FROM pgsql.asset;
-- .print 'portfolio_allocation_fact table'
-- SELECT * FROM pgsql.portfolio_allocation_fact;
