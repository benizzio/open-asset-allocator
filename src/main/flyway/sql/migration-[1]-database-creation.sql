-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler version: 1.1.3
-- PostgreSQL version: 16.0
-- Project Site: pgmodeler.io
-- Model Author: ---

-- Database creation must be performed outside a multi lined SQL file.
-- These commands were put in this file only as a convenience.
--
-- object: postgres | type: DATABASE --
-- DROP DATABASE IF EXISTS postgres;
--CREATE DATABASE postgres;
-- ddl-end --


-- object: asset | type: TABLE --
-- DROP TABLE IF EXISTS asset CASCADE;
CREATE TABLE asset (
	id serial NOT NULL,
	ticker text NOT NULL,
	name text NOT NULL,
	CONSTRAINT asset_pk PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON COLUMN asset.ticker IS E'Exchange full identifier, if relevant';
-- ddl-end --
COMMENT ON COLUMN asset.name IS E'Underlying asset name (company name, fund name, currency name, etc.)';
-- ddl-end --
ALTER TABLE asset OWNER TO postgres;
-- ddl-end --

-- object: asset_value_fact | type: TABLE --
-- DROP TABLE IF EXISTS asset_value_fact CASCADE;
CREATE TABLE asset_value_fact (
	asset_id integer NOT NULL,
	class text NOT NULL,
	cash_reserve boolean NOT NULL,
	asset_quantity numeric,
	asset_market_price numeric,
	total_market_value bigint NOT NULL,
	CONSTRAINT asset_value_fact_pk PRIMARY KEY (class,cash_reserve,asset_id)
);
-- ddl-end --
COMMENT ON TABLE asset_value_fact IS E'Fact table for portfolio asset values held';
-- ddl-end --
COMMENT ON COLUMN asset_value_fact.class IS E'Asset class degenerate dimension';
-- ddl-end --
COMMENT ON COLUMN asset_value_fact.cash_reserve IS E'Degenerate dimension informing that the asset is a cash reserve for the classifier';
-- ddl-end --
COMMENT ON COLUMN asset_value_fact.asset_quantity IS E'Helper field for asset unit quantity';
-- ddl-end --
COMMENT ON COLUMN asset_value_fact.asset_market_price IS E'Helper field for aggragated asset price';
-- ddl-end --
COMMENT ON COLUMN asset_value_fact.total_market_value IS E'Measure containing total market value of the asset in the portfolio';
-- ddl-end --
ALTER TABLE asset_value_fact OWNER TO postgres;
-- ddl-end --

-- object: asset_fk | type: CONSTRAINT --
-- ALTER TABLE asset_value_fact DROP CONSTRAINT IF EXISTS asset_fk CASCADE;
ALTER TABLE asset_value_fact ADD CONSTRAINT asset_fk FOREIGN KEY (asset_id)
REFERENCES asset (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: allocation_plan_unit | type: TABLE --
-- DROP TABLE IF EXISTS allocation_plan_unit CASCADE;
CREATE TABLE allocation_plan_unit (
	id serial NOT NULL,
	allocation_plan_id integer NOT NULL,
	structural_id text NOT NULL,
	asset_id integer,
	cash_reserve boolean NOT NULL,
	slice smallint,
	total_market_value smallint,
	CONSTRAINT slice_percentage_ck CHECK (slice BETWEEN 1 AND 100),
	CONSTRAINT allocation_plan_unit_pk PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE allocation_plan_unit IS E'Asset allocation planning detail';
-- ddl-end --
COMMENT ON COLUMN allocation_plan_unit.structural_id IS E'Identifier of the plan unit inside the hierarchical classification';
-- ddl-end --
COMMENT ON COLUMN allocation_plan_unit.cash_reserve IS E'Informs that the asset is a cash reserve for the lower classifier granularity';
-- ddl-end --
COMMENT ON COLUMN allocation_plan_unit.slice IS E'Allocation slice of the planned asset, bounded to the lower classifier granularity (in %), for ALLOCATION_PLANs';
-- ddl-end --
COMMENT ON COLUMN allocation_plan_unit.total_market_value IS E'Planned allocation size for EXECUTION_PLANs';
-- ddl-end --
ALTER TABLE allocation_plan_unit OWNER TO postgres;
-- ddl-end --

-- object: asset_fk | type: CONSTRAINT --
-- ALTER TABLE allocation_plan_unit DROP CONSTRAINT IF EXISTS asset_fk CASCADE;
ALTER TABLE allocation_plan_unit ADD CONSTRAINT asset_fk FOREIGN KEY (asset_id)
REFERENCES asset (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: allocation_plan | type: TABLE --
-- DROP TABLE IF EXISTS allocation_plan CASCADE;
CREATE TABLE allocation_plan (
	id serial NOT NULL,
	name text NOT NULL,
	type text NOT NULL,
	structure text NOT NULL,
	planned_execution_date date,
	CONSTRAINT allocation_plan_pk PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE allocation_plan IS E'Asset allocation planning classification';
-- ddl-end --
COMMENT ON COLUMN allocation_plan.type IS E'Allocation plan type, such as ALLOCATION_PLAN (slice sizing) or EXECUTION_PLAN (real asset positioning)';
-- ddl-end --
COMMENT ON COLUMN allocation_plan.structure IS E'Definition of structure of asset allocation plan in hierarchical levels, using the "|" (pipe) charcter as a divider. Ex: "ASSET_CLASS|ASSET"';
-- ddl-end --
ALTER TABLE allocation_plan OWNER TO postgres;
-- ddl-end --

-- object: allocation_plan_fk | type: CONSTRAINT --
-- ALTER TABLE allocation_plan_unit DROP CONSTRAINT IF EXISTS allocation_plan_fk CASCADE;
ALTER TABLE allocation_plan_unit ADD CONSTRAINT allocation_plan_fk FOREIGN KEY (allocation_plan_id)
REFERENCES allocation_plan (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: allocation_plan_unit_uk | type: CONSTRAINT --
-- ALTER TABLE allocation_plan_unit DROP CONSTRAINT IF EXISTS allocation_plan_unit_uk CASCADE;
ALTER TABLE allocation_plan_unit ADD CONSTRAINT allocation_plan_unit_uk UNIQUE (allocation_plan_id,structural_id);
-- ddl-end --
