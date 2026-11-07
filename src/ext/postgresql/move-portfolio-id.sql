-- Authored by: OpenCode and benizzio
BEGIN TRANSACTION;

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

CREATE TEMP TABLE moved_observation_ids (
    original_observation_id INTEGER PRIMARY KEY,
    new_observation_id INTEGER NOT NULL UNIQUE
);

INSERT INTO moved_observation_ids (original_observation_id, new_observation_id)
SELECT observation.id, nextval('portfolio_allocation_obs_time_id_seq')::INTEGER
FROM portfolio_allocation_obs_time observation
WHERE observation.portfolio_id = (SELECT value::integer FROM temp_vars WHERE key = 'portfolio_id')
  AND EXISTS (
      SELECT 1
      FROM portfolio_allocation_fact fact
      WHERE fact.portfolio_id = observation.portfolio_id
        AND fact.observation_time_id = observation.id
  );

INSERT INTO portfolio_allocation_obs_time (
    id,
    portfolio_id,
    observation_time_tag,
    observation_timestamp
)
SELECT moved.new_observation_id,
       (SELECT value::integer FROM temp_vars WHERE key = 'new_portfolio_id'),
       observation.observation_time_tag,
       observation.observation_timestamp
FROM moved_observation_ids moved
JOIN portfolio_allocation_obs_time observation ON observation.id = moved.original_observation_id;

UPDATE portfolio_allocation_fact paf
SET portfolio_id = (SELECT value::integer FROM temp_vars WHERE key = 'new_portfolio_id')
    , observation_time_id = moved.new_observation_id
FROM moved_observation_ids moved
WHERE paf.portfolio_id = (SELECT value::integer FROM temp_vars WHERE key = 'portfolio_id')
  AND paf.observation_time_id = moved.original_observation_id
;

UPDATE allocation_plan ap
SET portfolio_id = (SELECT value::integer FROM temp_vars WHERE key = 'new_portfolio_id')
WHERE portfolio_id = (SELECT value::integer FROM temp_vars WHERE key = 'portfolio_id')
;

DELETE FROM portfolio_allocation_obs_time
WHERE portfolio_id = (SELECT value::integer FROM temp_vars WHERE key = 'portfolio_id');

DELETE FROM portfolio
WHERE id = (SELECT value::integer FROM temp_vars WHERE key = 'portfolio_id')
;

COMMIT;
