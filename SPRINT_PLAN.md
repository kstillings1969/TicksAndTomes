# Development Sprint Plan — TicksAndTomes

**Project Phases**: 3 phases total | **Current Status**: Phase 1 (Core Loop) | **Target MVP**: 12 weeks

---

## Phase 1: Rules + Core Loop (Weeks 1-8)

Single-player game engine with tick-based resource management, skill progression, and building system.

### Sprint 1.1: Game Engine Foundation (Weeks 1-2)

**Goal**: Build core tick system and action resolution engine.

**Tasks**:

Backend:
- [ ] Implement tick regeneration system
  - [ ] Tick cap (500 active, 200 box)
  - [ ] Regen rate (1 per 5 minutes)
  - [ ] Timestamp tracking
  - Tests: Unit tests for tick accumulation, overflow, regen rate
  
- [ ] Create action resolution engine
  - [ ] Action types enum (explore, meditate, drill, farm)
  - [ ] Base output calculations (50 land, 15 tomes, etc.)
  - [ ] Skill modifier application (level 0-10, +20% per level)
  - Tests: Verify all action outputs scale correctly with skill levels

- [ ] Implement skill progression system
  - [ ] Experience gain (0.01 per action)
  - [ ] Streak decay (0.85^n with 0.25 floor)
  - [ ] Level cap (0-10)
  - [ ] Streak reset (1 hour timeout)
  - Tests: Boundary tests for streak count, level cap, experience accumulation

Database:
- [ ] Design player collection schema (from DATABASE_SCHEMA.md)
- [ ] Design empire subdocument structure
- [ ] Set up MongoDB indices (user_id unique, clan_id, status)
- [ ] Create database initialization script

Configuration:
- [ ] Load RULESET.yaml values into config
- [ ] Verify all action outputs match RULESET
- [ ] Add environment variable overrides for testing

**Definition of Done**:
- All unit tests passing (>80% coverage)
- Tick system verified deterministic
- Actions produce expected outputs per RULESET
- Database indices created

**Deliverables**:
- ✅ Tick regeneration working
- ✅ Actions resolve correctly
- ✅ Skill progression functional
- ✅ Database schema initialized

---

### Sprint 1.2: Resource Management (Weeks 2-3)

**Goal**: Implement resource generation, control, and balance.

**Tasks**:

Backend:
- [ ] Implement resource calculations
  - [ ] Civilian generation (1 per empty acre per tick)
  - [ ] Tax rate control (0-100%, multiplier logic)
  - [ ] Gold generation (civilians * tax_rate / 100)
  - [ ] Automatic resource decay on action (food consumed, etc.)
  - Tests: Tax thresholds (9%, 23%, 35%), civilian cap (land * 1000)

- [ ] Implement building system
  - [ ] Building types (farms, barracks, towers, bazaars, meditation towers)
  - [ ] Production formulas per building (10 grain/farm, 15 troops/barracks, etc.)
  - [ ] Modifier application (farm_level, drill_level, tax_rate)
  - Tests: Verify production scales with modifiers

- [ ] Create tick simulation engine
  - [ ] Apply all production (buildings, actions, taxes)
  - [ ] Calculate per-tick deltas
  - [ ] Update empire state atomically
  - Tests: Simulate 100 ticks, verify state consistency

- [ ] Implement morale system
  - [ ] Initial morale (100%)
  - [ ] Offensive vs non-offensive action changes (-1% / +1%)
  - [ ] Bounds checking (0-100%)
  - Tests: Edge cases at 0% and 100%

Frontend:
- [ ] Create ResourcesSection component
  - [ ] Display gold, grain, tomes, troops, ticks
  - [ ] Show per-tick deltas with status colors
  - [ ] Format large numbers with K/M/B abbreviations
  
- [ ] Create BuildingsSection component
  - [ ] Display building counts
  - [ ] Show production per building type

**Definition of Done**:
- All resource calculations match ENGINE_CONTRACT
- Tax system working correctly (bonuses/penalties)
- Morale boundaries enforced
- Frontend displays resources with deltas

**Deliverables**:
- ✅ Resource generation system complete
- ✅ Tax rate affects economy correctly
- ✅ Morale system functional
- ✅ UI shows resource status

---

### Sprint 1.3: Spells & Defense (Weeks 4-5)

**Goal**: Implement spell system and Tome Tower defense.

**Tasks**:

Backend:
- [ ] Implement Tome Tower defense system
  - [ ] Defense formula (filled_towers * 100)
  - [ ] Filled tower logic (min(towers, floor(tomes / 100)))
  - [ ] Display defense status
  - Tests: Boundary tests at 100, 199, 200 tomes per tower

- [ ] Implement Love spell
  - [ ] Tome requirement validation (20:1 ratio)
  - [ ] Morale restoration (6-12 total: 2 ticks + 4-10 spell)
  - [ ] Seeded RNG for determinism
  - [ ] Spell cost (2 ticks)
  - Tests: Morale cap at 100, tome requirement strict

- [ ] Implement Spell Shield
  - [ ] Activation/deactivation
  - [ ] Mitigation factor (0.33x spell damage)
  - Tests: Verify 33%, not immunity

- [ ] Create spell API endpoints
  - [ ] POST /api/spell/love
  - [ ] POST /api/spell/shield
  - [ ] Response with new empire state

Frontend:
- [ ] Create ActionSection component
  - [ ] Buttons for explore, meditate, drill, farm
  - [ ] Show action costs (1 tick each)
  - [ ] Show skill level + XP progress
  
- [ ] Create SpellSection component
  - [ ] Love spell button (shows morale impact)
  - [ ] Shield spell toggle (shows defense impact)
  - [ ] Tome requirement display

**Definition of Done**:
- Spell system passes all invariant tests
- Tome Tower defense formula exact
- Love spell restoration within range
- UI shows all spell options

**Deliverables**:
- ✅ Spells functional and tested
- ✅ Defense system working
- ✅ Spell UI integrated

---

### Sprint 1.4: Civilization Management (Weeks 5-6)

**Goal**: Complete economic simulation with civilian & tax management.

**Tasks**:

Backend:
- [ ] Implement advanced civilian growth
  - [ ] Growth rate multiplier per tax rate
  - [ ] Low tax bonus (+25% at 0-9%)
  - [ ] High tax penalty (-15% to -35% at 23-34%)
  - [ ] Exodus logic (losses at 35%+)
  - [ ] Pro-rated penalty calculation
  - Tests: All tax brackets, verify multipliers

- [ ] Implement tax control API
  - [ ] PUT /api/empire (update tax_rate)
  - [ ] Return new gold/civilian projections
  - [ ] Debounce support (for slider)

Frontend:
- [ ] Create CivilizationSection component
  - [ ] Display current civilians
  - [ ] Show civilians cap (land * 1000)
  - [ ] Tax rate slider (0-100%)
  - [ ] Tax effect labels (bonus/penalty/exodus)
  - [ ] Gold per tick projection
  
- [ ] Add EmpireHeader component
  - [ ] Sticky header with empire name
  - [ ] Quick status indicator
  - [ ] Land display

- [ ] Implement useResourceManager hook
  - [ ] Calculate per-tick deltas
  - [ ] Apply skill modifiers
  - [ ] Determine resource status (surplus/deficit/neutral)
  - [ ] Memoized calculations

**Definition of Done**:
- Tax system matches ENGINE_CONTRACT formulas
- Civilian growth/exodus working
- UI slider updates gold projections
- All calculations deterministic

**Deliverables**:
- ✅ Full economic simulation
- ✅ Citizen management UI
- ✅ Tax control functional

---

### Sprint 1.5: Dashboard Integration (Weeks 6-7)

**Goal**: Complete management screen UI with all sections.

**Tasks**:

Frontend:
- [ ] Implement useEmpireState hook
  - [ ] Fetch empire data from API
  - [ ] Refresh every 5 seconds
  - [ ] Error handling & retry
  - [ ] Caching strategy

- [ ] Create ActionProgressionSection
  - [ ] Display all 4 actions (explore, meditate, drill, farm)
  - [ ] Show level (0/10) and XP progress bar
  - [ ] Show bonus output percentage
  
- [ ] Integrate all sections into ManagementScreen
  - [ ] Sticky header
  - [ ] Resource section (collapsed by default on mobile)
  - [ ] Buildings section
  - [ ] Civilization section
  - [ ] Action progression section
  
- [ ] Implement responsive layout
  - [ ] Mobile: Single column, collapsible sections
  - [ ] Web (768px+): Two columns (TBD layout)
  - [ ] Smooth animations for expand/collapse
  
- [ ] Add tailwind configuration
  - [ ] Custom colors from variables.css
  - [ ] Responsive breakpoints
  - [ ] Component styling

**Definition of Done**:
- All sections rendering correctly
- Data flows from API to UI
- Responsive design working
- No console errors

**Deliverables**:
- ✅ Complete dashboard UI
- ✅ Responsive design
- ✅ API integration complete

---

### Sprint 1.6: Testing & Polish (Week 7)

**Goal**: Comprehensive testing, documentation, bug fixes.

**Tasks**:

Backend:
- [ ] Run full test suite
  - [ ] Unit tests (engine, calculations, formulas)
  - [ ] Integration tests (action → resource change → UI update)
  - [ ] Boundary tests (all invariants)
  - Target: >85% coverage
  
- [ ] Test against all invariants
  - [ ] Morale [0, 100]
  - [ ] Tome Tower defense (filled_towers * 100)
  - [ ] Skill cap at 10
  - [ ] Determinism (identical seeds → identical results)

Frontend:
- [ ] Component unit tests
  - [ ] ResourceRow rendering
  - [ ] Slider value changes
  - [ ] Status badge color mapping
  - [ ] Progress bar calculations
  
- [ ] E2E tests (optional for Phase 1)
  - [ ] Load empire data
  - [ ] Perform action
  - [ ] Verify UI updates
  - [ ] Change tax rate, verify projection

Documentation:
- [ ] API documentation
  - [ ] All endpoints documented
  - [ ] Request/response examples
  - [ ] Error codes
  
- [ ] Developer guide
  - [ ] Code structure overview
  - [ ] How to add new actions
  - [ ] How to add new resources

**Definition of Done**:
- All tests passing (>85% coverage)
- No console/terminal errors
- All invariants verified
- Code documented

**Deliverables**:
- ✅ Test suite complete
- ✅ Bug-free Phase 1
- ✅ Developer documentation

---

### Sprint 1.7: Performance & Optimization (Week 8)

**Goal**: Optimize for production, prepare for Phase 2.

**Tasks**:

Backend:
- [ ] Performance optimization
  - [ ] Query optimization (MongoDB)
  - [ ] Batch operations where possible
  - [ ] Cache frequently accessed data
  
- [ ] Load testing
  - [ ] 100 concurrent users
  - [ ] 1000 tick simulations
  - [ ] Measure response times
  
- [ ] Error handling
  - [ ] Graceful degradation
  - [ ] User-friendly error messages
  - [ ] Logging for debugging

Frontend:
- [ ] Performance optimization
  - [ ] Code splitting with React.lazy
  - [ ] Memoization (useMemo, useCallback)
  - [ ] Image optimization
  - [ ] Bundle size analysis
  
- [ ] Browser compatibility
  - [ ] Test on Chrome, Safari, Firefox
  - [ ] Mobile responsiveness
  - [ ] Touch interactions
  
- [ ] Accessibility
  - [ ] WCAG AA compliance
  - [ ] Screen reader testing
  - [ ] Keyboard navigation

**Definition of Done**:
- Load tests passing
- Bundle size <150KB
- Performance benchmarks set
- Accessibility checklist complete

**Deliverables**:
- ✅ Production-ready Phase 1
- ✅ Performance baseline established
- ✅ Ready for Phase 2 (PvP)

---

## Phase 2: PvP + Market (Weeks 9-12)

**Goal**: Add combat, raiding, and player-to-player trading.

### Sprint 2.1: Combat System (Weeks 9-10)

**Tasks**:
- [ ] Implement attack mechanics
  - [ ] Military strike formula
  - [ ] Damage calculation with morale multiplier
  - [ ] Casualty calculations
  
- [ ] Implement morale in combat
  - [ ] Pre-attack validation
  - [ ] Morale decay on offensive actions
  - [ ] Recovery mechanics
  
- [ ] Create attack resolution
  - [ ] Defender notification
  - [ ] Damage application
  - [ ] Loot/loss calculations
  
- [ ] Add API endpoints
  - [ ] POST /api/action/military_strike
  - [ ] POST /api/action/military_raid
  - [ ] GET /api/empire/:id (opponent info)

**Deliverables**:
- ✅ Combat system functional
- ✅ Attacks resolve correctly
- ✅ Morale affects outcomes

---

### Sprint 2.2: Market System (Weeks 10-11)

**Tasks**:
- [ ] Implement market data model
  - [ ] Listings (resource, price, quantity)
  - [ ] Transactions (buyer, seller, amount)
  - [ ] History/audit log
  
- [ ] Create market API
  - [ ] GET /api/market/listings
  - [ ] POST /api/market/listings (create)
  - [ ] POST /api/market/trade (execute)
  
- [ ] Implement price discovery
  - [ ] Moving average prices
  - [ ] Price history
  - [ ] Supply/demand indicators
  
- [ ] Frontend market UI
  - [ ] Listing browser
  - [ ] Price charts
  - [ ] Order placement

**Deliverables**:
- ✅ Market operational
- ✅ Players can trade
- ✅ Price discovery working

---

### Sprint 2.3: Leaderboards & Stats (Week 11-12)

**Tasks**:
- [ ] Implement leaderboards
  - [ ] By net worth
  - [ ] By empire size
  - [ ] By kill count
  
- [ ] Player profiles
  - [ ] Stats display
  - [ ] Empire history
  - [ ] Attack history
  
- [ ] Achievements (optional)
- [ ] Seasonal resets (optional)

**Deliverables**:
- ✅ Leaderboards functional
- ✅ Player profiles visible
- ✅ Phase 2 complete

---

## Phase 3: Clans + Chat (Future)

Goals:
- Implement clan system
- Add chat (global, clan, DM)
- Clan warfare mechanics
- Alliance features (optional)

---

## Development Metrics & Checkpoints

### Weekly Standup (Every Monday)

- [ ] Sprint progress update
- [ ] Blockers & dependencies
- [ ] Risk assessment
- [ ] Resource allocation review

### Sprint Retrospective (End of each sprint)

- [ ] What went well?
- [ ] What didn't?
- [ ] Action items for next sprint

### Test Coverage

- **Phase 1 Target**: >85% coverage
- **Phase 2 Target**: >80% coverage
- **Invariants**: 100% tested

### Performance Targets

- API response time: <200ms (p95)
- Frontend load time: <3s (LTE)
- Database query time: <50ms (p95)

---

## Resource Requirements

### Team Composition (Suggested)

- **Backend (Go)**: 1-2 engineers
- **Frontend (React)**: 1 engineer
- **QA/Testing**: 0.5 engineer (early phases)
- **DevOps**: 0.5 engineer (setup complete)

### Tools & Infrastructure

- **Repository**: GitHub (✅ configured)
- **CI/CD**: GitHub Actions (to setup)
- **Monitoring**: TBD
- **Error tracking**: TBD

---

## Risk Register

| Risk | Impact | Probability | Mitigation |
|------|--------|-------------|-----------|
| Performance bottleneck at 1000 users | HIGH | MEDIUM | Load test weekly, optimize early |
| Database schema issues | HIGH | LOW | Schema review before Phase 2 |
| Morale system breaks game balance | HIGH | MEDIUM | Balance eval every sprint |
| React state management complexity | MEDIUM | MEDIUM | Plan for Context/Redux if needed |
| MongoDB document size limits | MEDIUM | LOW | Monitor document growth, archive old data |

---

## Next Steps

1. **Kick off Sprint 1.1** — Begin game engine foundation
2. **Set up CI/CD** — GitHub Actions for tests on every PR
3. **Assign tasks** — Allocate development resources
4. **Schedule standup** — Weekly Monday sync
5. **Track progress** — Use GitHub issues with sprint labels

See [SETUP.md](SETUP.md) for onboarding developers to this plan.

For detailed task breakdown, see [GitHub Issues](https://github.com/kenstillings/ticks-and-tomes) (to be created with sprint tags).
