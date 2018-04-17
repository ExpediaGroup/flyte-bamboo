### GetLifeCycle command
This command adds a comment to a bamboo build.


#### Input for GetLifeCycle
```
{
    "build": "ABC-DEF-123",
    "stage": "prod-verify"
}
```

#### Output Events
This command returns either a `GetLifeCycle` or `GetLifeCycleFailed` event

### GetLifeCycle Event Output

```
{
   "Stage":"stage"
   "Build":"build"
   "LifeCycle":"lifecycle"
   "State": "state"
}
```

### GetLifeCycleFailed Event Output 

```
{
    "build": "ABC-DEF-123",
    "stage": "staging"
    "error" : "some error occured"
}
```