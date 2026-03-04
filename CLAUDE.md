# Accrue Engine

Open-source fintech engine. Companies deploy it to offer trading, portfolio management, and cash accounts under their own brand. Think Supabase for brokerage infrastructure.

## What This Repo Is

The core backend. Trade routing, portfolio engine, cash ledger, broker adapter interface. Everything is orchestration — Accrue never holds funds, executes trades directly, or handles compliance.

## Tech Stack

- **Language:** Go (strictly — no second language)
- **Database:** PostgreSQL (one per service, or shared with schema isolation)
- **Inter-service:** gRPC + protobuf
- **External API:** REST with OpenAPI spec (gateway only)
- **Deployment:** Docker + docker-compose

## Architecture — Microservices

```
                    ┌──────────────────┐
    External ──────▶│   API Gateway    │ (REST + OpenAPI)
    Clients         │   (auth, routing)│
                    └──┬───┬───┬───────┘
                       │   │   │  gRPC
              ┌────────┘   │   └────────┐
              ▼            ▼            ▼
     ┌────────────┐ ┌───────────┐ ┌──────────┐
     │  Adapter   │ │ Portfolio │ │  Ledger  │
     │  Service   │ │  Service  │ │  Service │
     │            │ │           │ │          │
     │ ┌────────┐ │ │ positions │ │ double-  │
     │ │  Mock  │ │ │ P&L      │ │ entry    │
     │ │ Alpaca │ │ │ tracking │ │ balances │
     │ │ IBKR   │ │ │          │ │ holds    │
     │ └────────┘ │ │           │ │          │
     └────────────┘ └───────────┘ └──────────┘
           │
           ▼
     External Brokers
```

### Services

| Service | Responsibility | Port (dev) |
|---------|---------------|------------|
| **api-gateway** | External REST API, auth (API keys), request routing, rate limiting | 8080 |
| **adapter-service** | Broker adapter interface, order management, trade execution | 9001 |
| **portfolio-service** | Positions, P&L calculation, multi-asset tracking, performance | 9002 |
| **ledger-service** | Double-entry cash ledger, balances, credits/debits, holds | 9003 |

### Communication

- **External → Gateway:** REST (JSON)
- **Gateway → Services:** gRPC
- **Service → Service:** gRPC (e.g., adapter-service calls ledger-service to debit on trade fill)
- **Async events:** Start with synchronous gRPC calls. Add a message bus (NATS) later only when needed.

## Phase 1 Scope (Current)

Build these in order:

1. **Project scaffolding** — Go workspace, service directories, shared proto definitions, Makefile, docker-compose.yml
2. **Proto definitions** — define gRPC contracts between services (orders, positions, balances, accounts)
3. **Adapter service** — broker adapter interface + mock adapter (paper trading, in-memory state)
4. **Ledger service** — double-entry bookkeeping, balance tracking, holds
5. **Portfolio service** — multi-asset positions, P&L calculation
6. **API gateway** — REST handlers, auth middleware, routes to services via gRPC
7. **Docker deployment** — `docker-compose up` starts all services + postgres

## Project Structure

```
accrue-engine/
├── proto/                       # shared protobuf definitions
│   ├── adapter/v1/
│   │   └── adapter.proto
│   ├── portfolio/v1/
│   │   └── portfolio.proto
│   ├── ledger/v1/
│   │   └── ledger.proto
│   └── common/v1/
│       └── types.proto          # shared types (Money, Asset, etc.)
│
├── services/
│   ├── gateway/                 # API gateway service
│   │   ├── cmd/
│   │   │   └── main.go
│   │   ├── internal/
│   │   │   ├── handler/         # REST handlers
│   │   │   ├── middleware/      # auth, rate limiting
│   │   │   └── router.go
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── adapter/                 # adapter service
│   │   ├── cmd/
│   │   │   └── main.go
│   │   ├── internal/
│   │   │   ├── broker/          # broker interface + implementations
│   │   │   │   ├── broker.go    # interface definition
│   │   │   │   └── mock/        # mock adapter
│   │   │   ├── order/           # order management
│   │   │   └── server/          # gRPC server implementation
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── portfolio/               # portfolio service
│   │   ├── cmd/
│   │   │   └── main.go
│   │   ├── internal/
│   │   │   ├── position/        # position tracking
│   │   │   ├── pnl/             # P&L calculation
│   │   │   ├── store/           # PostgreSQL layer
│   │   │   └── server/          # gRPC server implementation
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   └── ledger/                  # ledger service
│       ├── cmd/
│       │   └── main.go
│       ├── internal/
│       │   ├── entry/           # double-entry logic
│       │   ├── balance/         # balance calculations
│       │   ├── store/           # PostgreSQL layer
│       │   └── server/          # gRPC server implementation
│       ├── Dockerfile
│       └── go.mod
│
├── pkg/                         # shared Go packages (minimal)
│   └── id/                      # UUID generation helpers
│
├── migrations/
│   ├── ledger/
│   ├── portfolio/
│   └── adapter/
│
├── docker-compose.yml           # all services + postgres
├── Makefile                     # build, test, proto-gen, docker targets
├── go.work                      # Go workspace file linking all modules
├── buf.yaml                     # protobuf tooling config
├── README.md
└── CLAUDE.md
```

## Go Conventions

- Each service is its own Go module (separate `go.mod`)
- Use Go workspaces (`go.work`) for local development across services
- `internal/` per service — services don't import each other's internal packages
- `pkg/` for genuinely shared utilities only (keep this minimal)
- Standard library first. Only add dependencies when they earn their place.
- Errors are values. Wrap with context: `fmt.Errorf("placing order: %w", err)`
- No frameworks for HTTP — use `net/http` + chi for the gateway
- Table-driven tests. Test files next to the code they test.
- No global state. Dependencies injected via constructors.

## Protobuf Conventions

- Use [buf](https://buf.build/) for linting, breaking change detection, and code generation
- Proto files in `proto/<service>/v1/`
- Version all proto packages (v1, v2)
- Generated Go code goes into each service's internal `gen/` directory
- Keep proto definitions minimal — don't leak implementation details into contracts

## Database Conventions

- Each service owns its data. No cross-service database queries.
- Use SQL migrations (golang-migrate)
- Migrations per service in `migrations/<service>/`
- Use transactions for anything touching money (ledger entries, order state changes)
- UUIDs for primary keys
- `created_at` and `updated_at` on every table
- snake_case for column names
- No ORMs — write SQL, use pgx directly

## Key Design Decisions

- **Microservices from day one.** Clean service boundaries: gateway, adapter, portfolio, ledger. Each deployable independently.
- **gRPC between services.** Typed contracts via protobuf. REST only at the gateway for external consumers.
- **Adapters are plugins.** The broker interface in the adapter service is the most important abstraction. Get this right first.
- **Double-entry ledger.** Every money movement creates two entries (debit + credit). No single-entry shortcuts.
- **Multi-tenant from day one.** Every resource belongs to an account. Even if there's only one tenant now, the schema should support many.
- **Service isolation.** Each service has its own database schema, its own migrations, its own Go module. No shared database access.

## Building Blocks to Study

- **[Blnk](https://github.com/blnkfinance/blnk)** — double-entry ledger (cash account patterns)
- **[Moov](https://github.com/moov-io)** — ACH, wire transfer, OFAC (money movement primitives)
- **[GoCryptoTrader](https://github.com/thrasher-corp/gocryptotrader)** — exchange connector abstraction
- **[Open Trading Platform](https://github.com/ettec/open-trading-platform)** — order management architecture (also microservices + gRPC)

## Git

- **Author:** `cnylum <cnylum@users.noreply.github.com>` (local config already set)
- **Remote:** `git@github-cnylum:cnylum/Accrue-engine.git`
- **Branch:** `main`
- Commit messages: imperative mood, concise. "Add broker adapter interface", not "Added broker adapter interface"

## Related Repos

- **[accrue-business](https://github.com/cnylum/accrue-business)** — strategy, roadmap, research, content
- **accrue-ui** (future) — reference React frontend
- **accrue-sdk** (future) — client SDKs

## What Not To Do

- Don't add compliance features — that's the deployer's problem
- Don't build a consumer-facing app — this is infrastructure
- Don't add dependencies without a clear reason
- Don't use ORMs — write SQL, use pgx directly
- Don't share databases between services — each service owns its data
- Don't add a message bus yet — start with synchronous gRPC, add async when there's a real need
