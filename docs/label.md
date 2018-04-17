### AddLabel
This command adds a label to a bamboo build. 

Command Input: 

```
{
    "build": "ABC-123",
    "label": "staging"
}
```

### Output Events
This command returns either a `LabelAdded` or `FailedToAddLabel` event


### LabelAdd Event Output 

```
{
    "build": "ABC-123",
    "label": "staging"
}
```

### FailedToAddLabel Event Output

```
{
    "build": "ABC-123",
    "label": "staging"
    "error": "some error occured"
}
```

### DeleteLabel
This command deletes a label to a bamboo build. 

Command Input: 

```
{
    "build": "ABC-123",
    "label": "staging"
}
```

### Output Events
This command returns either a `LabelDeleted` or `FailedToDeleteLabel` event

### LabelDeleted Event Output 

```
{
    "build": "ABC-123",
    "label": "staging"
}
```

### FailedToDeleteLabel Event Output

```
{
    "build": "ABC-123",
    "label": "staging"
     "error": "some error occured"   
}
```
