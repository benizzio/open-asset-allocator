# TODOs

- [ ] Color configuration per hierarchy level record for visual consistency
- [ ] Pie chart for portfolio history view if possible to control inner radius
- [ ] Allocation map:
  - DIVERGENCE: last portfolio history from a timeframe, select a allocation plan to analyze
    - show divergent value in currency units
    - show divergent value in percentage
    - Endpoint
      - `GET /api/portfolio/:portfolioId/divergence/:timeFrameTag/plan/:planId`
        - returns full divergence analysis in fractal hierachy
  - CONVERGENCE: create from mapped DIVERGENCE?
    - line for each divergent record 
    - field for how much to value to move and to what asset
    - show final state after convergence (chart maybe?)