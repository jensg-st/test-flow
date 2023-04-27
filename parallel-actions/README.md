# Parallel Error Handling

This workflows run three functions in parallel and contains an error handling example. There are three error scenarios baked in. In general the event to trigger that workflow should look like: 

```json
{
    "specversion" : "1.0",
    "type" : "testevent",
    "source" : "mysource",
    "datacontenttype" : "application/json",
    "data" : {
        "name1": "Jens",
        "name2": "Hello",
        "name3": "World"
    }
}
```

## Error Scenarios

### Empty Name

```json
{
    "specversion" : "1.0",
    "type" : "testevent",
    "source" : "mysource",
    "datacontenttype" : "application/json",
    "data" : {
        "name1": "Jens",
        "name2": "Hello",
        "name3": ""
    }
}
```

### Timeout

```json
{
    "specversion" : "1.0",
    "type" : "testevent",
    "source" : "mysource",
    "datacontenttype" : "application/json",
    "data" : {
        "name1": "timeout",
        "name2": "Hello",
        "name3": "World"
    }
}
```

### Wrong Name

```json
{
    "specversion" : "1.0",
    "type" : "testevent",
    "source" : "mysource",
    "datacontenttype" : "application/json",
    "data" : {
        "name1": "John",
        "name2": "Hello",
        "name3": "World"
    }
}
```