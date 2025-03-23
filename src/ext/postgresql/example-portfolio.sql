INSERT INTO portfolio (id, "name", allocation_structure)
VALUES (
           1,
           'My Portfolio Example',
           '{"hierarchy": [{"name": "Assets", "field": "assetTicker"}, {"name": "Classes", "field": "class"}]}'::jsonb
       )
    ON CONFLICT (id) DO UPDATE SET "name" = EXCLUDED."name", allocation_structure = EXCLUDED.allocation_structure
;

SELECT setval('portfolio_id_seq', (SELECT max(id) FROM portfolio));

INSERT INTO asset (id, ticker, "name") VALUES
    (1, 'ARCA:BIL', 'SPDR Bloomberg 1-3 Month T-Bill ETF'),
    (2, 'ARCA:STIP', 'iShares 0-5 Year TIPS Bond ETF'),
    (3, 'NasdaqGM:IEF', 'iShares 7-10 Year Treasury Bond ETF'),
    (4, 'NasdaqGM:TLT', 'iShares 20+ Year Treasury Bond ETF'),
    (5, 'NasdaqGM:SHV', 'iShares Short Treasury Bond ETF'),
    (6, 'ARCA:EWZ', 'iShares Msci Brazil ETF'),
    (7, 'ARCA:SPY', 'SPDR S&P 500 ETF Trust')
ON CONFLICT (id) DO
    UPDATE SET
        ticker = EXCLUDED.ticker,
        "name" = EXCLUDED."name"
;

SELECT setval('asset_id_seq', (SELECT max(id) FROM asset));

DELETE FROM portfolio_allocation_fact WHERE portfolio_id = 1;

-- BONDS total market value = 27000
INSERT INTO portfolio_allocation_fact (
    asset_id,
    "class",
    cash_reserve,
    asset_quantity,
    asset_market_price,
    total_market_value,
    time_frame_tag,
    portfolio_id
)
VALUES (
           1,
           'BONDS',
           FALSE,
           100,
           100,
           10000,
           to_char(CURRENT_DATE, 'YYYYMM'),
           1
       ),
       (
           2,
           'BONDS',
           FALSE,
           80,
           100,
           8000,
           to_char(CURRENT_DATE, 'YYYYMM'),
           1
       ),
       (
           3,
           'BONDS',
           FALSE,
           60,
           100,
           6000,
           to_char(CURRENT_DATE, 'YYYYMM'),
           1
       ),
       (
           4,
           'BONDS',
           FALSE,
           30,
           100,
           3000,
           to_char(CURRENT_DATE, 'YYYYMM'),
           1
       )
;

-- STOCKS total market value = 18000
INSERT INTO public.portfolio_allocation_fact (
    asset_id,
    "class",
    cash_reserve,
    asset_quantity,
    asset_market_price,
    total_market_value,
    time_frame_tag,
    portfolio_id
)
VALUES (
           5,
           'STOCKS',
           TRUE,
           80,
           100,
           9000,
           to_char(CURRENT_DATE, 'YYYYMM'),
           1
       ),
       (
           6,
           'STOCKS',
           FALSE,
           10,
           100,
           1000,
           to_char(CURRENT_DATE, 'YYYYMM'),
           1
       ),
       (
           7,
           'STOCKS',
           FALSE,
           90,
           100,
           8000,
           to_char(CURRENT_DATE, 'YYYYMM'),
           1
       )
;

INSERT INTO allocation_plan (id, "name", "type", planned_execution_date, portfolio_id)
VALUES (
           1,
           '60/40 Portfolio Classic - Example',
           'ALLOCATION_PLAN',
           NULL,
           1
       )
    ON CONFLICT (id) DO
UPDATE SET
    "name" = EXCLUDED."name",
    "type" = EXCLUDED."type",
    planned_execution_date = EXCLUDED.planned_execution_date,
    portfolio_id = EXCLUDED.portfolio_id
;

DELETE FROM planned_allocation WHERE allocation_plan_id = 1;

INSERT INTO planned_allocation
(allocation_plan_id, structural_id, asset_id, cash_reserve, slice_size_percentage, total_market_value)
VALUES
    (1, '{NULL, "STOCKS"}', NULL, false, 0.4, NULL),
    (1, '{NULL, "BONDS"}', NULL, false, 0.6, NULL)
;

INSERT INTO planned_allocation
(allocation_plan_id, structural_id, asset_id, cash_reserve, slice_size_percentage, total_market_value)
VALUES
    (1, '{"ARCA:BIL", "BONDS"}', 1, false, 0.4, NULL),
    (1, '{"NasdaqGM:IEF", "BONDS"}', 3, false, 0.3, NULL),
    (1, '{"NasdaqGM:TLT", "BONDS"}', 4, false, 0.2, NULL),
    (1, '{"ARCA:STIP", "BONDS"}', 2, false, 0.1, NULL)
;

INSERT INTO planned_allocation
(allocation_plan_id, structural_id, asset_id, cash_reserve, slice_size_percentage, total_market_value)
VALUES
    (1, '{"NasdaqGM:SHV", "STOCKS"}', 5, true, 0.5, NULL),
    (1, '{"ARCA:EWZ", "STOCKS"}', 6, false, 0.05, NULL),
    (1, '{"ARCA:SPY", "STOCKS"}', 7, false, 0.45, NULL)
;

SELECT setval('allocation_plan_id_seq', (SELECT max(id) FROM allocation_plan));