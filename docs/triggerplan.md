### TriggerPlan
This command starts a plan in bamboo.

Command Input:
```
{
    "plan": "ABCDEF-GHI"
}
```
#### Output TriggerPlan Events
This command returns either a `PlanTriggered` or `FailedToTriggerPlan`


#### PlanTriggered output
```
{
    "plan": "ABCDEF-GHI"
}
```

#### FailedToTriggerPlan Event output
```
{
    "plan": "ABCDEF-GHI"
    "error": "err..."
}
```