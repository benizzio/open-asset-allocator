-- Migration: Add field size constraints to text fields
-- This migration adds VARCHAR constraints to text fields to enforce size limits
-- that will be synchronized with front-end and back-end validation

-- Asset table: ticker and name constraints
ALTER TABLE asset
    ALTER COLUMN ticker SET DATA TYPE varchar(20),
    ALTER COLUMN name SET DATA TYPE varchar(100);

-- Portfolio table: name constraint
ALTER TABLE portfolio
    ALTER COLUMN name SET DATA TYPE varchar(100);

-- Allocation plan table: name and type constraints
ALTER TABLE allocation_plan
    ALTER COLUMN name SET DATA TYPE varchar(100),
    ALTER COLUMN type SET DATA TYPE varchar(50);

-- Portfolio allocation fact table: class constraint
ALTER TABLE portfolio_allocation_fact
    ALTER COLUMN class SET DATA TYPE varchar(100);

-- Portfolio allocation observation time table: observation_time_tag constraint
ALTER TABLE portfolio_allocation_obs_time
    ALTER COLUMN observation_time_tag SET DATA TYPE varchar(100);
