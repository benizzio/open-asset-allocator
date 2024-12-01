ALTER TABLE allocation_plan
    ALTER COLUMN structure SET DATA TYPE JSONB USING structure::jsonb,
    ADD COLUMN IF NOT EXISTS create_timestamp TIMESTAMP DEFAULT now()
;

COMMENT ON COLUMN allocation_plan.structure IS E'Definition of the structure of the asset allocation plan in hierarchical levels';

ALTER TABLE allocation_plan_unit
    ALTER COLUMN structural_id SET DATA TYPE text[] USING structural_id::text[],
    ALTER COLUMN slice SET DATA TYPE NUMERIC(5, 2)
;

ALTER TABLE allocation_plan_unit
    RENAME COLUMN slice TO slice_size_percentage
;

COMMENT ON COLUMN allocation_plan_unit.structural_id
        IS E'Identifier of the plan unit structure inside the hierarchical classification level';

COMMENT ON COLUMN allocation_plan_unit.slice_size_percentage
        IS E'Allocation slice of the planned asset, bounded to the lower classifier granularity (in %) inside the higher granularity, for ALLOCATION_PLANs';

ALTER TABLE public.allocation_plan_unit
    DROP CONSTRAINT slice_percentage_ck,
    ADD CONSTRAINT slice_percentage_ck CHECK (((slice_size_percentage >= (0)::numeric) AND (slice_size_percentage <= (100)::numeric)))
;