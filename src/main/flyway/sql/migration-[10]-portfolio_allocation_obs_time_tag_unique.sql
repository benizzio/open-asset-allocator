-- Migration: Add unique constraint to observation_time_tag in portfolio_allocation_obs_time
-- Authored by: GitHub Copilot
ALTER TABLE portfolio_allocation_obs_time
    ADD CONSTRAINT portfolio_allocation_obs_time_observation_time_tag_unique UNIQUE (observation_time_tag);
