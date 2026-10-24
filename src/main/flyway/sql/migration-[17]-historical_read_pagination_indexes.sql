-- Authored by: OpenCode
CREATE INDEX portfolio_allocation_obs_time_page_idx
    ON portfolio_allocation_obs_time (portfolio_id, observation_timestamp DESC, id DESC)
;

-- Authored by: OpenCode
CREATE INDEX portfolio_allocation_fact_observation_page_idx
    ON portfolio_allocation_fact (portfolio_id, observation_time_id)
;

-- Authored by: OpenCode
CREATE INDEX allocation_plan_page_idx
    ON allocation_plan (portfolio_id, type, planned_date DESC, id DESC)
;
