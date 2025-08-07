ALTER TABLE portfolio_allocation_fact DROP CONSTRAINT portfolio_allocation_fact_pk;

ALTER TABLE portfolio_allocation_fact
    ADD CONSTRAINT portfolio_allocation_fact_pk PRIMARY KEY (portfolio_id, class, cash_reserve, asset_id, observation_time_id)
;

ALTER TABLE portfolio_allocation_fact DROP COLUMN time_frame_tag;