-- Authored by: OpenCode
ALTER TABLE allocation_plan
    ADD COLUMN planned_date DATE
;

UPDATE allocation_plan
SET planned_date = COALESCE(create_timestamp::date, CURRENT_DATE)
;

ALTER TABLE allocation_plan
    ALTER COLUMN planned_date SET NOT NULL
;

COMMENT ON COLUMN allocation_plan.planned_date IS E'Calendar date assigned to the allocation plan';
