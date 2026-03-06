# Sprint Quick Reference

**Current Status**: Before Sprint 1.1 | **Today**: March 6, 2026

---

## Sprint 1.1: Game Engine Foundation (Weeks 1-2)

### Primary Goal
Build core tick system, action resolution, and skill progression.

### Key Deliverables
- ✅ Tick regeneration (500 cap, 1 per 5 min)
- ✅ Action resolution (explore, meditate, drill, farm)
- ✅ Skill progression (0-10 levels, +20% per level)
- ✅ Streak decay (0.85^n, 0.25 floor)
- ✅ Player/Empire database schema

### Acceptance Criteria
- [ ] `make test-backend` passes with >80% coverage
- [ ] All actions produce expected outputs (matching RULESET.yaml)
- [ ] Tick system is deterministic (same seed = same results)
- [ ] Database indices created and optimized

### Critical Dependencies
- RULESET.yaml (source of truth for all numbers)
- ENGINE_CONTRACT.md (specification)
- DATABASE_SCHEMA.md (data model)

### Resources
- **Backend** Lead: [Assign name]
- **Database**: [Assign name]
- **Tests**: [Assign name]
- **Frontend**: Can start Setup work in parallel

---

## Key Files to Reference During Development

### Specifications (DO NOT DEVIATE)
1. `docs/ENGINE_CONTRACT.md` — Game mechanics & formulas
2. `rules/RULESET.yaml` — All balance numbers
3. `docs/DATABASE_SCHEMA.md` — Data model

### Agent Prompts (If using agents)
1. `prompts/dev/IMPLEMENT_FEATURES.md` — Dev agent instructions
2. `prompts/test/VERIFY_INVARIANTS.md` — Test agent checklist
3. `prompts/balance/EVALUATE_RULES.md` — Balance evaluation

### Development Standards
- `docs/DEVELOPMENT_STANDARDS.md` — i18n, naming, code style
- `docs/DECISIONS_RESOLVED.md` — TDD decisions (10 key decisions resolved)
- `SETUP.md` — Developer onboarding

### Project Info
- `README.md` — Project overview
- `Makefile` — Development commands
- `SPRINT_PLAN.md` — Full sprint roadmap

---

## Daily Workflow

### Morning (Start of Day)
1. Check blockers from yesterday
2. Verify latest `main` branch code builds
3. Review your assigned tasks for the day
4. Start coding

### During Day
- Commit early & often
- Run tests frequently (`make test-backend` or `make test-frontend`)
- Push to feature branch

### Evening (End of Day)
- Ensure tests still pass
- Push final changes
- Leave notes on blockers/concerns

### Before Sprint Review
- [ ] All commits pushed
- [ ] All tests passing
- [ ] Code documented
- [ ] PR created with sprint label

---

## Common Commands

```bash
# Setup & installation
make setup              # First time only

# Development
make backend            # Start backend server
make frontend           # Start frontend dev server
make dev              # Start both (in background)

# Testing & quality
make test             # Run all tests
make test-backend     # Backend tests only
make test-frontend    # Frontend tests only
make lint             # Check code style

# Database
docker-compose up     # Start MongoDB
docker-compose down   # Stop MongoDB

# Help
make help             # Show all available commands
```

---

## Debugging Tips

### Backend (Go)
```bash
# Run specific test
cd backend && go test -v ./internal/tests -run TestTickRegen

# View detailed logs
cd backend && go run ./cmd/server/main.go 2>&1

# Check MongoDB connection
mongosh -u root -p rootpassword --authenticationDatabase admin
```

### Frontend (React)
```bash
# Run specific test
cd frontend && npm test -- ResourceRow.test

# View network requests
# Open DevTools → Network tab
# Watch console for errors
console.log('Debug:', value)  // Add and remove as needed
```

---

## When Stuck

1. **Check the docs first**
   - ENGINE_CONTRACT.md has the detailed spec
   - RULESET.yaml has all numbers
   - SETUP.md has troubleshooting

2. **Ask in standup or chat**
   - Blockers get raised immediately
   - No solo debugging for >30 min

3. **Check test failures**
   - Read test error message carefully
   - Tests are the spec

4. **Review game logic**
   - Is the calculation correct per ENGINE_CONTRACT?
   - Are all RULESET numbers being used?
   - Is the formula deterministic?

---

## Testing Checklist Before Committing

- [ ] All tests pass locally (`make test`)
- [ ] No console errors (backend or frontend)
- [ ] New code has tests (or updated existing)
- [ ] Code follows naming conventions
- [ ] No hardcoded values (use RULESET)
- [ ] Commit message is clear

---

## Sprint Tracking

**Sprint 1.1 Progress**:
- Week 1 (Mar 6-12): Engine foundation (ticks, actions, skills)
- Week 2 (Mar 13-19): Database setup, schema validation

**Weekly Standups**:
- Monday 10am: Full team standup
- Thursday 2pm: Progress check-in
- Friday 4pm: Pre-review testing

**Sprint Review**: Friday, March 19

---

## Contacts & Escalation

- **Technical Blocker**: Post in #dev channel or mention in standup
- **Design Question**: Reference ENGINE_CONTRACT.md or ask lead
- **Balance Issue**: Flag for balance agent in code comments
- **Schedule/Resource**: Escalate to project lead

---

## Phase 1 (Weeks 1-8) Timeline

| Sprint | Weeks | Focus | Status |
|--------|-------|-------|--------|
| 1.1 | 1-2 | Engine foundation | Starting →
| 1.2 | 2-3 | Resources & buildings | Queued |
| 1.3 | 4-5 | Spells & defense | Queued |
| 1.4 | 5-6 | Civilization management | Queued |
| 1.5 | 6-7 | Dashboard UI | Queued |
| 1.6 | 7 | Testing & polish | Queued |
| 1.7 | 8 | Performance & launch | Queued |

---

## Get Started!

1. Read ENGINE_CONTRACT.md (15 min)
2. Read SETUP.md step 5 (understand the code structure)
3. Check your Sprint 1.1 assignment
4. Start coding!

Questions? Ask in standup or check SETUP.md troubleshooting.

**Let's build something great! 🎮**
