### EnablePlan command
This command enables a plan in bamboo.

#### Input for EnablePlan

```
{
    "plan": "ABCDEF-GHI"
}
```

### Output Events
This command returns either a `PlanEnabled` or `FailedToEnablePlan` event

### PlanEnabled Event Output

```
{
    "plan": "ABC-123",
}
```

### FailedToAddComment Event Output 

```
{
    "plan": "ABC-123",
    "error" : "some error occured"
}
```