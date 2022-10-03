# How to run

`make start`

This will spin up a postgres server and a HTTP server serving on localhost:8080.

## Routes

These are the routes available:

```
POST /api/v1/flights
    Data: ["SFO", "EWR"]
    example response: {}

GET /api/v1/flights
    example response: [["SFO", "EWR"], [...]]

GET /api/v1/flights/<from destination>
    example response: [["SFO", "EWR"], [...]]
```
