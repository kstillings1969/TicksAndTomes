# Development Standards

## Internationalization (i18n)

### Requirement
All text displayed in the UI must use language strings rather than hardcoded strings. This enables supporting multiple languages without code changes.

### Why
- Allows adding new languages in the future without modifying application code
- Centralizes all UI text in language files for easy translation management
- Separates concerns: developers maintain code structure, translators maintain text content

### Implementation

#### Use Language Keys, Not Hardcoded Strings

**❌ Incorrect:**
```typescript
// Frontend
return <button>{i18n.t('action.explore')}</button>

// Backend
res.json({ message: "Exploration completed successfully" })
```

**✅ Correct:**
```typescript
// Frontend
return <button>{i18n.t('action.explore')}</button>

// Backend
const response = {
  resourceKey: "messages.exploration_complete",
  data: { landGained: 50 }
}
res.json(response)
// Client renders: i18n.t('messages.exploration_complete', { landGained: 50 })
```

#### Naming Convention for Language Keys

Use dot notation with hierarchical organization:
- `actions.*` — Action names and descriptions (explore, meditate, drill, farm)
- `resources.*` — Resource names (land, tomes, troops, food, gold, civilians)
- `buildings.*` — Building names and descriptions (farms, barracks, towers, bazaars, meditation_towers)
- `messages.*` — Game messages and notifications
- `labels.*` — UI labels (Level, Experience, Tax Rate, etc.)
- `errors.*` — Error messages
- `ui.*` — UI elements (buttons, menus, dialogs)

#### Example Language File Structure

```json
{
  "actions": {
    "explore": "Explore",
    "explore_description": "Scout new land for your empire",
    "meditate": "Meditate",
    "meditate_description": "Generate tomes to cast spells",
    "drill": "Drill",
    "drill_description": "Train troops for defend and attack",
    "farm": "Farm",
    "farm_description": "Cultivate food for your civilians"
  },
  "resources": {
    "land": "Land",
    "tomes": "Tomes",
    "troops": "Troops",
    "food": "Food",
    "gold": "Gold",
    "civilians": "Civilians",
    "ticks": "Ticks"
  },
  "buildings": {
    "farms": "Farms",
    "farms_description": "Produce grain",
    "barracks": "Barracks",
    "barracks_description": "Train troops",
    "towers": "Towers",
    "towers_description": "Provide defense",
    "bazaars": "Bazaars",
    "bazaars_description": "Generate gold",
    "meditation_towers": "Meditation Towers",
    "meditation_towers_description": "Generate tomes"
  },
  "messages": {
    "exploration_complete": "You gained {landGained} land",
    "meditation_complete": "You gained {tomesGained} tomes",
    "insufficient_tomes": "Insufficient tomes to cast this spell",
    "spell_failed": "Spell cast failed"
  },
  "labels": {
    "level": "Level",
    "experience": "Experience",
    "tax_rate": "Tax Rate"
  }
}
```

#### When to Use Language Keys

1. **All UI text** — buttons, labels, headings, descriptions
2. **Error messages** — validation errors, action failures
3. **Game notifications** — successful actions, events, warnings
4. **Dynamic content** — use placeholders for variable data

#### When NOT to Use Language Keys

- Internal code comments
- Variable/function names
- API field names (keep these stable)
- Log messages (can be English for development/debugging)
- Database field names

### Best Practices

1. Use meaningful key names — avoid generic names like `string1`, `text2`
2. Keep related strings grouped in the same category
3. Use placeholders for dynamic values: `{variableName}`
4. Avoid context-dependent constructs; rely on the key structure instead
5. Document context in the language file if translation is ambiguous

### Future Language Support

When adding a new language:
1. Create new language file (e.g., `en.json`, `es.json`, `fr.json`)
2. Translate all keys in the file
3. Update i18n configuration to register new language
4. No code changes required
