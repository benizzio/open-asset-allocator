package infra

const initialStateSQL = `
	INSERT INTO portfolio (id, "name", allocation_structure)
	VALUES (
			   1,
			   'My Portfolio Example',
			   '{"hierarchy": [{"name": "Assets", "field": "assetTicker"}, {"name": "Classes", "field": "class"}]}'::jsonb
		   )
		ON CONFLICT (id) DO UPDATE set "name" = EXCLUDED."name", allocation_structure = EXCLUDED.allocation_structure
	;

	INSERT INTO portfolio (id, "name", allocation_structure)
	VALUES (
			   2,
			   'Test Portfolio 2',
			   '{"hierarchy": [{"name": "Assets", "field": "assetTicker"}]}'::jsonb
		   )
		ON CONFLICT (id) DO UPDATE set "name" = EXCLUDED."name", allocation_structure = EXCLUDED.allocation_structure
	;
	
	-- ###################################################################
	-- ASSET TABLE
	-- ###################################################################
	INSERT INTO asset (id, ticker, "name") VALUES
		(1, 'ARCA:BIL', 'SPDR Bloomberg 1-3 Month T-Bill ETF'),
		(2, 'ARCA:STIP', 'iShares 0-5 Year TIPS Bond ETF'),
		(3, 'NasdaqGM:IEF', 'iShares 7-10 Year Treasury Bond ETF'),
		(4, 'NasdaqGM:TLT', 'iShares 20+ Year Treasury Bond ETF'),
		(5, 'NasdaqGM:SHV', 'iShares Short Treasury Bond ETF'),
		(6, 'ARCA:EWZ', 'iShares Msci Brazil ETF'),
		(7, 'ARCA:SPY', 'SPDR S&P 500 ETF Trust')
	ON CONFLICT (id) DO
		UPDATE set
			ticker = EXCLUDED.ticker,
			"name" = EXCLUDED."name"
	;
	
	-- ###################################################################
	-- ALLOCATION OBSERVATION TIME TABLE
	-- ###################################################################
	INSERT INTO portfolio_allocation_obs_time (id, observation_time_tag, observation_timestamp)
	VALUES 
	    (1, '202501', '2025-01-01 00:00:00'::TIMESTAMP),
		(2, '202503', '2025-03-01 00:00:00'::TIMESTAMP),
		(3, '202504', '2025-04-01 00:00:00'::TIMESTAMP),
		(4, '202506', '2025-06-01 00:00:00'::TIMESTAMP),
		(5, '202507', '2025-07-01 00:00:00'::TIMESTAMP) 
	;

	-- ###################################################################
	-- ALLOCATION FACT TABLE
	-- ###################################################################
	DELETE FROM portfolio_allocation_fact WHERE portfolio_id = 1;
	
	-- Portfolio 1 BONDS total market value = 27000
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
	VALUES (
			   1,
			   'BONDS',
			   FALSE,
			   100,
			   100,
			   10000,
			   1,
			   1
		   ),
		   (
			   2,
			   'BONDS',
			   FALSE,
			   80,
			   100,
			   8000,
			   1,
			   1
		   ),
		   (
			   3,
			   'BONDS',
			   FALSE,
			   60,
			   100,
			   6000,
			   1,
			   1
		   ),
		   (
			   4,
			   'BONDS',
			   FALSE,
			   30,
			   100,
			   3000,
			   1,
			   1
		   )
	;
	
	-- Portfolio 1 STOCKS total market value = 18000
	INSERT INTO public.portfolio_allocation_fact (
		asset_id,
		"class",
		cash_reserve,
		asset_quantity,
		asset_market_price,
		total_market_value,
		portfolio_id,
		create_timestamp,
		observation_time_id
	)
	VALUES (
			   5,
			   'STOCKS',
			   TRUE,
			   80,
			   100,
			   9000,
			   1,
			   now() - INTERVAL '1 minute',
			   1
		   ),
		   (
			   6,
			   'STOCKS',
			   FALSE,
			   10,
			   100,
			   1000,
			   1,
			   now() - INTERVAL '1 minute',
			   1
		   ),
		   (
			   7,
			   'STOCKS',
			   FALSE,
			   90,
			   100,
			   8000,
			   1,
			   now() - INTERVAL '1 minute',
			   1
		   )
	;
	
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
	VALUES (
			   1,
			   'BONDS',
			   FALSE,
			   100.00009,
			   100,
			   10000,
			   1,
			   2
		   )
	;

	-- ###################################################################
	-- ALLOCATION PLAN TABLE
	-- ###################################################################
	INSERT INTO allocation_plan (id, "name", "type", planned_execution_date, portfolio_id)
	VALUES (1, '60/40 Portfolio Classic - Example', 'ALLOCATION_PLAN', NULL, 1)
	;
	
	-- ###################################################################
	-- PLANNED ALLOCATION TABLE
	-- ###################################################################
	DELETE FROM planned_allocation WHERE allocation_plan_id = 1;
	
	INSERT INTO planned_allocation
	(id, allocation_plan_id, hierarchical_id, asset_id, cash_reserve, slice_size_percentage, total_market_value)
	VALUES
		(1, 1, '{NULL, "STOCKS"}', NULL, FALSE, 0.4, NULL),
		(2, 1, '{NULL, "BONDS"}', NULL, FALSE, 0.6, NULL)
	;
	
	INSERT INTO planned_allocation
	(id, allocation_plan_id, hierarchical_id, asset_id, cash_reserve, slice_size_percentage, total_market_value)
	VALUES
		(3, 1, '{"ARCA:BIL", "BONDS"}', 1, FALSE, 0.4, NULL),
		(4, 1, '{"NasdaqGM:IEF", "BONDS"}', 3, FALSE, 0.3, NULL),
		(5, 1, '{"NasdaqGM:TLT", "BONDS"}', 4, FALSE, 0.2, NULL),
		(6, 1, '{"ARCA:STIP", "BONDS"}', 2, FALSE, 0.1, NULL)
	;
	
	INSERT INTO planned_allocation
	(id, allocation_plan_id, hierarchical_id, asset_id, cash_reserve, slice_size_percentage, total_market_value)
	VALUES
		(7, 1, '{"NasdaqGM:SHV", "STOCKS"}', 5, TRUE, 0.5, NULL),
		(8, 1, '{"ARCA:EWZ", "STOCKS"}', 6, FALSE, 0.05, NULL),
		(9, 1, '{"ARCA:SPY", "STOCKS"}', 7, FALSE, 0.45, NULL)
	;

	-- ###################################################################
	-- *******************************************************************
	-- DIVERGENCE ANALYSIS SET DIFFERENCE TESTS
	-- *******************************************************************
	-- ###################################################################
	
	INSERT INTO portfolio (id, "name", allocation_structure)
	VALUES (
	    3,
	   	'Set difference test portfolio',
	   	'{"hierarchy": [{"name": "Assets", "field": "assetTicker"}, {"name": "Classes", "field": "class"}]}'::jsonb
	)
	;
	
	-- ==================================================================
	-- Lower level set difference tests
	-- ==================================================================
	
	-- Single bond and dual stock portfolio allocation
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
	VALUES (
			1, --'ARCA:BIL'
			'BONDS',
			 FALSE,
			 50.00009,
			 100,
			 5000,
			 3,
			 2
		),
		(
			6, --'ARCA:EWZ'
			'STOCKS',
			 FALSE,
			 25.00009,
			 100,
			 2500,
			 3,
			 2
		),
		(
			7, --'ARCA:SPY'
			'STOCKS',
			 FALSE,
			 24.00001,
			 100,
			 2500,
			 3,
			 2
		)
	;

	-- Dual bond and single stock portfolio allocation plan
	INSERT INTO allocation_plan (id, "name", "type", planned_execution_date, portfolio_id)
	VALUES (2, 'Dual bond single stock plan', 'ALLOCATION_PLAN', NULL, 3)
	;

	INSERT INTO planned_allocation
	(id, allocation_plan_id, hierarchical_id, asset_id, cash_reserve, slice_size_percentage, total_market_value)
	VALUES
		(10, 2, '{NULL, "BONDS"}', NULL, FALSE, 0.5, NULL),
		(11, 2, '{NULL, "STOCKS"}', NULL, FALSE, 0.5, NULL)
	;
	
	INSERT INTO planned_allocation
	(id, allocation_plan_id, hierarchical_id, asset_id, cash_reserve, slice_size_percentage, total_market_value)
	VALUES
		(12, 2, '{"ARCA:BIL", "BONDS"}', 1, FALSE, 0.4, NULL),
		(13, 2, '{"ARCA:STIP", "BONDS"}', 2, FALSE, 0.6, NULL)
	;
	
	INSERT INTO planned_allocation
	(id, allocation_plan_id, hierarchical_id, asset_id, cash_reserve, slice_size_percentage, total_market_value)
	VALUES
		(14, 2, '{"ARCA:EWZ", "STOCKS"}', 6, FALSE, 1, NULL)
	;

	-- ==================================================================
	-- Higher level set difference tests
	-- ==================================================================
	
	-- Single bond and no stock portfolio allocation
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
	VALUES (
			1, --'ARCA:BIL'
			'BONDS',
			FALSE,
			10.00009,
			100,
			10000,
			3,
			3
		)
	;

	-- Single bond and single stock portfolio allocation
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
	VALUES (
			1, --'ARCA:BIL'
			'BONDS',
			FALSE,
			40.00009,
			100,
			4000,
			3,
			4
		),
	    (
			6, --'ARCA:EWZ'
			'STOCKS',
			FALSE,
			59.00001,
			100,
			6000,
			3,
			4
		)
	;

	-- Single unplanned asset
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
	VALUES (
			1, --'ARCA:BIL'
			'A_TEST_CLASS',
			FALSE,
			10.00009,
			100,
			10000,
			3,
			5
		)
	;


	-- Single bond and single stock portfolio allocation plan
	INSERT INTO allocation_plan (id, "name", "type", planned_execution_date, portfolio_id)
	VALUES (3, 'Single bond single stock plan', 'ALLOCATION_PLAN', NULL, 3)
	;

	INSERT INTO planned_allocation
	(id, allocation_plan_id, hierarchical_id, asset_id, cash_reserve, slice_size_percentage, total_market_value)
	VALUES
		(15, 3, '{NULL, "BONDS"}', NULL, FALSE, 0.4, NULL),
		(16, 3, '{NULL, "STOCKS"}', NULL, FALSE, 0.6, NULL)
	;
	
	INSERT INTO planned_allocation
	(id, allocation_plan_id, hierarchical_id, asset_id, cash_reserve, slice_size_percentage, total_market_value)
	VALUES
		(17, 3, '{"ARCA:BIL", "BONDS"}', 1, FALSE, 1, NULL)
	;
	
	INSERT INTO planned_allocation
	(id, allocation_plan_id, hierarchical_id, asset_id, cash_reserve, slice_size_percentage, total_market_value)
	VALUES
		(18, 3, '{"ARCA:EWZ", "STOCKS"}', 6, FALSE, 1, NULL)
	;

	-- Single bond and no stock portfolio allocation plan
	INSERT INTO allocation_plan (id, "name", "type", planned_execution_date, portfolio_id)
	VALUES (4, 'Single bond no stock plan', 'ALLOCATION_PLAN', NULL, 3)
	;

	INSERT INTO planned_allocation
	(id, allocation_plan_id, hierarchical_id, asset_id, cash_reserve, slice_size_percentage, total_market_value)
	VALUES
		(19, 4, '{NULL, "BONDS"}', NULL, FALSE, 1, NULL)
	;
	
	INSERT INTO planned_allocation
	(id, allocation_plan_id, hierarchical_id, asset_id, cash_reserve, slice_size_percentage, total_market_value)
	VALUES
		(20, 4, '{"ARCA:BIL", "BONDS"}', 1, FALSE, 1, NULL)
	;

	-- Single bond for total divergence test
	INSERT INTO allocation_plan (id, "name", "type", planned_execution_date, portfolio_id)
	VALUES (5, 'Single bond total divergence test', 'ALLOCATION_PLAN', NULL, 3)
	;

	INSERT INTO planned_allocation
	(id, allocation_plan_id, hierarchical_id, asset_id, cash_reserve, slice_size_percentage, total_market_value)
	VALUES
		(21, 5, '{NULL, "BONDS"}', NULL, FALSE, 1, NULL)
	;
	
	INSERT INTO planned_allocation
	(id, allocation_plan_id, hierarchical_id, asset_id, cash_reserve, slice_size_percentage, total_market_value)
	VALUES
		(22, 5, '{"NasdaqGM:TLT", "BONDS"}', 4, FALSE, 1, NULL)
	;


	-- ###################################################################
	-- *******************************************************************
	-- SEQUENCE RESETS AFTER MANUAL ID INSERTIONS
	-- *******************************************************************
	-- ###################################################################

	SELECT setval('portfolio_id_seq', (SELECT max(id) FROM portfolio));
	SELECT setval('allocation_plan_id_seq', (SELECT max(id) FROM allocation_plan));
	SELECT setval('allocation_plan_unit_id_seq', (SELECT max(id) FROM planned_allocation));
	SELECT setval('portfolio_allocation_obs_time_id_seq', (SELECT max(id) FROM portfolio_allocation_obs_time));
	SELECT setval('asset_id_seq', (SELECT max(id) FROM asset));
`

func InitializeDBState() error {
	return ExecuteDBQuery(initialStateSQL)
}
