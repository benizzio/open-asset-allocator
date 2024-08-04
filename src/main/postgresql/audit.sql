-- query to verify ingestion os Simply Wall Street with original data if needed
select a.ticker, sum(avf.asset_quantity), sum(avf.total_market_value)
from asset_value_fact avf
join asset a on a.id = avf.asset_id
where 1=1
	group by a.ticker
order by a.ticker
;
