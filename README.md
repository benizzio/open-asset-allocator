# open-asset-allocator

Open source tool to manage an asset portfolio using asset allocation strategies as a first-class citizen.

> [!CAUTION]
> This project is in pre-alpha stage and is not ready for production use. It can evolve to a cloud ready service in the
> future, but is currently a work in progress without minimum security for anything but local usage.

## What is an asset allocation strategy?

> An asset allocation strategy is a plan for how to distribute your investments across different asset classes, such as
> stocks, bonds, and cash. The goal of an asset allocation strategy is to balance risk and reward by diversifying your
> investments.

Yes, this is the AI generated, statistically predominant, by the book definition. It can also be found in places like
[Investopedia](https://www.investopedia.com/terms/a/assetallocation.asp)
and [Wikipedia](https://en.wikipedia.org/wiki/Asset_allocation).
It is bland and generally not very helpful.

Some definitions I believe are a lot more useful:

- A strategy to manage assets with focus in the long term, with clear fallbacks for short term volatility and
  deviations;
- A technique to store cognitive energy in a plan that allows for intuitive, "automatic" behavior when moving assets,
  detaching it from emotional influences and deviations;
- A way to time the markets without targets, taking advantage of its cyclic nature (where appliable), and to buy "cheap"
  and sell "expensive" without having to know about the price at the time of action;
- A method that allows for higher focus on some very important risk assessment
  and [convex payoff](https://youtu.be/ovEPIQR65hc) topics:
    - Maintaining diversification with
      [uncorrelated assets](https://www.investopedia.com/articles/financial-theory/09/uncorrelated-assets-diversification.asp)
      (without forgetting correlation is fickle, as it breaks and is resstablished dynamically during cycles);
    - Exploring
      [barbell (bimodal) strategies](https://www.investopedia.com/articles/investing/013114/barbell-investment-strategy.asp);
    - Taking care of [hedging tail risks](https://youtu.be/o3Qno1rT-nw).

TODO diagram of stages of the normal asset allocation strategy

## What is this tool?

It is a web application that allows the continued management of the application of asset allocation strategies in a
fractal structure.

TODO continue