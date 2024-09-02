DROP VIEW asset_price_last_market_data;

ALTER TABLE asset_price_market_data
    ALTER COLUMN market_close_price SET DATA TYPE numeric(18,8)
;

ALTER TABLE asset_value_fact
    ALTER COLUMN asset_quantity SET DATA TYPE numeric(18,8),
    ALTER COLUMN asset_market_price SET DATA TYPE numeric(18,8)
;

CREATE VIEW asset_price_last_market_data AS
	SELECT DISTINCT ass.id AS asset_id, ass.ticker, amds.id AS asset_data_source_id, amds.data_source, apmd.market_date, apmd.market_close_price
	FROM asset_market_data_source amds
	JOIN asset ass ON amds.asset_id = ass.id
	LEFT JOIN (
		SELECT asset_data_source_id, max(market_date) as max_marked_date FROM asset_price_market_data
		GROUP BY asset_data_source_id
	) apmdmd ON amds.id = apmdmd.asset_data_source_id
	LEFT JOIN asset_price_market_data apmd ON amds.id = apmd.asset_data_source_id AND apmd.market_date = apmdmd.max_marked_date
;