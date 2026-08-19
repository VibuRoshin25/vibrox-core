# vibrox-core

`vibrox-core` is the public API gateway for the Vibrox Systems Lab. The web
frontend reaches it through the `/api` reverse proxy, while experiment services
use internal protocols such as gRPC behind it.

The former user CRUD and custom authentication code was intentionally removed:
the portfolio does not need user accounts, and retaining placeholder identity
infrastructure would add complexity without demonstrating a real requirement.

## Current endpoint

```text
GET /health
```

```text
POST /arena/moves
```

The Arena endpoint translates browser JSON into an internal gRPC call.
Authentication should use a maintained identity provider if a future experiment
genuinely requires accounts.
