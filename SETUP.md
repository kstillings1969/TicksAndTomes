# SETUP.md - Detailed Development Guide

Complete onboarding for developers and agentic workflows.

## Table of Contents

1. [New Developer Onboarding](#new-developer-onboarding)
2. [Environment Setup](#environment-setup)
3. [Development Workflow](#development-workflow)
4. [Agentic Development](#agentic-development)
5. [Common Tasks](#common-tasks)
6. [Troubleshooting](#troubleshooting)

---

## New Developer Onboarding

### Step 1: Understand the Project

**Read these first** (15 min):
1. [README.md](README.md) - Project overview & quick start
2. [DESIGN_DECISIONS.md](docs/DESIGN_DECISIONS.md) - High-level architecture
3. [ENGINE_CONTRACT.md](docs/ENGINE_CONTRACT.md) - Core game mechanics (this is THE source of truth)

**Then explore**:
- [COMPONENT_STRUCTURE.md](docs/COMPONENT_STRUCTURE.md) - Frontend layout & React components
- [DEVELOPMENT_STANDARDS.md](docs/DEVELOPMENT_STANDARDS.md) - Code style & i18n conventions
- [GLOSSARY.md](docs/GLOSSARY.md) - Term definitions

### Step 2: Set Up Your Machine

1. **Install prerequisites**:
   ```bash
   # macOS
   brew install golang node docker
   
   # Or verify versions
   go version      # Should be 1.21+
   node --version  # Should be 18+
   npm --version   # Should be 9+
   ```

2. **Clone & configure**:
   ```bash
   cd /path/to/TicksAndTomes
   cp .env.example .env
   # Edit .env if needed (defaults work locally)
   ```

3. **Install dependencies**:
   ```bash
   make setup
   ```

### Step 3: Run the Project

```bash
# Option A: Local development (2 terminals)
# Terminal 1
make backend

# Terminal 2
make frontend

# Option B: Docker (1 command)
make docker-up
```

Visit **http://localhost:3000** to see frontend.

### Step 4: Verify Installation

**Backend check**:
```bash
curl http://localhost:8080/health
# Expected: {"status":"ok"}
```

**Frontend check**:
- Open http://localhost:3000
- Should see management dashboard (stub)

**Database check**:
```bash
# If using Docker
docker-compose exec mongo mongosh --authenticationDatabase admin -u root -p rootpassword --eval "db.adminCommand('ping')"
# Expected: { ok: 1 }
```

### Step 5: Read the Code

Start with these files:
- **Backend**: `backend/cmd/server/main.go` - Entry point & routes
- **Frontend**: `frontend/src/App.tsx` → `components/screens/ManagementScreen.tsx`
- **Config**: `backend/internal/config/config.go`
- **Types**: `docs/DATABASE_SCHEMA.md` - Data model reference

---

## Environment Setup

### Configuration Files

#### `.env` - Local Configuration

```bash
cp .env.example .env
```

**Key variables for local development**:
- `BACKEND_PORT=8080` - Backend server port
- `FRONTEND_PORT=3000` - Frontend dev server
- `DB_HOST=localhost` - Local MongoDB
- `REACT_APP_API_URL=http://localhost:8080/api` - API endpoint

**For production**, replace with actual values and never share `.env`.

#### `docker-compose.yml` - Container Orchestration

Starts all services:
- **MongoDB** (port 27017) - Database
- **Backend** (port 8080) - Go server
- **Frontend** (port 3000) - React dev server

```bash
make docker-up
# Services available in ~10s
make docker-down
```

### Database Setup

**Local MongoDB**:
```bash
# Install: brew install mongodb-community
# Start: brew services start mongodb-community
# Verify: mongosh
```

**Docker MongoDB**:
```bash
make docker-up
# Already initialized with collections & indexes
```

**Manual Init** (if needed):
```bash
cd backend
go run ./cmd/migrate/main.go      # Run migrations
# Or use mongo shell directly
mongosh admin -u root -p rootpassword < scripts/init-mongo.js
```

---

## Development Workflow

### Daily Commands

```bash
# Start development
make setup                    # First time only
make dev                      # Or in 2 terminals: make backend && make frontend

# Quality checks
make lint                     # Check code style
make test                     # Run all tests

# Debugging
make backend                  # Backend logs to terminal
docker-compose logs -f mongo  # MongoDB logs
```

### Code Changes & Reload

**Backend**:
- Edit Go files in `backend/`
- Changes auto-reload (Gin watch mode recommended)
- Or restart: `Ctrl+C` → `make backend`

**Frontend**:
- Edit React files in `frontend/src/`
- Hot reload automatic (Vite)
- Check browser console for errors

### Git Workflow

```bash
# Create feature branch
git checkout -b feat/epic-name

# Make changes, commit regularly
git add .
git commit -m "feat: description of change"

# Push & create PR
git push origin feat/epic-name
```

**Commit message format**:
```
feat: Add tax rate slider to civilization section
fix: Correct morale calculation in combat
docs: Update ENGINE_CONTRACT with new defense formula
test: Add test for negative gold handling
chore: Update dependencies
```

---

## Agentic Development

This project uses **autonomous agents** coordinating via handoff payloads.

### Key Roles

**Dev Agent**:
- Implements features per [ENGINE_CONTRACT.md](docs/ENGINE_CONTRACT.md)
- Follows [RULESET.yaml](rules/RULESET.yaml) for all balance values
- Creates tests, documents changes
- **Constraint**: Never weaken invariants (Tom Tower 100:1, Spell Shield mitigation, etc.)

**Test Agent**:
- Verifies implementation against invariants
- Creates property-based & boundary tests
- Reports defects in `test_to_dev` payload
- **Constraint**: Cannot modify production code

**Balance Agent** (future):
- Evaluates ruleset changes
- Simulates economic/combat scenarios
- Recommends balance adjustments

### Agent Workflow Cycle

```
1. Human requests feature → specifies in GitHub issue
2. Dev Agent implements
   - Reads ENGINE_CONTRACT + RULESET
   - Creates/modifies code + tests
   - Outputs `dev_to_test` handoff payload
3. Test Agent verifies
   - Runs tests, attempts falsification
   - Reports pass/fail in `test_to_dev` payload
4. If pass → Merged
   If fail → Back to Dev Agent with defects
```

### Agent Prompts

Located in `prompts/`:
- `dev/IMPLEMENT_FEATURES.md` - Dev agent constraints & output format
- `test/VERIFY_INVARIANTS.md` - Test agent checks
- `balance/EVALUATE_RULES.md` - Balance evaluation

See `workflows/AGENTIC_WORKFLOW.yaml` for agent definitions.

### Developer as Agent

If you're working as an agent:

1. **Always read**:
   - `docs/ENGINE_CONTRACT.md` (source of truth)
   - `rules/RULESET.yaml` (all numbers)
   - `docs/DATABASE_SCHEMA.md` (data model)

2. **When implementing**:
   - Keep resolution deterministic (same inputs = same outputs)
   - Use RULESET for all balance values
   - Add tests alongside code
   - Create `dev_to_test` handoff payload:
     ```yaml
     change_intent: "Implement love spell morale restoration"
     affected_files:
       - backend/internal/spells/love.go
       - backend/internal/tests/spells_test.go
     risks:
       - Morale calculation might exceed 100
       - Tome requirement validation critical
     tests_expected:
       - Morale capped at 100
       - Insufficient tomes rejects spell
       - Random bonus within 4-10 range
     ```

3. **Document decisions**:
   - If ENGINE_CONTRACT is ambiguous → note it
   - If RULESET needs clarification → flag it
   - Prefer questions over assumptions

### Resolving Ambiguities

If ENGINE_CONTRACT or RULESET is unclear:

1. **Note the ambiguity** in your change summary
2. **Make conservative assumption** (usually the most restrictive)
3. **Leave a TODO comment** with the question
4. **Flag in test handoff** for review

Example:
```go
// TODO: ACTION_PROGRESSION experience multiplier not specified in RULESET
// Assuming multiplicative (1 + level * 0.2) not additive
// Confirm in ENGINE_CONTRACT v2
```

---

## Common Tasks

### Implementing a New Feature

**Example**: Add fire spell

1. **Read requirements**:
   - Check `docs/` for spell mechanics
   - Check `rules/RULESET.yaml` for balance numbers
   - Check `API_CONTRACT.md` for endpoints

2. **Implement backend**:
   ```bash
   cd backend
   # Edit internal/spells/fire.go
   # Edit internal/handlers/spells.go
   # Add spell to main.go routes
   # Create tests in internal/tests/fire_test.go
   ```

3. **Implement frontend**:
   ```bash
   cd frontend
   # Add SpellButton component
   # Add spell API call to services/spellsAPI.ts
   # Update UI to show spell in action panel
   ```

4. **Test everything**:
   ```bash
   make test-backend
   make test-frontend
   ```

5. **Create handoff payload** (if using agent workflow):
   ```yaml
   change_intent: Add fire spell with 60:1 tome requirement
   affected_files: [...]
   risks: Spell damage might unbalance PvP
   tests_expected: Tome check, damage calculation, cooldown
   ```

### Adding a New Endpoint

1. **Update API_CONTRACT.md** with endpoint spec
2. **Add handler** in `backend/internal/handlers/ `
3. **Register route** in `backend/cmd/server/main.go`
4. **Add test** in `backend/internal/tests/`
5. **Add frontend client** in `frontend/src/services/`

Example handler:
```go
func CastFireSpell(c *gin.Context) {
    var req CastSpellRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    // Validate via RULESET
    // Execute spell logic
    // Return result
    
    c.JSON(200, gin.H{"success": true})
}
```

### Writing Tests

**Backend** (Go):
```go
// backend/internal/tests/spells_test.go
func TestLoveSpellMoraleRestore(t *testing.T) {
    // Arrange
    empire := createTestEmpire()
    empire.Morale = 50
    empire.Tomes = 2100  // 20 per meditation tower
    
    // Act
    result := CastLoveSpell(empire)
    
    // Assert
    assert.True(t, result.Morale >= 54)  // 50 + 2 ticks + 4 min spell
    assert.True(t, result.Morale <= 62)  // 50 + 2 ticks + 10 max spell
    assert.True(t, result.Morale <= 100) // Capped
}
```

**Frontend** (React):
```typescript
// frontend/src/components/__tests__/ResourceRow.test.tsx
import { render, screen } from '@testing-library/react'
import { ResourceRow } from '@/components/common/ResourceRow'

test('displays resource with surplus status', () => {
    render(
        <ResourceRow
            label="Gold"
            current={12987}
            perTick={+2389}
            status="surplus"
        />
    )
    
    expect(screen.getByText('Gold')).toBeInTheDocument()
    expect(screen.getByText('12,987')).toBeInTheDocument()
    expect(screen.getByText('+2,389/tick')).toBeInTheDocument()
})
```

### Updating Rules

**If you need to change RULESET.yaml**:

1. **Edit** `rules/RULESET.yaml` with new numbers
2. **Note impact** in a comment:
   ```yaml
   morale:
     starting_value: 100  # Changed from 50 (impact: gives players more buffer)
   ```

3. **Update ENGINE_CONTRACT.md** to match
4. **Update tests** to match new expected values
5. **Create balance evaluation** in handoff payload

Never update RULESET without updating ENGINE_CONTRACT + tests.

### Debugging

**Backend**:
```bash
# Add logging
import "log"
log.Printf("Value: %v", someVar)

# Or use debugger
# Use Delve: https://github.com/go-delve/delve
dlv debug ./cmd/server
```

**Frontend**:
```tsx
// Browser console (F12)
console.log('Empire:', empire)

// React DevTools extension
// Check component renders & state changes

// Network tab
// Inspect API calls to backend
```

**Database**:
```bash
# Connect with mongo shell
mongosh -u root -p rootpassword --authenticationDatabase admin

# List collections
show collections

# Query player
db.players.findOne()

# Check indexes
db.players.getIndexes()
```

---

## Troubleshooting

### "Port 8080 already in use"

```bash
# Find process using port
lsof -i :8080

# Kill it (if safe)
kill -9 <PID>

# Or use different port
BACKEND_PORT=8081 make backend
```

### "MongoDB connection refused"

```bash
# Check if running
docker-compose ps

# If using Docker
make docker-up

# If local MongoDB
brew services start mongodb-community
mongosh  # Verify connection
```

### "Frontend can't reach backend"

1. Check backend is running: `curl http://localhost:8080/health`
2. Check REACT_APP_API_URL in `.env`
3. Check browser console for CORS errors
4. Verify ports match (backend 8080, frontend 3000)

### "Tests failing"

```bash
# Run with verbose output
make test-backend
cd backend && go test -v ./... -run TestName

make test-frontend
cd frontend && npm test -- MyComponent.test --verbose
```

Check test setup in:
- Backend: `backend/internal/tests/setup.go`
- Frontend: `frontend/package.json` jest config

### "Linter errors"

```bash
# See what's wrong
make lint

# Fix automatically
make lint:fix   # Frontend
# Backend linter fixes vary (review manually)
```

---

## Next Steps

1. **Complete onboarding**: Follow steps 1-5 above
2. **Pick a task**: Check GitHub issues labeled `good-first-issue`
3. **Create feature branch**: `git checkout -b feat/your-feature`
4. **Implement + test**: Follow workflows above
5. **Submit PR**: Reference issue number

**Questions?** 
- Check [GLOSSARY.md](docs/GLOSSARY.md) for term definitions
- Read relevant docs in `docs/` folder
- Ask in GitHub discussions

Happy coding! 🎮
