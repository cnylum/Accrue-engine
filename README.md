# Accrue Engine

Open-source fintech engine for building trading, portfolio management, and cash account products.

Companies deploy Accrue under their own brand and licenses. Plug in your broker, your compliance, your frontend. Accrue handles the orchestration.

> Think Supabase, but for brokerage infrastructure.

## What It Does

- **Trade routing** — pluggable broker adapter interface. Connect any broker or exchange.
- **Portfolio engine** — multi-asset positions (stocks, crypto, ETFs), P&L, performance tracking.
- **Cash ledger** — double-entry balance tracking for cash accounts.
- **REST API** — OpenAPI-documented. Build any frontend on top.

## Quick Start

```bash
docker-compose up
```

> Not yet available — engine is under active development.

## Architecture

```
┌─────────────────────────────┐
│        REST API              │
├─────────────────────────────┤
│  Portfolio Engine  │  Ledger │
├─────────────────────────────┤
│    Broker Adapter Interface  │
├──────┬──────┬──────┬────────┤
│ Mock │Alpaca│ IBKR │ ...    │
└──────┴──────┴──────┴────────┘
         ▼
    External Brokers
```

Adapters are plugins. Accrue ships with a mock adapter for development. Real broker adapters (Alpaca, Interactive Brokers, crypto exchanges) are added incrementally.

## Tech Stack

- **Language:** Go
- **Database:** PostgreSQL
- **Deployment:** Docker

## Project Status

**Phase 1 — Engine Core** (in progress)

- [ ] Project scaffolding
- [ ] Broker adapter interface
- [ ] Mock adapter (paper trading)
- [ ] Portfolio engine
- [ ] Cash ledger
- [ ] REST API
- [ ] Auth / API keys
- [ ] Docker deployment

See the full [roadmap](https://github.com/cnylum/accrue-business/blob/main/ROADMAP.md).

## What Accrue Is Not

- **Not a consumer app** — the reference UI is a demo, not the product.
- **Not a brokerage** — Accrue does not hold funds, execute trades directly, or handle compliance. That's your responsibility.
- **Not financial advice** — it's infrastructure.

## License

MIT

## Links

- [Roadmap & Strategy](https://github.com/cnylum/accrue-business)
