# open-asset-allocator

Open source tool to manage an asset portfolio using asset allocation strategies as a first-class citizen.

> [!CAUTION]
> This project is in pre-alpha stage and is not ready for production use. It can evolve to a cloud ready service in the
> future, but is currently a work in progress without minimum security for anything but local usage.

## What is an asset allocation strategy?

> An asset allocation strategy is a plan for how to distribute your investments across different asset classes, such as
> stocks, bonds, and cash. The goal of an asset allocation strategy is to balance risk and reward by diversifying your
> investments.

Yes, this is the AI generated, statistically predominant, by the book definition. Something similar can also be found in
places like [Investopedia](https://www.investopedia.com/terms/a/assetallocation.asp)
and [Wikipedia](https://en.wikipedia.org/wiki/Asset_allocation).
It is bland and generally not very helpful.

Some definitions I believe are a lot more useful:

- A strategy to manage assets with focus in the long term, with clear fallbacks for short term volatility and
  deviations;
- A technique to store cognitive energy in a plan that allows for intuitive, "automatic" behavior when moving assets,
  helping detach it from emotional influences and deviations;
- A way to time the markets without targets, taking advantage of its cyclic nature (where appliable), and to buy "cheap"
  and sell "expensive" without having to know about the price at the time of action;
- A process that is used to manage large ammounts of assets, enjoying large accumulation zones and permanent time in the
  market;
- A method that allows for higher focus on some very important risk assessment
  and [convex payoff](https://youtu.be/ovEPIQR65hc) topics:
    - Maintaining diversification with
      [uncorrelated assets](https://www.investopedia.com/articles/financial-theory/09/uncorrelated-assets-diversification.asp)
      (without forgetting correlation is fickle, as it breaks and is resstablished dynamically during cycles);
    - Exploring
      [barbell (bimodal) strategies](https://www.investopedia.com/articles/investing/013114/barbell-investment-strategy.asp);
    - Taking care of [hedging tail risks](https://youtu.be/o3Qno1rT-nw).

![asset allocation flow](/docs/images/asset-allocation-flow.png)

## What is this tool?

It is an application that allows the continued management of the usage of asset allocation strategies for long term
portfolios in a "fractal" structure. The allocation needs to follow this bottom-up, repeatable structure, composed of
assets on the lowest level.
Those assets can be grouped by any criteria defined by the implemented model, and that will form the top levels of the
hierarchical structure.

Example portfolio (concrete 60/40 classic portfolio) with 2 layers hierarchy (CLASSES and ASSETS, this portfolio is
available as an example dataset in the project):

> [!CAUTION]
> The examples shown in this document or any part of the project are for mere illustration purposes of the application
> features with real data. They are not, at any point in time, to be considered investment advice. Asset values are also
> fictional.

![class level](/docs/images/example-portfolio-classes.png)
![asset level bonds](/docs/images/example-portfolio-bonds.png)
![asset level stocks](/docs/images/example-portfolio-stocks.png)

### Features

- **Manage multiple portfolios**:
    - Each portfolio has it own customizable hierarchical structure that must follow a "fractal" pattern
      (the "hierarchy");
    - Inside each portfolio, the assets are classified with data properties that can be chosen as levels (as many as
      needed) of the hierarchy and group them;
    - Each level of the hierarchy has their own set proportions, and can be monitored and balanced as such;
    - In the lowest level of the hierarchy, the assets, there is a classifier for "cash" reserves (assets that can be
      used for risk mitigation
      and [optionality](https://medium.com/@hannes.rollin/antifragile-system-design-1-optionality-17b60fa0842d) if the
      allocation phase demands it);
- **Historical data and visualization**:
    - At any planned (or unplanned) interval a snapshot of the portfolio can be stored to be visualized and used for
      re-balacing (reallocation) analysis;

![portfolio page](/docs/images/portfolio-page.png)

- **Allocation planning**:
    - Allocation plans following the portfolio structure can be created and used for re-balancing analysis;
    - As time passes, monitoring intervals are reached or maret conditions change, new allocation plans can be created
      and old plans are stored for
      historical analysis;

![allocation plan page](/docs/images/allocation-plan-page.png)

- **Divergence analysis**:
    - Any historical snapshot of the (allocation of the) portfolio can be conbined to any plan of the same portfolio to
      create a Divergence analysis, illustrating how out of balance the portfolio is from the selected plan;

![divergence analysis page](/docs/images/divergence-analysis-page.png)

- **Convergence analysis and planning**:
    - Planned as a future feature;

### Roadmap(ish)

- [ ] Color configuration per hierarchy level record for visual consistency
- [x] Pie chart for portfolio history view if possible to control inner radius
- [ ] Allocation map:
    - [x] DIVERGENCE: last portfolio history from a timeframe, select an allocation plan to analyze
        - [ ] persistence of generated analysis on timeframe for convenience
        - [x] show divergent value in currency units
        - [x] show divergent value in percentage
        - [ ] allow easy UI elements to add extenal cash inflow
    - [ ] CONVERGENCE: create from mapped DIVERGENCE?
        - bound to a divergence analysis and its plan
        - line for each divergent record
        - field for how much to value to move and to what asset
        - show final state after convergence (chart maybe?)