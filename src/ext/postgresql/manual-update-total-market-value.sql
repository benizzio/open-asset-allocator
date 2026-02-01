-- Manual script to update total_market_value based on asset_quantity * asset_market_price
-- for a specific portfolio and observation time

-- Update total_market_value for a specific portfolio_id and observation_time_id
UPDATE portfolio_allocation_fact
SET total_market_value = asset_quantity * asset_market_price
WHERE portfolio_id = 1
  AND observation_time_id = (SELECT id FROM portfolio_allocation_obs_time WHERE observation_time_tag = '202602')
  AND "class" = 'CRYPTO'
;

-- Verify the update
SELECT paf."class",
       paf.asset_id,
       paf.asset_quantity,
       paf.asset_market_price,
       paf.total_market_value,
       (paf.asset_quantity * paf.asset_market_price) AS calculated_value
FROM portfolio_allocation_fact paf
WHERE paf.portfolio_id = 1
  AND paf.observation_time_id = (SELECT id FROM portfolio_allocation_obs_time WHERE observation_time_tag = '202602')
ORDER BY paf.asset_id ASC
;