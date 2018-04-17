### StartStage

This command starts a stage in bamboo. Key/value pairs passed in `vars` will be made available as bamboo inject variables.

Repeated attempts at triggering the command can be configured by setting values for:

`maxRetry` - The maximum number of attempts to make

`delay` - The time (in seconds) between attempts

Command Input: 
```
{
    "build": "ABC-DEF-123",
    "stage" : "initiate-staging",
    "maxRetry" : 180,
    "delay" : 20,
    "vars": {
        "deployStatus": "success"
    }
}
```

#### Output StartStage Event 
This command returns either a `StartStageSuccess` or `FailedToStartStage`

###  StartStageSuccess Event Output 

```
{
    "build": "ABC-DEF-123",
    "stage" : "initiate-staging",
    "vars": {
        "deployStatus": "success"
    }
}
```

###  FailedToStartStage Event Output 

```
{
    "build": "ABC-DEF-123",
    "stage" : "initiate-staging"
    "error: : "some error occured"
    "vars": {
        "deployStatus": "success"
    }
}
```

