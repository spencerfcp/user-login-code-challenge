# Running Tests

- `go test ./...`

## Protobuf Generation

PB generates TypeScript and GO that should not be updated manually. To make changes, update `./pb/api.proto`

`./pb` and run `./generate.sh` to generate new type definitions for both the frontend and backend.
