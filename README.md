# GotByond

A WIP web requests library for Byond, written in Go.

## Current and Planned Features

- [x] GET Requests
- [ ] POST Requests

### Why GET requests? Doesn't BYOND already support that?

Yes, but BYOND considers some successful response codes to be errors and logs them as a bug in the console. It's very annoying if you're doing a lot of requests.