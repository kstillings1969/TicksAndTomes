# Database Schema

Document-based database structure for storing player state across empires, wars, and death mechanics.

## Player Document

```json
{
  "user_id": "string (unique)",
  "email": "string",

  // Account-level (persists across empire deaths/restarts)
  "stars": number,
  "total_empires_created": number,
  "all_time_stats": {
    "empires_killed": number,
    "invasions_participated": number,
    "resource_destroyed": number
  },

  // Current empire (resets on death at 0 land)
  "empire": {
    "name": "string",
    "created_at": "timestamp",
    "status": "alive" | "dead",

    // Resources
    "resources": {
      "land": number,           // 0 = dead/eliminated
      "tomes": number,
      "troops": number,
      "food": number,
      "ticks": number,
      "civilians": number,       // Auto-generated from land (capped at land * 1000)
      "gold": number,            // Generated from tax_rate * civilians
      "tax_rate": number         // 0-100 (affects civilian generation speed)
    },

    // Action progression
    "action_progression": {
      "meditate": {
        "level": number,        // 0-10
        "experience": number,   // 0.0-1.0 (progress to next level)
        "last_action_timestamp": number
      },
      "drill": {
        "level": number,
        "experience": number,
        "last_action_timestamp": number
      },
      "farm": {
        "level": number,
        "experience": number,
        "last_action_timestamp": number
      },
      "explore": {
        "level": number,
        "experience": number,
        "last_action_timestamp": number
      }
    },

    // Buildings
    "buildings": {
      "farms": number,
      "barracks": number,
      "towers": number,
      "bazaars": number,
      "meditation_towers": number
    },

    // Streak tracking (for decay calculation per copilot-instructions)
    "streaks": {
      "meditate": number,
      "drill": number,
      "farm": number,
      "explore": number
    },

    // Clan & combat
    "clan_id": "string | null",
    "last_attacked": "timestamp | null",
    "attack_history": [
      {
        "attacker_user_id": "string",
        "attacker_clan_id": "string | null",
        "damage": {
          "land_lost": number,
          "troops_lost": number,
          "tomes_lost": number,
          "food_lost": number
        },
        "timestamp": "timestamp"
      }
    ]
  }
}
```

## Collection Indexes

For efficient querying:

- `user_id` (unique)
- `email` (unique)
- `empire.name` (to support leaderboards)
- `empire.status` (to find active/dead players)
- `empire.clan_id` (to find clan members)
- `empire.created_at` (for time-based queries)
- `empire.resources.gold` (for economic rankings)
- `empire.resources.civilians` (for population queries)
- `stars` (for ranking)

## Empire Lifecycle

### Creation
- New player creates empire
- `empire.name` set by player
- `empire.status = "alive"`
- All resources start at 0
- Action levels start at 0
- Account `stars` from purchases carry over

### Active Play
- Player performs actions (explore/meditate/drill/farm)
- Action experience accumulates per `RULESET.yaml` (0.01 per action)
- Resources increase via multipliers based on action levels
- Streaks tracked for decay calculation

### Death (Land = 0)
- Triggered when `empire.resources.land` reaches 0 (via coordinated attacks)
- `empire.status = "dead"`
- Resources remain at 0
- Action levels & experience: **[DECISION PENDING]** — persist or reset?
- Account `stars` remain (purchasable currency)

### Restart
- Player can create new empire with fresh `empire` object
- New `empire.name` can be set
- Account `stars` available for purchases
- Previous action levels: **[DEPENDS ON ABOVE DECISION]**

## Empire Death Decision Point

**Question: Should action progression persist across empire deaths?**

- **Option A: Reset on death** — Each new empire starts fresh at level 0. Encourages replayability, but resets skill progression.
- **Option B: Persist across empires** — Keep action levels on account. Makes returning players powerful faster, but may reduce difficulty curve for experienced players restarting.

This affects balancing significantly and should be decided before implementation.
