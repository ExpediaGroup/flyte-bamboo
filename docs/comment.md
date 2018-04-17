### AddComment command
This command adds a comment to a bamboo build.


#### Input for AddComment
```
{
    "build": "ABC-123",
    "comment": "staging"
}
```

#### Output Events
This command returns either a `CommentAdded` or `FailedToAddComment` event

### CommentAdded Event Output

```
{
    "build": "ABC-123",
    "comment": "staging"
}
```

### FailedToAddComment Event Output 

```
{
    "build": "ABC-123",
    "comment": "staging"
    "error" : "some error occured"
}
```