# Sha Sum API

This API assists HTTP downloaders that use range requests
achieve reliable downloads by returning SHA256 sums of specific
byte ranges. These can be compared to downloaded ranges to allow retries
if there are errors.

## Building

When building for prod 
```
export GIN_MODE=release
go build -ldflags "-s -w"
```