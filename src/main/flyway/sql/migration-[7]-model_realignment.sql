ALTER TABLE allocation_plan_unit RENAME TO planned_allocation;
ALTER TABLE planned_allocation RENAME CONSTRAINT allocation_plan_unit_pk TO planned_allocation_pk;
ALTER TABLE planned_allocation RENAME CONSTRAINT allocation_plan_unit_uk TO planned_allocation_uk;

COMMENT ON COLUMN planned_allocation.structural_id
        IS E'Identifier of the planned allocation structure inside the hierarchical classification levels';

ALTER TABLE asset_value_fact RENAME TO portfolio_allocation_fact;
ALTER TABLE portfolio_allocation_fact RENAME CONSTRAINT asset_value_fact_pk TO portfolio_allocation_fact_pk;

CREATE TABLE portfolio (
    id serial NOT NULL,
    name text NOT NULL,
    allocation_structure jsonb NOT NULL,
    CONSTRAINT portfolio_pk PRIMARY KEY (id)
);

COMMENT ON COLUMN portfolio.allocation_structure
        IS E'Definition of the structure of the asset allocation in hierarchical levels';

ALTER TABLE allocation_plan
    DROP COLUMN structure,
    ADD COLUMN portfolio_id integer NOT NULL
;

ALTER TABLE allocation_plan ADD CONSTRAINT portfolio_fk FOREIGN KEY (portfolio_id)
    REFERENCES portfolio (id) MATCH FULL
    ON DELETE RESTRICT ON UPDATE CASCADE
;

ALTER TABLE portfolio_allocation_fact
    ADD COLUMN portfolio_id integer NOT NULL
;

ALTER TABLE portfolio_allocation_fact ADD CONSTRAINT portfolio_fk FOREIGN KEY (portfolio_id)
    REFERENCES portfolio (id) MATCH FULL
    ON DELETE RESTRICT ON UPDATE CASCADE
;

ALTER TABLE portfolio_allocation_fact DROP CONSTRAINT portfolio_allocation_fact_pk;

ALTER TABLE portfolio_allocation_fact
    ADD CONSTRAINT portfolio_allocation_fact_pk PRIMARY KEY (portfolio_id, class, cash_reserve, asset_id, time_frame_tag)
;