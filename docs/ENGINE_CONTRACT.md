
# Engine Contract

## Action Progression
When a player performs an action, they gain experience toward that action's level. At each full level, a bonus multiplier is applied.

output = base_output * (1 + (multiplier_per_level * current_level))

Where:
- base_output = the action's raw output (e.g., tomes_per_tick)
- multiplier_per_level = 0.2 (20% bonus per level)
- current_level = current level for that action (0-10)

## Tome Tower Defense
filled_towers = min(towers, floor(tomes / 100))
defense = filled_towers * 100

## Civilians & Taxation

### Civilian Growth
Civilians automatically generate based on empty land (land without troops).
- **Civilian cap**: land * 1000
- **Base generation**: Gains 1 civilian per empty acre per tick (modified by tax rate)

### Tax Rate Multiplier
Tax rate affects civilian generation speed:

**Low tax (0-9%)**:
- multiplier = 1.25 (+25% boost)

**Normal tax (10-22%)**:
- multiplier = 1.0 (no bonus or penalty)

**High tax (23-34%)**:
- Pro-rated penalty from -15% to -35%
- multiplier = 1.0 - (0.20 * ((tax_rate - 23) / 12))
- At 23%: multiplier = 1.0 - 0 = 1.0
- At 35%: multiplier = 1.0 - 0.20 = 0.8

**Exodus threshold (35-99%)**:
- Civilians start leaving (negative growth)
- multiplier = 1.0 - (0.20 * ((tax_rate - 23) / 12))
- At 35%: -0% (threshold point, no growth/loss)
- Continues scaling toward 100%

**Total exodus (100%)**:
- civilians_lost_per_tick = current_civilians * 0.07
- All civilians flee regardless of land

### Final Civilian Growth Formula
```
change_per_tick = empty_acres * 1.0 * multiplier
new_civilians = min(land * 1000, current_civilians + change_per_tick)
```

### Gold Income
```
gold_per_tick = civilians * (tax_rate / 100)
```

## Buildings

### Construction
Players construct buildings on empty land. Construction consumes resources and time.

**Constraints:**
- Cost per building: 10 gold + 5 grain (uniform across all types)
- Build rate: 3% of total land can be built per tick
- Example: 1000 land can build floor(1000 * 0.03) = 30 buildings per tick maximum

**Formula:**
```
buildings_constructable_per_tick = floor(total_land * 0.03)
```

### Building Types & Production

**Farms**
- Production: 10 grain per building per tick
- Modifier: farm_level (action progression level 0-10)
- `grain_per_tick = num_farms * 10 * (1 + (0.2 * farm_level))`

**Barracks**
- Production: 15 troops per building per tick
- Modifier: drill_level (action progression level 0-10)
- `troops_per_tick = num_barracks * 15 * (1 + (0.2 * drill_level))`

**Towers**
- Production: 100 defense points per building (no modifier)
- Defense scales directly with count
- `total_defense = num_towers * 100 + (tome_tower_defense)`

**Bazaars**
- Production: 25 gold per building per tick
- Modifier: tax_rate_based (affects local market multiplier)
- `gold_per_tick = num_bazaars * 25 * tax_rate_modifier`

**MeditationTowers**
- Production: 15 tomes per building per tick
- Modifier: meditation_level (action progression level 0-10)
- Affects spell strength via tome ratio system
- `tomes_per_tick = num_meditation_towers * 15 * (1 + (0.2 * meditation_level))`

### Meditation Tower Tome Ratio System

Spell strength depends on the ratio of current tomes to meditation towers.

**Optimal ratio: 125 tomes per tower (0% modifier)**

**Above optimal (125-175 tomes/tower)**:
- Bonus ramps from +15% to +20%
- `spell_strength_modifier = 1.0 + (0.15 + (0.001 * (ratio - 125)))`
- At 125:1 → modifier = 1.15 (+15%)
- At 175:1 → modifier = 1.20 (+20%)

**Exodus (175+ tomes/tower)**:
- 3% of tomes lost per tick
- `tomes_lost_per_tick = current_tomes * 0.03`

**Below optimal (0-125 tomes/tower)**:
- No penalty
- Modifier = 1.0 (0%)
- Spells fail if required tome ratio is not met

### Spell Minimum Tome Ratios

For a spell to execute, caster must have sufficient tomes relative to meditation towers:

| Spell Type | Minimum Ratio |
|---|---|
| Spell Shield | 15:1 |
| Non-offensive actions | 20:1 |
| Intelligence spells | 30:1 |
| Magic attacks | 60:1 |

**Check:**
```
required_tomes = meditation_towers * minimum_ratio
if current_tomes < required_tomes:
  spell_fails()
```

## Spell Shield
If active: spell_effect *= mitigation_factor (0.33)
Applies to spell effects only.

## Morale System

Morale tracks empire psychological state and affects combat effectiveness.

### Morale Basics
- **Range**: 0-100 (percentage)
- **Starting value**: 100
- **Multiplier on attacks**: `attack_strength = base_damage * (morale / 100)`

### Morale Changes

**Offensive actions** (cost 1% morale each):
- Military strike
- Magic attack spells
- Any spell dealing damage to other players

Example: After 5 attacks, morale = 95% (loses 5%)

**Non-offensive actions** (gain 1% morale each, capped at 100):
- Explore
- Meditate
- Drill (training troops)
- Farm
- Build structures
- Economic actions

Example: After 3 non-offensive actions with 85% morale → morale = 88%

### Morale Formula

```
if action is offensive:
  new_morale = max(0, current_morale - 1)

if action is non_offensive:
  new_morale = min(100, current_morale + 1)
```

### Attack Strength Multiplier

```
damage_output = base_damage * (current_morale / 100)

Examples:
- 100% morale = 100% damage
- 90% morale = 90% damage
- 50% morale = 50% damage
- 0% morale = 0% damage (cannot attack)
```

---

## Love Spell

A restorative morale spell that requires meditation.

### Requirements
- **Tick cost**: 2 ticks
- **Tome requirement**: 20:1 ratio (20 tomes per meditation tower minimum)
- **Cannot be used in combat** (it's non-offensive)

### Morale Restoration

Morale increases from two sources:

1. **Tick cost itself** (non-offensive action): +2% morale
2. **Love spell effect**: +4 to +10 morale points (randomized)

**Total morale gain**: 6 to 12 points

```
love_spell_bonus = random(4, 10)  // Variable effect
tick_bonus = 2                     // From 2 ticks spent
total_morale_gain = love_spell_bonus + tick_bonus

Formula:
new_morale = min(100, current_morale + total_morale_gain)

Examples:
- At 85% morale, love spell bonus = 5 → new morale = min(100, 85 + 2 + 5) = 92%
- At 95% morale, love spell bonus = 8 → new morale = min(100, 95 + 2 + 8) = 100% (capped)
- At 88% morale, love spell bonus = 4 → new morale = min(100, 88 + 2 + 4) = 94%
```

### Balance Notes

- Love spell is the primary way to recover morale when heavily damaged by combat
- Players must choose between aggressive warfare (morale decay) and morale recovery (2 ticks cost)
- Non-offensive actions naturally rebuild morale at 1% per action
- Love spell is valuable for players planning extended PvP campaigns

