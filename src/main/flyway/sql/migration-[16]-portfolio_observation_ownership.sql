-- Authored by: OpenCode
ALTER TABLE portfolio_allocation_obs_time
    ADD COLUMN portfolio_id INTEGER;

-- Authored by: OpenCode
SELECT setval(
    pg_get_serial_sequence('portfolio_allocation_obs_time', 'id'),
    GREATEST(
        COALESCE((SELECT MAX(id) FROM portfolio_allocation_obs_time), 0),
        sequence_state.last_value
    ),
    COALESCE((SELECT MAX(id) IS NOT NULL FROM portfolio_allocation_obs_time), FALSE)
        OR sequence_state.is_called
)
FROM portfolio_allocation_obs_time_id_seq AS sequence_state;

-- Authored by: OpenCode
CREATE TEMPORARY TABLE portfolio_observation_ownership_map (
    original_observation_id INTEGER NOT NULL,
    portfolio_id INTEGER NOT NULL,
    target_observation_id INTEGER NOT NULL,
    PRIMARY KEY (original_observation_id, portfolio_id),
    UNIQUE (target_observation_id)
) ON COMMIT DROP;

-- Authored by: OpenCode
INSERT INTO portfolio_observation_ownership_map (
    original_observation_id,
    portfolio_id,
    target_observation_id
)
WITH observation_portfolios AS (
    SELECT DISTINCT observation_time_id AS original_observation_id, portfolio_id
    FROM portfolio_allocation_fact
),
ranked_observation_portfolios AS (
    SELECT
        original_observation_id,
        portfolio_id,
        ROW_NUMBER() OVER (
            PARTITION BY original_observation_id
            ORDER BY portfolio_id
        ) AS portfolio_rank
    FROM observation_portfolios
)
SELECT
    original_observation_id,
    portfolio_id,
    CASE
        WHEN portfolio_rank = 1 THEN original_observation_id
        ELSE nextval('portfolio_allocation_obs_time_id_seq')::INTEGER
    END
FROM ranked_observation_portfolios;

-- Authored by: OpenCode
ALTER TABLE portfolio_allocation_fact
    DROP CONSTRAINT portfolio_allocation_obs_time_fk;

-- Authored by: OpenCode
ALTER TABLE portfolio_allocation_obs_time
    DROP CONSTRAINT portfolio_allocation_obs_time_observation_time_tag_unique;

-- Authored by: OpenCode
DROP TRIGGER trg_set_observation_time_tag ON portfolio_allocation_obs_time;

-- Authored by: OpenCode
DROP FUNCTION set_observation_time_tag();

-- Authored by: OpenCode
UPDATE portfolio_allocation_obs_time AS observation
SET portfolio_id = ownership.portfolio_id
FROM portfolio_observation_ownership_map AS ownership
WHERE ownership.original_observation_id = observation.id
  AND ownership.target_observation_id = observation.id;

-- Authored by: OpenCode
INSERT INTO portfolio_allocation_obs_time (
    id,
    portfolio_id,
    observation_time_tag,
    observation_timestamp
)
SELECT
    ownership.target_observation_id,
    ownership.portfolio_id,
    original.observation_time_tag,
    original.observation_timestamp
FROM portfolio_observation_ownership_map AS ownership
JOIN portfolio_allocation_obs_time AS original
  ON original.id = ownership.original_observation_id
WHERE ownership.target_observation_id <> ownership.original_observation_id;

-- Authored by: OpenCode
UPDATE portfolio_allocation_fact AS fact
SET observation_time_id = ownership.target_observation_id
FROM portfolio_observation_ownership_map AS ownership
WHERE fact.observation_time_id = ownership.original_observation_id
  AND fact.portfolio_id = ownership.portfolio_id;

-- Authored by: OpenCode
DELETE FROM portfolio_allocation_obs_time AS observation
WHERE NOT EXISTS (
    SELECT 1
    FROM portfolio_allocation_fact AS fact
    WHERE fact.observation_time_id = observation.id
);

-- Authored by: OpenCode
ALTER TABLE portfolio_allocation_obs_time
    ALTER COLUMN portfolio_id SET NOT NULL,
    ADD CONSTRAINT portfolio_allocation_obs_time_portfolio_fk
        FOREIGN KEY (portfolio_id) REFERENCES portfolio (id)
        ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT portfolio_allocation_obs_time_portfolio_tag_uk
        UNIQUE (portfolio_id, observation_time_tag),
    ADD CONSTRAINT portfolio_allocation_obs_time_portfolio_id_uk
        UNIQUE (portfolio_id, id);

-- Authored by: OpenCode
ALTER TABLE portfolio_allocation_fact
    ADD CONSTRAINT portfolio_allocation_fact_observation_owner_fk
        FOREIGN KEY (portfolio_id, observation_time_id)
        REFERENCES portfolio_allocation_obs_time (portfolio_id, id)
        ON DELETE RESTRICT ON UPDATE RESTRICT;

-- Authored by: OpenCode
COMMENT ON COLUMN portfolio_allocation_obs_time.portfolio_id
    IS E'Portfolio that owns this allocation observation';

-- Authored by: OpenCode
COMMENT ON TABLE portfolio_observation_ownership_map
    IS E'Temporary mapping used to split legacy shared observations by portfolio';
