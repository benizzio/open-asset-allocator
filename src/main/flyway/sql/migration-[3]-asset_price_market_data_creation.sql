CREATE TABLE asset_market_data_source (
	id serial NOT NULL,
	asset_id integer NOT NULL,
	data_source text NOT NULL,
	CONSTRAINT asset_market_data_source_pk PRIMARY KEY (id)
);

COMMENT ON TABLE asset_market_data_source IS E'Market data sources per asset';

ALTER TABLE asset_market_data_source ADD CONSTRAINT asset_fk FOREIGN KEY (asset_id)
	REFERENCES asset (id) MATCH FULL
	ON DELETE SET NULL ON UPDATE CASCADE
;

ALTER TABLE asset_market_data_source ADD CONSTRAINT asset_market_data_source_uk UNIQUE (asset_id, data_source);

CREATE TABLE asset_price_market_data (
	asset_data_source_id integer NOT NULL,
	market_date date NOT NULL,
	market_close_price numeric NOT NULL,
	CONSTRAINT asset_price_market_data_pk PRIMARY KEY (asset_data_source_id, market_date)
);

COMMENT ON TABLE asset_price_market_data IS E'Market data for prices per asset and datasource';

ALTER TABLE asset_price_market_data ADD CONSTRAINT asset_market_data_source_fk FOREIGN KEY (asset_data_source_id)
	REFERENCES asset_market_data_source (id) MATCH FULL
	ON DELETE SET NULL ON UPDATE CASCADE
;

CREATE VIEW asset_price_last_market_data AS
	SELECT DISTINCT ass.ticker, amds.data_source, apmd.market_date, apmd.market_close_price
	FROM asset_market_data_source amds
	JOIN asset ass ON amds.asset_id = ass.id
	LEFT JOIN (
		SELECT asset_data_source_id, max(market_date) as max_marked_date FROM asset_price_market_data
		GROUP BY asset_data_source_id
	) apmdmd ON amds.id = apmdmd.asset_data_source_id
	LEFT JOIN asset_price_market_data apmd ON amds.id = apmd.asset_data_source_id AND apmd.market_date = apmdmd.max_marked_date
;

COMMENT ON VIEW asset_price_last_market_data IS E'Last available price market data for asset that has an available data source registered';