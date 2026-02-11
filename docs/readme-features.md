# Features

## Manage multiple portfolios

![portfolios](/docs/images/portfolio-selection.png)

- Each portfolio has it own customizable hierarchical structure that must follow a "fractal" pattern (the "hierarchy")
  (**pre-alpha**: fixed structure with 2 levels, class and asset);
- Inside each portfolio, the assets are classified with data properties that can be chosen as levels (as many as
  needed) of the hierarchy and group them (**pre-alpha**: fixed structure with 2 levels, class and asset);
- Each level of the hierarchy has their own set proportions, and can be monitored and balanced as such;
- In the lowest level of the hierarchy, the assets, there is a classifier for "cash" reserves (assets that can be
  used for risk mitigation;
  and [optionality](https://www.notion.so/Asset-allocation-applied-to-an-Antifragility-theoretical-basis-235acfc60ff080d08fabe96819a6e680?source=copy_link#25facfc60ff08010b14bcda99230b22c)
  if the allocation phase demands it);

## Historical data and visualization

![detail class level](/docs/images/portfolio-history-detail1.png)
![detail asset level](/docs/images/portfolio-history-detail2.png)
![history management](/docs/images/portfolio-history-management.png)

- At any planned (or unplanned) interval a snapshot of the portfolio can be stored to be visualized and used for
  re-balacing (reallocation) analysis;
- Historical data can be manually edited in management screen.

## Allocation planning

![detail class level](/docs/images/allocation-plan-detail.png)
![plan management](/docs/images/allocation-plan-management.png)

- Allocation plans following the portfolio structure can be created and used for re-balancing analysis;
- As time passes, monitoring intervals are reached or market conditions change, new allocation plans can be created
  and old plans are stored for historical analysis;
- Plans can be manually edited in management screen.

## Divergence analysis

![divergence analysis](/docs/images/divergence-analysis.png)

- Any historical snapshot of the (allocation of the) portfolio can be conbined to any plan of the same portfolio to
  create a Divergence analysis, illustrating how out of balance the portfolio is from the selected plan;
- Multiple indicators to signal size of divergence and help with re-balancing decisions.

## Convergence analysis and planning

- **pre-alpha**: Planned as a future feature;