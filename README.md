# TicksAndTomes

A browser-based empire building & PvP game featuring skill progression, resource management, tactical combat, clan warfare, and player-driven markets.

**Status**: Early development (Phase 1: Core Rules + Game Loop)

## Quick Start

### Prerequisites
- **Go 1.21+**
- **Node.js 18+** (npm 9+)
- **Docker** & **Docker Compose** (optional, for containerized setup)

### Option 1: Local Development (Recommended)

```bash
# Clone and enter project
cd /path/to/TicksAndTomes

# Install dependencies
make setup

# Copy environment template
cp .env.example .env

# Start background + frontend in separate terminals
# Terminal 1: Backend
make backend

# Terminal 2: Frontend
make frontend
```

Then visit **http://localhost:3000** 🎮

### Option 2: Docker (Containerized)

```bash
# Start all services (backend, frontend, MongoDB)
make docker-up

# View logs
docker-compose logs -f

# Stop services
make docker-down
```

Visit **http://localhost:3000** after services initialize (~10s).

## Project Structure

```
├── backend/                    # Go backend (Gin + MongoDB)
│   ├── cmd/server/            # Server entry point
│   ├── internal/              # Business logic & handlers
│   ├── pkg/                   # Shared packages
│   ├── scripts/               # Database init, migrations
│   ├── go.mod                 # Go dependencies
│   └── Dockerfile             # Container definition
│
├── frontend/                   # React + TypeScript frontend (Vite)
│   ├── src/                   # React components, hooks, types
│   │   ├── components/        # UI components
│   │   ├── hooks/             # Custom hooks
│   │   ├── types/             # TypeScript interfaces
│   │   ├── services/          # API clients
│   │   └── styles/            # CSS & Tailwind config
│   ├── package.json           # npm dependencies
│   └── Dockerfile             # Container definition
│
├── docs/                      # Game design & architecture
│   ├── ENGINE_CONTRACT.md     # Core game mechanics
│   ├── COMPONENT_STRUCTURE.md # Frontend architecture
│   ├── DATABASE_SCHEMA.md     # Data model
│   └── ...
│
├── rules/                     # Game balance configuration
│   ├── RULESET.yaml          # Core numbers & formulas
│   ├── UNITS.yaml            # Unit definitions
│   └── ATTACKS.yaml          # Combat rules
│
├── workflows/                 # Agentic development
│   └── AGENTIC_WORKFLOW.yaml # Agent roles & flow
│
├── prompts/                   # LLM prompts for agents
│   ├── dev/                   # Dev agent prompts
│   ├── test/                  # Test agent prompts
│   └── balance/               # Balance evaluation prompts
│
├── Makefile                   # Development commands
├── .env.example               # Environment template
├── docker-compose.yml         # Container orchestration
└── README.md                  # This file
```

## Key Commands

```bash
# Development
make dev              # Start backend + frontend
make backend          # Start backend only
make frontend         # Start frontend only
make setup           # Install dependencies

# Quality
make lint            # Run linters
make test            # Run all tests
make test-backend    # Backend tests
make test-frontend   # Frontend tests

# Docker
make docker-up       # Start containers
make docker-down     # Stop containers
make db-init         # Initialize database

# Help
make help            # Show all commands
```

## Architecture Overview

### Backend (Go + Gin + MongoDB)

**Core Layers**:
- **Handlers** - HTTP endpoints (auth, empire, actions, spells)
- **Services** - Business logic (tick calculation, resource generation, combat)
- **Models** - Database models (Player, Empire, Building, Attack)
- **Config** - Environment & ruleset configuration

**Key Features**:
- RESTful API matching [API_CONTRACT.md](docs/API_CONTRACT.md)
- MongoDB document store for flexible empire state
- Deterministic game resolution (all calculations from [RULESET.yaml](rules/RULESET.yaml))
- JWT authentication

### Frontend (React + TypeScript + Vite + Tailwind)

**Key Layers**:
- **Screens** - Full-page components (ManagementScreen)
- **Sections** - Reusable dashboard sections (Resources, Buildings, Civilization)
- **Components** - Atomic UI components (ResourceRow, Slider, ProgressBar)
- **Hooks** - State management (useEmpireState, useResourceManager)
- **Services** - API clients (empireAPI, spellsAPI)

**Responsive Design**:
- Mobile-first architecture
- Single column on mobile, 2 columns on web (768px+)
- Collapsible sections for content density
- Minimalist color scheme

## Game Mechanics (TL;DR)

- **Ticks** - Time resource regenerating 1 per 5 minutes (capped at 500)
- **Actions** - Explore, Meditate, Drill, Farm (each costs 1 tick, gains skill XP)
- **Resources** - Land, Tomes, Troops, Food, Gold, Civilians
- **Buildings** - Farms, Barracks, Towers, Bazaars, Meditation Towers
- **Skills** - 4 action types with progression 0-10 (bonuses scale +20% per level)
- **Spells** - Love (morale) & Shield (defense) + future spells
- **PvP** - Military strikes & raids affect land, troops, resources
- **Clans** - Player groups with shared chat
- **Market** - Trade resources with other players
- **Morale** - Combat multiplier (down on attacks, up on non-offensive actions)

See [DESIGN_DECISIONS.md](docs/DESIGN_DECISIONS.md) & [ENGINE_CONTRACT.md](docs/ENGINE_CONTRACT.md) for full rules.

## Agentic Development

This project uses **agentic development workflows** to coordinate feature work:

- **Dev Agent** - Implements features per [ENGINE_CONTRACT.md](docs/ENGINE_CONTRACT.md) & [RULESET.yaml](rules/RULESET.yaml)
- **Test Agent** - Verifies correctness, attempts to break invariants
- **Prompts** - Handoff payloads between agents for clear requirements

See [SETUP.md](SETUP.md) for agent-specific instructions.

## Environment Variables

All variables are optional (defaults work for local development). Create `.env` from template:

```bash
cp .env.example .env
```

Key variables:
- `BACKEND_PORT`, `FRONTEND_PORT` - Server ports (default: 8080, 3000)
- `DB_HOST`, `DB_NAME` - MongoDB connection
- `JWT_SECRET` - Auth token secret
- `TICK_INTERVAL_SECONDS` - Tick regen rate (default: 300)

## Testing

```bash
# Run all tests
make test

# Backend (Go)
make test-backend
cd backend && go test -v ./...

# Frontend (React)
make test-frontend
cd frontend && npm test -- --coverage
```

## Troubleshooting

### Backend won't start
```bash
# Check MongoDB is running (Docker)
docker-compose ps

# Check port 8080 is free
lsof -i :8080
```

### Frontend won't connect to backend
- Verify `REACT_APP_API_URL=http://localhost:8080/api` in `.env`
- Check backend is running on port 8080
- Check browser console for CORS errors

### Database issues
```bash
# Reinitialize database
make docker-down
rm -rf backend/scripts/mongo_data
make docker-up
```

## Learning Resources

1. **New to the project?** → Start with [SETUP.md](SETUP.md)
2. **Understand game rules?** → Read [ENGINE_CONTRACT.md](docs/ENGINE_CONTRACT.md)
3. **Working on UI?** → See [COMPONENT_STRUCTURE.md](docs/COMPONENT_STRUCTURE.md)
4. **Contributing code?** → Follow [DEVELOPMENT_STANDARDS.md](docs/DEVELOPMENT_STANDARDS.md)
5. **Agentic development?** → See [AGENTIC_WORKFLOW.yaml](workflows/AGENTIC_WORKFLOW.yaml)
6. **On a sprint?** → Check [SPRINT_QUICK_REF.md](SPRINT_QUICK_REF.md) & [SPRINT_PLAN.md](SPRINT_PLAN.md)

## Current Phase

**Phase 1: Core Rules + Game Loop**
- Core action system (explore, meditate, drill, farm)
- Resource generation & management
- Skill progression with streaks
- Basic combat (morale system)

**Sprints**: See [SPRINT_PLAN.md](SPRINT_PLAN.md) for detailed 8-week roadmap with deliverables

**Phase 2** (planned): PvP + Market
**Phase 3** (planned): Clans + Chat

## Contributing

See [DEVELOPMENT_STANDARDS.md](docs/DEVELOPMENT_STANDARDS.md) for code style, commit conventions, and i18n requirements.

## License

[Add your license here]

## Contact

Questions? Issues? See [SETUP.md](SETUP.md) for detailed onboarding.
