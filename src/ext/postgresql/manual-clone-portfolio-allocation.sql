-- Clone into new observation of same portfolio
INSERT INTO portfolio_allocation_fact (
    asset_id,
    class,
    cash_reserve,
    asset_quantity,
    asset_market_price,
    total_market_value,
    portfolio_id,
    observation_time_id
)
SELECT
    paf.asset_id,
    paf.class,
    paf.cash_reserve,
    paf.asset_quantity,
    paf.asset_market_price,
    paf.total_market_value,
    paf.portfolio_id,
    (SELECT id FROM portfolio_allocation_obs_time WHERE observation_time_tag = '202509') as observation_time_id
FROM portfolio_allocation_fact paf
JOIN portfolio_allocation_obs_time paot ON paf.observation_time_id = paot.id
WHERE paot.observation_time_tag = '202507' AND paf.portfolio_id = 1
;

-- Check cloning after
SELECT *
FROM portfolio_allocation_fact paf
WHERE paf.observation_time_id IN (
    SELECT id FROM portfolio_allocation_obs_time WHERE observation_time_tag IN ('202509', '202507')
)
ORDER BY asset_id, observation_time_id ASC;