
---
description: "Adjust numbers only via RULESET.yaml (no code changes). Keep invariants intact."
---

Task:
Update **only** `rules/RULESET.yaml`. Do not modify engine code. [3](https://twodegrees1.sharepoint.com/teams/GlobalTechnologyTechnicalEnablement/Shared%20Documents/Partner%20Learning/Microsoft/GitHub%20Copilot/Exam%20Prep%20Documents%20-%20Nov%202025/GHCP%20Exam%20Prep%20Guide%20-%20%20Josh%20Highlighted.pdf?web=1)

Input:
${input:rule_change:What should be changed in the ruleset?}

Constraints:
- Preserve invariants in RULESET (do not remove them). [3](https://twodegrees1.sharepoint.com/teams/GlobalTechnologyTechnicalEnablement/Shared%20Documents/Partner%20Learning/Microsoft/GitHub%20Copilot/Exam%20Prep%20Documents%20-%20Nov%202025/GHCP%20Exam%20Prep%20Guide%20-%20%20Josh%20Highlighted.pdf?web=1)
- Spell Shield remains mitigation (not immunity). [3](https://twodegrees1.sharepoint.com/teams/GlobalTechnologyTechnicalEnablement/Shared%20Documents/Partner%20Learning/Microsoft/GitHub%20Copilot/Exam%20Prep%20Documents%20-%20Nov%202025/GHCP%20Exam%20Prep%20Guide%20-%20%20Josh%20Highlighted.pdf?web=1)
- Tome Tower defense remains based on filled towers (≥100 tomes). [3](https://twodegrees1.sharepoint.com/teams/GlobalTechnologyTechnicalEnablement/Shared%20Documents/Partner%20Learning/Microsoft/GitHub%20Copilot/Exam%20Prep%20Documents%20-%20Nov%202025/GHCP%20Exam%20Prep%20Guide%20-%20%20Josh%20Highlighted.pdf?web=1)

Required output:
- Minimal patch (diff-style)
- Explanation of gameplay impact
- Any test updates needed (but do not implement tests in this prompt)
