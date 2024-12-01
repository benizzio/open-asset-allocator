ALTER TABLE allocation_plan_unit
    ALTER COLUMN slice_size_percentage SET DATA TYPE NUMERIC(10, 5)
;

UPDATE allocation_plan_unit SET slice_size_percentage = slice_size_percentage / 100;

ALTER TABLE allocation_plan_unit
    ALTER COLUMN slice_size_percentage SET DATA TYPE NUMERIC(6, 5)
;

ALTER TABLE public.allocation_plan_unit
    DROP CONSTRAINT slice_percentage_ck,
    ADD CONSTRAINT slice_percentage_ck CHECK (((slice_size_percentage >= (0)::numeric) AND (slice_size_percentage <= (1)::numeric)))
;