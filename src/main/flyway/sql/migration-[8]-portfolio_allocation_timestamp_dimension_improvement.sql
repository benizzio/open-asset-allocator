CREATE TABLE portfolio_allocation_obs_time (
    id serial NOT NULL,
    observation_timestamp timestamp with time zone NOT NULL,
    observation_time_tag text,
    CONSTRAINT portfolio_allocation_obs_time_pk PRIMARY KEY (id)
)
;

COMMENT ON TABLE portfolio_allocation_obs_time IS E'Time dimension table for portfolio allocation observations';
COMMENT ON COLUMN portfolio_allocation_obs_time.observation_timestamp
        IS E'Timestamp of the observation of the portfolio allocation'
;
COMMENT ON COLUMN portfolio_allocation_obs_time.observation_time_tag IS E'User defined tag of the time of observation';

ALTER TABLE portfolio_allocation_fact
    ADD COLUMN observation_time_id integer
;

ALTER TABLE portfolio_allocation_fact ADD CONSTRAINT portfolio_allocation_obs_time_fk FOREIGN KEY (observation_time_id)
    REFERENCES portfolio_allocation_obs_time (id) MATCH FULL
    ON DELETE SET NULL ON UPDATE CASCADE
;

INSERT INTO portfolio_allocation_obs_time (observation_time_tag, observation_timestamp)
SELECT DISTINCT ON (time_frame_tag) time_frame_tag as observation_time_tag, create_timestamp as observation_timestamp
FROM portfolio_allocation_fact pa
ORDER BY time_frame_tag ASC, create_timestamp ASC
;

UPDATE portfolio_allocation_fact paf
SET observation_time_id = pot.id
FROM portfolio_allocation_obs_time pot
WHERE paf.time_frame_tag = pot.observation_time_tag
;

ALTER TABLE portfolio_allocation_fact ALTER COLUMN observation_time_id SET NOT NULL;