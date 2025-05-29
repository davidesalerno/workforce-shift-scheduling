# workforce-shift-scheduling
A simple workforce shift scheduling app

# Introduction
The purpose of this application is to facilitate a workforce shift scheduling problem using a greedy algorithm.

The scheduler input are the details about the workers and the shifts and the output it will be the proposed scheduling.

In order to interact with the application you can use the REST API described into the example.

# Example

Endpoint: POST http://localhost:8080/schedule

```
{
  "shifts": [
    {"date":"2025-07-11","slot":"morning","required":2},
    {"date":"2025-07-11","slot":"evening","required":1}
  ],
  "workers": [
    {"id":"w1","name":"Alice","max_shifts":2,"worked_shifts":0,"availability":{"2025-07-11":["morning"]}},
    {"id":"w2","name":"Bob","max_shifts":1,"worked_shifts":0,"availability":{"2025-07-11":["morning","evening"]}},
    {"id":"w3","name":"Charlie","max_shifts":1,"worked_shifts":0,"availability":{"2025-07-11":["evening"]}}
  ]
}
```

The expected response
```
[
  {"date":"2025-07-11","slot":"morning","assigned":[{"id":"w1","name":"Alice"},{"id":"w2","name":"Bob"}]},
  {"date":"2025-07-11","slot":"evening","assigned":[{"id":"w3","name":"Charlie"}]}
]
```