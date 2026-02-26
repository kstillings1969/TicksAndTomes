
version: v5
includes:
  units: rules/UNITS.yaml
  attacks: rules/ATTACKS.yaml

progression:
  model: experience_based
  cap_level: 10
  streak_decay:
    enabled: true
    exponential_base: 0.85
    min_multiplier: 0.25
    reset_after_seconds: 3600

time_resource:
  name: ticks
  regen: 1_per_5_minutes
  active_cap: 500
  tick_box_cap: 200

removed_systems:
  ages: true
  eras: true
  portals: true

defense_rules:
  tome_towers:
    tomess_per_tower: 100
    defense_per_filled_tower: 100
    requires_full_staffing: true

spells_allowed:
  - spell_shield
  - arcane_expansion
