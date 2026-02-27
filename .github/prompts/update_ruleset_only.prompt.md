
---
description: "Adjust numbers only via RULESET.yaml (no code changes). Keep invariants intact."
---

Task:
Update **only** `rules/RULESET.yaml`. Do not modify engine code. [3]

Input:
${input:rule_change:What should be changed in the ruleset?}

Constraints:
- Preserve invariants in RULESET (do not remove them). [3]
- Spell Shield remains mitigation (not immunity). [3]
- Tome Tower defense remains based on filled towers (≥100 tomes). [3]

Required output:
- Minimal patch (diff-style)
- Explanation of gameplay impact
- Any test updates needed (but do not implement tests in this prompt)
