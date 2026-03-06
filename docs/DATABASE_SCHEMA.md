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

  // Skill progression (ACCOUNT-LEVEL: persists across empire deaths)
  // Decision: Skills carry over to new empires (see DECISIONS_RESOLVED.md)
  "action_progression": {
    "meditate": {
      "level": number,           // 0-10 (hard cap)
      "experience": number,      // 0.0-1.0 (progress to next level)
      "streak_count": number,    // Consecutive actions (for streak decay)
      "last_action_timestamp": number
    },
    "drill": {
      "level": number,
      "experience": number,
      "streak_count": number,
      "last_action_timestamp": number
    },
    "farm": {
      "level": number,
      "experience": number,
      "streak_count": number,
      "last_action_timestamp": number
    },
    "explore": {
      "level": number,
      "experience": number,
      "streak_count": number,
      "last_action_timestamp": number
    }
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
      "ticks": number,          // 0-500 (active ticks)
      "tick_box": number,       // 0-200 (stored ticks for overflow)
      "civilians": number,      // Auto-generated from land (capped at land * 1000)
      "gold": number,           // Generated from tax_rate * civilians
      "tax_rate": number        // 0-100 (affects civilian generation speed)
    },

    // Buildings
    "buildings": {
      "farms": number,
      "barracks": number,
      "towers": number,
      "bazaars": number,
      "meditation_towers": number
    },

    // Morale system (affects attack strength)
    "morale": number,          // 0-100 (affects attack strength multiplier)

    // Spell states
    "spells": {
      "shield_active": boolean,    // Is Spell Shield active?
      "shield_active_until": "timestamp | null"
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
- Resources reset to 0
- **Action levels PERSIST** on account (see [DECISIONS_RESOLVED.md](DECISIONS_RESOLVED.md))
- Account `stars` remain (purchasable currency)

### Restart
- Player can create new empire with fresh `empire` object
- New `empire.name` can be set
- Account `stars` available for purchases
- **Action progression carries over** — Skills 0-10 maintained at account level
- New empire has 0 resources but same skill levels as old empire
- Encourages players to restart and experience new strategies
