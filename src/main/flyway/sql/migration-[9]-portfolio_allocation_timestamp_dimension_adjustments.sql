-- Update existing null observation_time_tag values with ISO string from observation_timestamp
-- Authored by: GitHub Copilot
UPDATE portfolio_allocation_obs_time
SET observation_time_tag = to_char(observation_timestamp AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"')
WHERE observation_time_tag IS NULL;

-- Add NOT NULL constraint to observation_time_tag
-- Authored by: GitHub Copilot
ALTER TABLE portfolio_allocation_obs_time
    ALTER COLUMN observation_time_tag SET NOT NULL;

-- Create a trigger function to set observation_time_tag on insert or update
-- Authored by: GitHub Copilot
CREATE OR REPLACE FUNCTION set_observation_time_tag()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.observation_time_tag IS NULL THEN
        NEW.observation_time_tag := to_char(NEW.observation_timestamp AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"');
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create a trigger to automatically set observation_time_tag when NULL
-- Authored by: GitHub Copilot
CREATE TRIGGER trg_set_observation_time_tag
BEFORE INSERT OR UPDATE ON portfolio_allocation_obs_time
FOR EACH ROW
EXECUTE FUNCTION set_observation_time_tag();

-- Update the comment for the column to reflect the changes
-- Authored by: GitHub Copilot
COMMENT ON COLUMN portfolio_allocation_obs_time.observation_time_tag
    IS 'User defined tag of the time of observation (defaults to ISO 8601 format of observation_timestamp via trg_set_observation_time_tag trigger)';
