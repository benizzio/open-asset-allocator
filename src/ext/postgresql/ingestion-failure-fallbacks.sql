-- Manually inserting portfolio allocation fact data
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
SELECT
    asset_id,
    "class",
    cash_reserve,
    asset_quantity,
    asset_market_price,
    asset_quantity * asset_market_price AS total_market_value,
    time_frame_tag,
    portfolio_id
FROM (
    VALUES (
       6,
       'STORE_VALUE',
       false,
       1,
       1,
       '202505',
       1
    )
) AS t(
       asset_id,
       "class",
       cash_reserve,
       asset_quantity,
       asset_market_price,
       time_frame_tag,
       portfolio_id
);