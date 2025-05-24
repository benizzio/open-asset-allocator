-- query to verify ingestion os Simply Wall Street with original data if needed
select a.ticker, sum(avf.asset_quantity), sum(avf.total_market_value)
from asset_value_fact avf
join asset a on a.id = avf.asset_id
where 1=1
	group by a.ticker
order by a.ticker
;

-- example query to verify if allocation plan fractal slices are correctly proportioned
select 'All Classes' as category, sum(pa.slice_size_percentage) as total_percentage
from planned_allocation pa
where 1=1
  and pa.allocation_plan_id = 1
  and pa.structural_id[1] IS NULL
  and pa.structural_id[2] IS NOT NULL
union
select pa.structural_id[2] as category, sum(pa.slice_size_percentage) as total_percentage
from planned_allocation pa
where 1=1
  and pa.allocation_plan_id = 1
  and pa.structural_id[1] IS NOT NULL
group by pa.structural_id[2]
;