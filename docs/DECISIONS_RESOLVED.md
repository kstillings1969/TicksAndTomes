# Resolved Decisions — TicksAndTomes v5

This document resolves pending decisions from earlier design phases. Reference these when implementing.

---

## Decision 1: Action Progression Across Empire Deaths

**Status**: ✅ RESOLVED

**Question** (from DATABASE_SCHEMA.md):
> Should action progression (skill levels & experience) persist across empire deaths?
> - Option A: Reset on death — Each new empire starts fresh at level 0
> - Option B: Persist across empires — Keep action levels on account

**Decision**:
### **OPTION B: Persist Across Empires** ✅

**Rationale**:
1. **Account progression** — Rewards long-term play, not just single empire
2. **Engagement** — Returning players stay powerful, fight their way back up
3. **Alt empires** — Players can experiment with different strategies at full power
4. **Achievement preservation** — Skills earned are not permanent loss on death

**Implementation**:
- Move `action_progression` out of `empire` document
- Store on `player` document (account-level)
- Does NOT reset on empire death
- All empires share the same skill levels
- BUT: Each empire has separate resources, buildings, streaks

**Database structure**:
```typescript
// User document
{
  user_id: "string",
  email: "string",
  stars: number,
  
  // Account-level (persists across empires)
  action_progression: {
    explore: { level: 5, experience: 0.3 },
    meditate: { level: 8, experience: 0.75 },
    drill: { level: 3, experience: 0.1 },
    farm: { level: 6, experience: 0.5 }
  },
  
  // Current empire (resets on death)
  empire: {
    resources: { ... },
    buildings: { ... },
    streaks: { ... }
    // NO action_progression here
  }
}
```

**Code implications**:
- Backend: Read action levels from `player.action_progression`, not `empire.action_progression`
- Frontend: Display skill levels from player account
- Actions: XP gains update `player.action_progression`

**Impacts**:
- ✅ Skill progress preserved → Encourages restarting
- ✅ Account feels rewarding → "My skills")
- ⚠️ New players won't have high-level players' skills advantage (fair)
- ⚠️ Returning players start strong (by design)

---

## Decision 2: Streak Decay & Skill Levels

**Status**: ✅ RESOLVED

**Question** (from copilot-instructions.md & RULESET.yaml):
> How should streaks interact with skill levels? Current state: Streak decay enabled, but implementation unclear.

**Decision**:
### **Streak Decay is a Multiplier on XP Gain** ✅

**Spec** (from copilot-instructions.md):
```
streak_decay:
  enabled: true
  exponential_base: 0.85
  min_multiplier: 0.25
  reset_after_seconds: 3600  # 1 hour
```

**Formula**:
```
xp_gain = base_xp * (exponential_base ^ (streak_count - 1))
xp_gain = max(xp_gain, base_xp * min_multiplier)  # Floor at 0.25x

Where:
- base_xp = 0.01 per action (from RULESET)
- streak_count = consecutive actions of this type
- exponential_base = 0.85 (each action worth 85% of previous)
- min_multiplier = 0.25 (never go below 25% of base)
- reset = streak lost if no action of that type for 1 hour
```

**Examples**:
```
Action 1: XP = 0.01 * (0.85^0) = 0.01 (100%)
Action 2: XP = 0.01 * (0.85^1) = 0.0085 (85%)
Action 3: XP = 0.01 * (0.85^2) = 0.00723 (72.3%)
Action 4: XP = 0.01 * (0.85^3) = 0.00614 (61.4%)
Action 5: XP = 0.01 * (0.85^4) = 0.00522 (52.2%)

...continuing until:
Action 20+: XP = 0.01 * 0.25 = 0.0025 (25% floor)

If player goes 1+ hour without farming: streak resets to 0
Next farm: Back to 0.01 XP
```

**Rationale**:
- **Anti-grind**: Can't farm same action repeatedly for massive XP
- **Renewable reward**: Streak resets after 1 hour, encouraging diverse play
- **Skill still reachable**: 25% floor means slowdown, not stop

**Database structure**:
```typescript
action_progression: {
  explore: {
    level: 5,
    experience: 0.3,
    streak_count: 3,
    last_action_timestamp: 1705000000
  },
  ...
}
```

**Calculation in code**:
```go
func CalculateXPGain(baseXP float64, streakCount int32) float64 {
    decayMultiplier := math.Pow(0.85, float64(streakCount-1))
    decayMultiplier = math.Max(decayMultiplier, 0.25)  // Floor
    return baseXP * decayMultiplier
}

func ResetStreakIfExpired(skill *SkillProgress, currentTime int64) {
    if currentTime-skill.LastActionTimestamp > 3600 {
        skill.StreakCount = 0
    }
}
```

---

## Decision 3: Spell Shield vs Immunity

**Status**: ✅ RESOLVED (from copilot-instructions)

**Question**:
> Should Spell Shield (arcane_expansion) be mitigation or immunity?

**Decision**:
### **Spell Shield = Mitigation (0.33x), NOT Immunity** ✅

**Spec** (from ENGINE_CONTRACT.md):
```
if spell_shield active:
  spell_effect *= 0.33  # 33% of original damage

Examples:
- Unshielded damage: 100 → 100
- Shielded damage: 100 → 33
- Can still harm shielded players, just much less
```

**Why not immunity**:
- ✅ Creates interesting spell vs shield tactics
- ✅ Doesn't break game balance
- ❌ Immunity would make shield overpowered

**Rationale**: Spell Shield reduces spell effectiveness to 1/3, but doesn't eliminate it. Shielded players can still take damage from spells.

---

## Decision 4: Tome Tower Defense Ratio

**Status**: ✅ RESOLVED (from ENGINE_CONTRACT.md)

**Question**:
> Confirm Tome Tower defense is capped by tomes, not just tower count.

**Decision**:
### **Filled Tower = 100+ Tomes Per Tower** ✅

**Spec**:
```
filled_towers = min(towers, floor(tomes / 100))
defense = filled_towers * 100

Examples:
- 200 tomes, 5 towers → defense = min(5, 2) * 100 = 200
- 550 tomes, 5 towers → defense = min(5, 5) * 100 = 500
- 99 tomes, 5 towers → defense = min(5, 0) * 100 = 0
```

**Why this formula**:
- ✅ Tomes are precious (also used for spells)
- ✅ Creates meaningful choice: spells OR defense
- ✅ Tower investment alone insufficient

---

## Decision 5: Morale System Scope

**Status**: ✅ RESOLVED

**Question**:
> When does morale affect combat? Only direct attacks, or spells too?

**Decision**:
### **Morale Affects All Damage (Direct + Spells)** ✅

**Spec**:
```
damage_output = base_damage * (current_morale / 100)

Applies to:
- Military strikes (troops vs troops)
- Spell attacks (magic damage)
- Raids (resource theft)

Does NOT apply to:
- Spell Shield mitigation (Shield works same regardless of morale)
- Love spell restoration (restorative, not damage)
```

**Example**:
```
Magic attack base: 500 damage
Attacker morale: 60%
Actual damage: 500 * 0.6 = 300
```

---

## Decision 6: Game Phase Scope — Phase 1 MVP

**Status**: ✅ RESOLVED

**Question**:
> What's minimum viable product for Phase 1?

**Decision**:
### **Phase 1: Core Loop Only** ✅

**Includes**:
- ✅ Actions (explore, meditate, drill, farm)
- ✅ Skill progression (0-10 with streaks)
- ✅ Resources (land, tomes, troops, food, gold, civilians)
- ✅ Buildings (farms, barracks, towers, bazaars, meditation towers)
- ✅ Morale system
- ✅ Spell Shield & Love spell (defensive/restorative)
- ✅ Basic taxes & civilian management
- ✅ Tick system & regeneration

**NOT in Phase 1**:
- ❌ PvP attacks/raids (Phase 2)
- ❌ Full spell system beyond Shield & Love (Phase 2)
- ❌ Clans (Phase 3)
- ❌ Market/trading (Phase 2)
- ❌ Chat (Phase 3)

**Why this scope**:
- ✅ Testable single-player loop
- ✅ Verify economy balances
- ✅ Foundation for PvP/clans
- ✅ Can ship MVP and iterate

---

## Decision 7: Skill Cap Confirmation

**Status**: ✅ RESOLVED

**Question**:
> Confirm skill levels are capped at 10, not higher with expansions.

**Decision**:
### **Hard Cap: Level 10 Across All Skills** ✅

**Spec**:
```
For each action:
- Level range: [0, 10]
- Once at 10, further XP has no effect
- Progression bonus: +20% per level
  - Level 0: 0% bonus (1x multiplier)
  - Level 5: 100% bonus (2x multiplier)
  - Level 10: 200% bonus (3x multiplier)
```

**Examples**:
```
Exploration base: 50 land/tick
- Level 0: 50 * 1.0 = 50
- Level 5: 50 * 2.0 = 100
- Level 10: 50 * 3.0 = 150
- Level 11 (impossible): Still 150
```

---

## Decision 8: Tick Box Mechanic

**Status**: ✅ RESOLVED

**Question**:
> How do Tick Box (tick storage) and Max Ticks (active ticks) interact?

**Decision**:
### **Two-Tier Tick System** ✅

**Spec** (from copilot-instructions):
```
active_cap: 500         # Max ticks available for actions
tick_box_cap: 200       # Max stored "overflow" ticks

Regen flow:
1. If ticks < 500: Gain 1 tick every 5 minutes
2. If ticks >= 500 and tick_box < 200:
   - Tick regenerates into tick_box instead
   - Gives player "bank" for later
3. If both are full: Stop regenerating
```

**Examples**:
```
Scenario 1: Normal play
- Ticks: 287/500
- Regen: +1 tick/5min (standard)
- After 5 min: Ticks = 288/500

Scenario 2: Inactive (ticks full)
- Ticks: 500/500
- Tick Box: 0/200
- Regen: Gives to tick box instead
- After 5 min: Ticks = 500/500, Tick Box = 1/200

Scenario 3: Using box
- Ticks: 450/500
- Tick Box: 150/200
- Player spends 200 ticks (goes to -1, pulls from box)
- Ticks: 250/500, Tick Box: 149/200
```

**Why this system**:
- ✅ Prevents gated progress (no max cap waste)
- ✅ Encourages regular play
- ✅ Rewards dedicated players (tick box)
- ✅ Prevents multi-week idle gain

---

## Decision 9: Determinism Requirement

**Status**: ✅ RESOLVED

**Question**:
> How strict is determinism? Do random values matter?

**Decision**:
### **Determinism Required; Random Values Must Be Seeded** ✅

**Spec**:
```
Same input state → Same output state (always)

For random values:
- Seed RNG with deterministic value (e.g., empire ID + timestamp)
- Love spell: random(4, 10) always same for given empire + tick
- Not cryptographically random, just reproducible
```

**Why needed**:
- ✅ Audit logging (did spell do what it claimed?)
- ✅ Dispute resolution (player says spell bugged)
- ✅ Regression testing (formula changes work correctly)
- ✅ Encryption verification (can replay state)

**Code pattern**:
```go
func CastLoveSpell(empire *Empire, seed int64) int {
    rng := rand.NewSource(empire.ID + seed)
    randomBonus := rng.Intn(7) + 4  // 4-10 range
    totalMorale := min(100, empire.Morale + 2 + randomBonus)
    return totalMorale
}
```

---

## Decision 10: API Contract Priorities

**Status**: ✅ RESOLVED

**Question**:
> Which endpoints are Phase 1 vs later?

**Decision**:
### **Phase 1 Endpoints** ✅

**REQUIRED** (Phase 1):
```
Auth:
  POST /api/auth/login
  POST /api/auth/register
  POST /api/auth/logout

Empire:
  GET /api/empire              # Get current empire state
  POST /api/empire             # Create new empire
  PUT /api/empire              # Update (tax rate, troop distribution)

Actions:
  POST /api/action/explore
  POST /api/action/meditate
  POST /api/action/drill
  POST /api/action/farm

Spells:
  POST /api/spell/love
  POST /api/spell/shield      # Cast/activate shield
```

**NOT Phase 1**:
- ❌ Attack/raid endpoints (Phase 2)
- ❌ Chat/clan endpoints (Phase 3)
- ❌ Market/trade endpoints (Phase 2)
- ❌ Spell search/intelligence (Phase 2)

---

## Summary Table

| Decision | Status | Owner | Impact |
|----------|--------|-------|--------|
| Action Progression Persistence | ✅ Persist across empires | Game Design | HIGH |
| Streak Decay Formula | ✅ 0.85^n with 0.25 floor | Balance | MEDIUM |
| Spell Shield Type | ✅ Mitigation (0.33x) | Game Design | HIGH |
| Tome Tower Defense | ✅ Confirm: filled = tomes/100 | Engine | CRITICAL |
| Morale Damage Scope | ✅ All damage types | Engine | HIGH |
| Phase 1 Scope | ✅ Single-player loop | Product | CRITICAL |
| Skill Cap | ✅ Level 10 hard cap | Engine | MEDIUM |
| Tick Box Mechanics | ✅ 500 active + 200 storage | Engine | MEDIUM |
| Determinism | ✅ Required, seeded RNG | Architecture | CRITICAL |
| API Priorities | ✅ Core actions only | Backend | MEDIUM |

---

## Next Steps for Implementation

1. **Backend developers**: Reference this for skill/action/streak implementation
2. **Frontend developers**: Confirm progress persistence in UI (account-level skills)
3. **Database**: Update schema to move action_progression to player level
4. **Tests**: Verify all decisions with test cases from VERIFY_INVARIANTS.md

---

## Questions or Additions?

If new decisions arise during implementation:
1. Document question clearly
2. Propose options with rationale
3. Reference game design goals
4. Add to this file once approved

See [SETUP.md](SETUP.md) for how to flag ambiguities during development.
