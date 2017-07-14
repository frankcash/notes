# Operability in Go

## Maintenance

Equal objectives:

1. Fix it
2. Identify what is wrong

## Fail Well

- Fail immediately when unrecoverable errors occur. Otherwise data will corrupt
- Fail the smalled execution as possible
- In general, an unhadled error should panic

## Logging

- Provide context, otherwise it will go nowhere
- Some errors provide context, Named errors do not
- You can added context within the `github.com/pkg/errors` lib.
- Structured logging adds k-v pairs to your log `github.com/sirupsen/logrus`.  This adds context.
- Structured loggers output JSON, making them easy to be consumed.
- Log environment and flag values
- `expvar` shows current state
- `expvar` for state, `log` for crashing
- Monitoring: prometheus

