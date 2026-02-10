-- Example portfolio data script.
-- This script is idempotent and can be run multiple times without conflicts.
-- If records with the same name already exist, a timestamp suffix is appended.
--
-- Author: Github Copilot

-- Assets (insert only if not present, comparison by ticker)
INSERT INTO asset (ticker, "name")
VALUES
    ('ARCA:BIL', 'SPDR Bloomberg 1-3 Month T-Bill ETF'),
    ('ARCA:STIP', 'iShares 0-5 Year TIPS Bond ETF'),
    ('NasdaqGM:IEF', 'iShares 7-10 Year Treasury Bond ETF'),
    ('NasdaqGM:TLT', 'iShares 20+ Year Treasury Bond ETF'),
    ('NasdaqGM:SHV', 'iShares Short Treasury Bond ETF'),
    ('ARCA:EWZ', 'iShares Msci Brazil ETF'),
    ('ARCA:SPY', 'SPDR S&P 500 ETF Trust')
ON CONFLICT (ticker) DO NOTHING
;

DO $$
DECLARE

    v_portfolio_name TEXT;
    v_portfolio_id INTEGER;
    v_obs_time_id INTEGER;
    v_plan_name TEXT;
    v_plan_id INTEGER;

BEGIN

    -- Generate unique portfolio name (append timestamp if base name exists)
    IF EXISTS (SELECT 1 FROM portfolio WHERE "name" = 'My Portfolio Example') THEN
        v_portfolio_name := 'My Portfolio Example - ' || to_char(now(), 'YYYYMMDD-HH24MISS');
    ELSE
        v_portfolio_name := 'My Portfolio Example';
    END IF;

    -- Portfolio
    INSERT INTO portfolio ("name", allocation_structure)
    VALUES (
        v_portfolio_name,
        '{"hierarchy": [{"name": "Assets", "field": "assetTicker"}, {"name": "Classes", "field": "class"}]}'::jsonb
    )
    RETURNING id INTO v_portfolio_id;

    -- Observation time
    INSERT INTO portfolio_allocation_obs_time (observation_timestamp, observation_time_tag)
    VALUES (CURRENT_DATE, to_char(CURRENT_DATE, 'YYYYMM'))
    ON CONFLICT (observation_time_tag) DO UPDATE SET
        observation_timestamp = EXCLUDED.observation_timestamp
    RETURNING id INTO v_obs_time_id;

    -- Portfolio allocation facts
    INSERT INTO portfolio_allocation_fact (
        asset_id,
        "class",
        cash_reserve,
        asset_quantity,
        asset_market_price,
        total_market_value,
        portfolio_id,
        observation_time_id
    )
    VALUES
        -- BONDS total market value = 27000
        (
            (SELECT id FROM asset WHERE ticker = 'ARCA:BIL'),
            'BONDS',
            FALSE,
            100,
            100,
            10000,
            v_portfolio_id,
            v_obs_time_id
        ),
        (
            (SELECT id FROM asset WHERE ticker = 'ARCA:STIP'),
            'BONDS',
            FALSE,
            80,
            100,
            8000,
            v_portfolio_id,
            v_obs_time_id
        ),
        (
            (SELECT id FROM asset WHERE ticker = 'NasdaqGM:IEF'),
            'BONDS',
            FALSE,
            60,
            100,
            6000,
            v_portfolio_id,
            v_obs_time_id
        ),
        (
            (SELECT id FROM asset WHERE ticker = 'NasdaqGM:TLT'),
            'BONDS',
            FALSE,
            30,
            100,
            3000,
            v_portfolio_id,
            v_obs_time_id
        ),
        -- STOCKS total market value = 18000
        (
            (SELECT id FROM asset WHERE ticker = 'NasdaqGM:SHV'),
            'STOCKS',
            TRUE,
            80,
            100,
            9000,
            v_portfolio_id,
            v_obs_time_id
        ),
        (
            (SELECT id FROM asset WHERE ticker = 'ARCA:EWZ'),
            'STOCKS',
            FALSE,
            10,
            100,
            1000,
            v_portfolio_id,
            v_obs_time_id
        ),
        (
            (SELECT id FROM asset WHERE ticker = 'ARCA:SPY'),
            'STOCKS',
            FALSE,
            90,
            100,
            8000,
            v_portfolio_id,
            v_obs_time_id
        )
    ;

    -- Generate unique allocation plan name (append timestamp if base name exists)
    IF EXISTS (SELECT 1 FROM allocation_plan WHERE "name" = '60/40 Portfolio Classic - Example') THEN
        v_plan_name := '60/40 Portfolio Classic - Example - ' || to_char(now(), 'YYYYMMDD-HH24MISS');
    ELSE
        v_plan_name := '60/40 Portfolio Classic - Example';
    END IF;

    -- Allocation plan
    INSERT INTO allocation_plan ("name", "type", planned_execution_date, portfolio_id)
    VALUES (
        v_plan_name,
        'ALLOCATION_PLAN',
        NULL,
        v_portfolio_id
    )
    RETURNING id INTO v_plan_id;

    -- Planned allocations
    INSERT INTO planned_allocation (
        allocation_plan_id,
        hierarchical_id,
        asset_id,
        cash_reserve,
        slice_size_percentage,
        total_market_value
    )
    VALUES
        -- Class level allocations
        (
            v_plan_id,
            '{NULL, "STOCKS"}'::text[],
            NULL,
            FALSE,
            0.4,
            NULL
        ),
        (
            v_plan_id,
            '{NULL, "BONDS"}'::text[],
            NULL,
            FALSE,
            0.6,
            NULL
        ),
        -- BONDS asset allocations
        (
            v_plan_id,
            '{"ARCA:BIL", "BONDS"}'::text[],
            (SELECT id FROM asset WHERE ticker = 'ARCA:BIL'),
            FALSE,
            0.4,
            NULL
        ),
        (
            v_plan_id,
            '{"NasdaqGM:IEF", "BONDS"}'::text[],
            (SELECT id FROM asset WHERE ticker = 'NasdaqGM:IEF'),
            FALSE,
            0.3,
            NULL
        ),
        (
            v_plan_id,
            '{"NasdaqGM:TLT", "BONDS"}'::text[],
            (SELECT id FROM asset WHERE ticker = 'NasdaqGM:TLT'),
            FALSE,
            0.2,
            NULL
        ),
        (
            v_plan_id,
            '{"ARCA:STIP", "BONDS"}'::text[],
            (SELECT id FROM asset WHERE ticker = 'ARCA:STIP'),
            FALSE,
            0.1,
            NULL
        ),
        -- STOCKS asset allocations
        (
            v_plan_id,
            '{"NasdaqGM:SHV", "STOCKS"}'::text[],
            (SELECT id FROM asset WHERE ticker = 'NasdaqGM:SHV'),
            TRUE,
            0.5,
            NULL
        ),
        (
            v_plan_id,
            '{"ARCA:EWZ", "STOCKS"}'::text[],
            (SELECT id FROM asset WHERE ticker = 'ARCA:EWZ'),
            FALSE,
            0.05,
            NULL
        ),
        (
            v_plan_id,
            '{"ARCA:SPY", "STOCKS"}'::text[],
            (SELECT id FROM asset WHERE ticker = 'ARCA:SPY'),
            FALSE,
            0.45,
            NULL
        )
    ;

END $$;

