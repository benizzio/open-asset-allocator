-- query to verify ingestion os Simply Wall Street with original data if needed
select a.ticker, sum(avf.asset_quantity), sum(avf.total_market_value)
from asset_value_fact avf
join asset a on a.id = avf.asset_id
where 1=1
	group by a.ticker
order by a.ticker
;

-- example query to verify if allocation plan fractal slices are correctly proportioned
select 'All Classes' as category, sum(apu.slice) as total_percentage
from allocation_plan_unit apu
where 1=1
    and apu.allocation_plan_id = 1
    and apu.structural_id[1] IS NULL
    and apu.structural_id[2] IS NOT NULL
union
select apu.structural_id[2] as category, sum(apu.slice) as total_percentage
from allocation_plan_unit apu
where 1=1
    and apu.allocation_plan_id = 1
    and apu.structural_id[1] IS NOT NULL
group by apu.structural_id[2]
;