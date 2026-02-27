# React Component Structure

Mobile-first React architecture for the Management Dashboard.

## Folder Structure

```
src/
├── components/
│   ├── screens/
│   │   └── ManagementScreen.tsx          # Main dashboard container
│   ├── sections/
│   │   ├── EmpireHeader.tsx              # Sticky header with empire name/status
│   │   ├── ResourcesSection.tsx          # Resources & per-tick rates
│   │   ├── BuildingsSection.tsx          # Building counts & production
│   │   ├── CivilizationSection.tsx       # Civilians, tax rate control
│   │   ├── ActionProgressionSection.tsx  # Skill levels & experience
│   │   └── IntelligenceReportSection.tsx # Intelligence spell results
│   ├── common/
│   │   ├── ResourceRow.tsx               # Reusable resource display
│   │   ├── StatRow.tsx                   # Reusable stat display
│   │   ├── ProgressBar.tsx               # Experience/progress bar
│   │   ├── Slider.tsx                    # Custom slider control
│   │   ├── StatusBadge.tsx               # Status indicator (surplus/deficit/etc)
│   │   ├── CollapsibleSection.tsx        # Collapsible container
│   │   └── Spinner.tsx                   # Loading state
│   └── layout/
│       └── ScreenLayout.tsx              # Responsive grid layout
├── hooks/
│   ├── useEmpireState.ts                 # Empire data fetching & caching
│   ├── useResourceManager.ts             # Resource calculations
│   ├── useDebounce.ts                    # Debounced API calls
│   └── useMobileLayout.ts                # Mobile-specific layout logic
├── types/
│   ├── empire.ts                         # Interfaces for empire data
│   ├── resources.ts                      # Resource type definitions
│   └── ui.ts                             # UI state interfaces
├── services/
│   ├── empireAPI.ts                      # API calls to Go backend
│   ├── spells.ts                         # Spell-related API calls
│   └── calculations.ts                   # Game logic calculations
├── utils/
│   ├── formatting.ts                     # Number/percentage formatting
│   ├── colors.ts                         # Status color mapping
│   └── storage.ts                        # localStorage helpers
└── styles/
    ├── tailwind.config.js                # Tailwind configuration
    ├── variables.css                     # CSS custom properties
    └── global.css                        # Global styles
```

---

## Key Components

### 1. ManagementScreen (Container)

**Purpose:** Orchestrates data fetching and section rendering.

```typescript
// components/screens/ManagementScreen.tsx
interface ManagementScreenProps {}

export const ManagementScreen: React.FC<ManagementScreenProps> = () => {
  const { empire, loading, error, refetch } = useEmpireState();
  const { currentResources, perTickDelta, status } = useResourceManager(empire);

  if (loading) return <Spinner />;
  if (error) return <ErrorBoundary error={error} />;

  return (
    <ScreenLayout>
      <EmpireHeader empire={empire} />
      <ResourcesSection
        resources={currentResources}
        delta={perTickDelta}
        status={status}
      />
      <BuildingsSection buildings={empire.buildings} />
      <CivilizationSection
        civilization={empire.civilization}
        onTaxRateChange={handleTaxRateChange}
      />
      <ActionProgressionSection actions={empire.action_progression} />
    </ScreenLayout>
  );
};
```

**Data Flow:**
- Fetches empire state on mount and at interval (e.g., 5 seconds)
- Passes immutable state down to children
- Callback props handle user interactions (tax rate, troop distribution)
- Debounced API calls on changes

---

### 2. ResourcesSection

**Purpose:** Display current resources with per-tick deltas and status.

```typescript
// components/sections/ResourcesSection.tsx
interface ResourcesSectionProps {
  resources: ResourceState;
  delta: ResourceDelta;
  status: ResourceStatus;
}

export const ResourcesSection: React.FC<ResourcesSectionProps> = ({
  resources,
  delta,
  status,
}) => {
  return (
    <CollapsibleSection title="Resources" defaultOpen>
      <ResourceRow
        label="Gold"
        current={resources.gold}
        perTick={delta.gold}
        status={status.gold}
      />
      <ResourceRow
        label="Grain"
        current={resources.food}
        perTick={delta.food}
        status={status.food}
      />
      <ResourceRow
        label="Tomes"
        current={resources.tomes}
        perTick={delta.tomes}
        ratio={calculateTomeRatio(resources.tomes, resources.meditation_towers)}
        ratioStatus={calculateTomeRatioStatus(ratio)}
      />
      <ResourceRow
        label="Troops"
        current={resources.troops}
        perTick={delta.troops}
        breakdown={resources.troop_breakdown}
      />
      <ResourceRow
        label="Ticks"
        current={`${resources.ticks} / 500`}
        regenTime={getNextTickTime()}
      />
    </CollapsibleSection>
  );
};
```

---

### 3. ResourceRow (Reusable Component)

**Purpose:** Display a single resource with flexible content.

```typescript
// components/common/ResourceRow.tsx
interface ResourceRowProps {
  label: string;
  current: number | string;
  perTick?: number;
  status?: 'surplus' | 'deficit' | 'neutral' | 'optimal';
  ratio?: string;
  ratioStatus?: 'optimal' | 'warning' | 'critical';
  breakdown?: Record<string, number>;
  regenTime?: string;
}

export const ResourceRow: React.FC<ResourceRowProps> = ({
  label,
  current,
  perTick,
  status,
  ratio,
  ratioStatus,
  breakdown,
  regenTime,
}) => {
  return (
    <div className="resource-row">
      <div className="resource-label">{label}</div>

      <div className="resource-value">
        {formatNumber(current)}
      </div>

      {perTick !== undefined && (
        <div className={`resource-delta ${perTick >= 0 ? 'positive' : 'negative'}`}>
          {perTick >= 0 ? '+' : ''}{formatNumber(perTick)}/tick
        </div>
      )}

      {status && <StatusBadge status={status} />}

      {ratio && (
        <div className={`resource-ratio ${ratioStatus}`}>
          Ratio: {ratio}
        </div>
      )}

      {breakdown && (
        <div className="resource-breakdown">
          {Object.entries(breakdown).map(([type, count]) => (
            <div key={type} className="breakdown-item">
              {type}: {formatNumber(count)}
            </div>
          ))}
        </div>
      )}

      {regenTime && (
        <div className="resource-regen">
          Next tick in {regenTime}
        </div>
      )}
    </div>
  );
};
```

---

### 4. CivilizationSection

**Purpose:** Display civilians and tax rate control.

```typescript
// components/sections/CivilizationSection.tsx
interface CivilizationSectionProps {
  civilization: CivilizationState;
  onTaxRateChange: (rate: number) => void;
}

export const CivilizationSection: React.FC<CivilizationSectionProps> = ({
  civilization,
  onTaxRateChange,
}) => {
  const [taxRate, setTaxRate] = useState(civilization.tax_rate);
  const debouncedChange = useDebounce((rate: number) => {
    onTaxRateChange(rate);
  }, 500);

  const handleSliderChange = (newRate: number) => {
    setTaxRate(newRate);
    debouncedChange(newRate);
  };

  const getTaxRateInfo = (rate: number) => {
    if (rate < 9) return { label: 'Low Tax', status: 'bonus', effect: '+25% growth' };
    if (rate < 23) return { label: 'Normal Tax', status: 'neutral', effect: 'No modifier' };
    if (rate < 35) return { label: 'High Tax', status: 'warning', effect: '-15% to -35% growth' };
    if (rate < 100) return { label: 'Exodus', status: 'critical', effect: 'Civilians fleeing' };
    return { label: 'Maximum Tax', status: 'critical', effect: '-7% loss/tick' };
  };

  const taxInfo = getTaxRateInfo(taxRate);

  return (
    <CollapsibleSection title="Civilization & Economy">
      <StatRow
        label="Civilians"
        value={formatNumber(civilization.civilians)}
        secondary={`Cap: ${formatNumber(civilization.cap)}`}
        delta={civilization.growth_per_tick}
      />

      <div className="tax-control">
        <label>Tax Rate: {taxRate}%</label>
        <Slider
          min={0}
          max={100}
          value={taxRate}
          onChange={handleSliderChange}
          className="tax-slider"
        />
        <div className={`tax-info ${taxInfo.status}`}>
          <div className="tax-label">{taxInfo.label}</div>
          <div className="tax-effect">{taxInfo.effect}</div>
        </div>
      </div>

      <StatRow
        label="Gold/Tick from Tax"
        value={`+${formatNumber(civilization.gold_per_tick_from_tax)}`}
      />
    </CollapsibleSection>
  );
};
```

---

### 5. BuildingsSection

**Purpose:** Display building counts and troop distribution control.

```typescript
// components/sections/BuildingsSection.tsx
interface BuildingsSectionProps {
  buildings: BuildingState;
  onTroopDistributionChange?: (distribution: TroopDistribution) => void;
}

export const BuildingsSection: React.FC<BuildingsSectionProps> = ({
  buildings,
  onTroopDistributionChange,
}) => {
  return (
    <CollapsibleSection title="Buildings & Production">
      <StatRow label="Farms" value={buildings.farms} secondary="10 grain/tick" />
      <StatRow label="Barracks" value={buildings.barracks} secondary="15 troops/tick" />
      <StatRow label="Towers" value={buildings.towers} secondary="100 def/building" />
      <StatRow label="Bazaars" value={buildings.bazaars} secondary="25 gold/tick" />
      <StatRow label="Meditation Towers" value={buildings.meditation_towers} secondary="15 tomes/tick" />

      {onTroopDistributionChange && (
        <TroopDistributionControl onChange={onTroopDistributionChange} />
      )}
    </CollapsibleSection>
  );
};
```

---

### 6. Custom Hooks

**useEmpireState** - Fetches and caches empire data

```typescript
// hooks/useEmpireState.ts
interface UseEmpireStateReturn {
  empire: EmpireState;
  loading: boolean;
  error: Error | null;
  refetch: () => Promise<void>;
}

export const useEmpireState = (
  refreshIntervalMs = 5000
): UseEmpireStateReturn => {
  const [empire, setEmpire] = useState<EmpireState | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    const fetchEmpire = async () => {
      try {
        const data = await empireAPI.getEmpire();
        setEmpire(data);
        setError(null);
      } catch (err) {
        setError(err as Error);
      } finally {
        setLoading(false);
      }
    };

    fetchEmpire();
    const interval = setInterval(fetchEmpire, refreshIntervalMs);

    return () => clearInterval(interval);
  }, [refreshIntervalMs]);

  const refetch = async () => {
    try {
      const data = await empireAPI.getEmpire();
      setEmpire(data);
    } catch (err) {
      setError(err as Error);
    }
  };

  return { empire: empire!, loading, error, refetch };
};
```

**useResourceManager** - Calculates derived resource values

```typescript
// hooks/useResourceManager.ts
export const useResourceManager = (empire: EmpireState) => {
  const currentResources = useMemo(() => ({
    gold: empire.resources.gold,
    food: empire.resources.food,
    tomes: empire.resources.tomes,
    troops: empire.resources.troops,
    ticks: empire.resources.ticks,
    civilians: empire.resources.civilians,
  }), [empire.resources]);

  const perTickDelta = useMemo(() => ({
    gold: calculateGoldPerTick(empire),
    food: calculateFoodPerTick(empire),
    tomes: calculateTomesPerTick(empire),
    troops: calculateTroopsPerTick(empire),
  }), [empire]);

  const status = useMemo(() => ({
    gold: perTickDelta.gold >= 0 ? 'surplus' : 'deficit',
    food: perTickDelta.food >= 0 ? 'surplus' : 'deficit',
    tomes: 'neutral',
    troops: 'neutral',
  }), [perTickDelta]);

  return { currentResources, perTickDelta, status };
};
```

**useDebounce** - Debounce function calls

```typescript
// hooks/useDebounce.ts
export const useDebounce = <T extends (...args: any[]) => any>(
  callback: T,
  delay: number
) => {
  const timeoutRef = useRef<NodeJS.Timeout | null>(null);

  return useCallback(
    (...args: Parameters<T>) => {
      if (timeoutRef.current) clearTimeout(timeoutRef.current);
      timeoutRef.current = setTimeout(() => {
        callback(...args);
      }, delay);
    },
    [callback, delay]
  );
};
```

---

## TypeScript Interfaces

```typescript
// types/empire.ts

export interface EmpireState {
  user_id: string;
  empire: {
    name: string;
    created_at: number;
    status: 'alive' | 'dead';
    resources: ResourceState;
    action_progression: ActionProgressionState;
    buildings: BuildingState;
    streaks: StreakState;
    clan_id?: string;
    last_attacked?: number;
  };
  stars: number;
}

export interface ResourceState {
  land: number;
  tomes: number;
  troops: number;
  food: number;
  ticks: number;
  civilians: number;
  gold: number;
  tax_rate: number;
}

export interface ActionProgressionState {
  [action: string]: {
    level: number;
    experience: number;
    last_action_timestamp: number;
  };
}

export interface BuildingState {
  farms: number;
  barracks: number;
  towers: number;
  bazaars: number;
  meditation_towers: number;
}

export type ResourceStatus = 'surplus' | 'deficit' | 'neutral' | 'optimal';

export interface ResourceDelta {
  gold: number;
  food: number;
  tomes: number;
  troops: number;
}
```

---

## Styling Approach

Use **Tailwind CSS** with custom CSS variables for minimalist design:

```css
/* styles/variables.css */
:root {
  --color-bg-primary: #ffffff;
  --color-bg-secondary: #fafafa;
  --color-text-primary: #333333;
  --color-text-secondary: #666666;
  --color-border: #eeeeee;
  --color-surplus: #4caf50;
  --color-deficit: #f44336;
  --color-optimal: #2196f3;
  --color-neutral: #9e9e9e;
  --color-warning: #ff9800;

  --spacing-xs: 8px;
  --spacing-sm: 12px;
  --spacing-md: 16px;
  --spacing-lg: 24px;

  --font-size-sm: 12px;
  --font-size-base: 14px;
  --font-size-lg: 16px;
  --font-size-xl: 18px;
  --font-size-2xl: 24px;
}

/* Mobile-first breakpoints */
@media (min-width: 768px) {
  /* Tablet/Web styles */
}
```

---

## State Management Pattern

This architecture uses:
1. **React hooks** for component-level state (`useState`)
2. **Custom hooks** for shared logic (`useEmpireState`, `useResourceManager`)
3. **useCallback** + **debouncing** for optimized API calls
4. **Context API** (optional) if you need global state (user auth, theme, etc.)

**NO Redux/Zustand required** for this dashboard—hooks are sufficient.

---

## Performance Considerations

- Memoize expensive calculations with `useMemo`
- Use `useCallback` for stable function references
- Debounce rapid API calls (tax rate slider, troop distribution)
- Virtual scrolling (optional) if section list becomes very long
- Lazy load intelligence report section only when needed

