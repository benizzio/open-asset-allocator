CREATE TEMP TABLE temp_vars (key text PRIMARY KEY, value text);

INSERT INTO temp_vars (key, value) VALUES ('portfolio_id', '1');

INSERT INTO portfolio ("name", allocation_structure)
SELECT "name", allocation_structure
FROM portfolio
WHERE id = (SELECT value::integer FROM temp_vars WHERE key = 'portfolio_id')
;

INSERT INTO temp_vars (key, value)
SELECT 'new_portfolio_id', id::text
FROM portfolio
WHERE id = (SELECT max(id) FROM portfolio)
;

UPDATE portfolio_allocation_fact paf
SET portfolio_id = (SELECT value::integer FROM temp_vars WHERE key = 'new_portfolio_id')
WHERE portfolio_id = (SELECT value::integer FROM temp_vars WHERE key = 'portfolio_id')
;

UPDATE allocation_plan ap
SET portfolio_id = (SELECT value::integer FROM temp_vars WHERE key = 'new_portfolio_id')
WHERE portfolio_id = (SELECT value::integer FROM temp_vars WHERE key = 'portfolio_id')
;