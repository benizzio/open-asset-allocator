-- Script to delete a portfolio and all its dependent records.
-- Uses a temp table variable pattern to avoid typing the portfolio ID multiple times.
--
-- Usage:
--   1. Set the portfolio_id value in the INSERT statement below
--   2. Run the entire script within a transaction
--
-- Deletion order follows foreign key dependencies:
--   1. planned_allocation (depends on allocation_plan)
--   2. allocation_plan (references portfolio)
--   3. portfolio_allocation_fact (references portfolio)
--   4. portfolio (the main table)
--
-- Author: Github Copilot

CREATE TEMP TABLE temp_vars (key text PRIMARY KEY, value text);

INSERT INTO temp_vars (key, value)
SELECT 'portfolio_id', id::text
FROM portfolio
WHERE id IN (<REPLACE_WITH_COMMA_SEPARATED_IDS>);

-- Verify portfolio(s) exist before deletion
SELECT id, name
FROM portfolio
WHERE id IN (SELECT value::integer FROM temp_vars WHERE key = 'portfolio_id');

-- Delete planned_allocation records (child of allocation_plan)
DELETE FROM planned_allocation
WHERE allocation_plan_id IN (
    SELECT id
    FROM allocation_plan
    WHERE portfolio_id IN (SELECT value::integer FROM temp_vars WHERE key = 'portfolio_id')
);

-- Delete allocation_plan records
DELETE FROM allocation_plan
WHERE portfolio_id IN (SELECT value::integer FROM temp_vars WHERE key = 'portfolio_id');

-- Delete portfolio_allocation_fact records
DELETE FROM portfolio_allocation_fact
WHERE portfolio_id IN (SELECT value::integer FROM temp_vars WHERE key = 'portfolio_id');

-- Delete the portfolio(s)
DELETE FROM portfolio
WHERE id IN (SELECT value::integer FROM temp_vars WHERE key = 'portfolio_id');

-- Verify deletion completed
SELECT 'Portfolio(s) deleted successfully. Remaining portfolios:' AS status;

SELECT id, name FROM portfolio ORDER BY id;

DELETE FROM portfolio_allocation_obs_time
WHERE id NOT IN (
    SELECT DISTINCT observation_time_id
    FROM portfolio_allocation_fact
    WHERE observation_time_id IS NOT NULL
);

