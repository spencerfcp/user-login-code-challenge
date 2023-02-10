## Getting Started

First, run the development server:

- `pnpm install`
- `pnpm dev`

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

## Running Tests

- Run: `pnpm test`

If unfamiliar, libraries can be found here:

https://jestjs.io/docs/getting-started
https://playwright.dev/docs/intro

## Protobuff

`./pb/api.ts` are type defintions automatically generated from the backend and should not be updated manually.
If there is a mismatch, you will need to rereun `backend/pb/generate.sh`
