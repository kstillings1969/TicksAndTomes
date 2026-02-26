
# Go Appendix (excerpt) — Engine Model & Tick Regen

```go
package game

import "time"

const (
    TickInterval = 5 * time.Minute
    MaxTicks     = int64(500)
    MaxTickBox   = int64(200)
)

type ActionKind string

const (
    ActionExplore        ActionKind = "explore"
    ActionBuild          ActionKind = "build"
    ActionFarm           ActionKind = "farm"
    ActionCash           ActionKind = "cash"
    ActionIndustry       ActionKind = "industry"
    ActionMeditate       ActionKind = "meditate"
    ActionMilitaryAttack ActionKind = "military_attack"
    ActionMilitaryRaid   ActionKind = "military_raid"
    ActionCastSpell      ActionKind = "cast_spell"
)

type SkillID string
type SkillProgress struct{ Level int32; XP int64 }
type SkillStreak struct{ Count int32; LastAward time.Time }
type SkillProfile struct {
    Skills  map[SkillID]SkillProgress
    Streaks map[SkillID]SkillStreak
}

type EmpireState struct {
    Ticks      int64
    TickBox    int64
    LastTickAt time.Time
    Skill      SkillProfile
}

func ApplyTickRegen(s *EmpireState, now time.Time) {
    if s.LastTickAt.IsZero() { s.LastTickAt = now; return }
    steps := int64(now.Sub(s.LastTickAt) / TickInterval)
    for i := int64(0); i < steps; i++ {
        if s.Ticks < MaxTicks {
            s.Ticks++
            if s.TickBox > 0 && s.Ticks < MaxTicks { s.Ticks++; s.TickBox-- }
        } else if s.TickBox < MaxTickBox {
            s.TickBox++
        }
    }
    s.LastTickAt = s.LastTickAt.Add(time.Duration(steps) * TickInterval)
}

func streakMultiplier(base, min float64, count int32) float64 {
    m := 1.0
    for i := int32(1); i < count; i++ { m *= base }
    if m < min { m = min }
    return m
}
