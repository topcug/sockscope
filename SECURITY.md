# Security Policy

## Reporting a vulnerability

If you believe you have found a security issue in `sockscope`, please do **not** open a public GitHub issue. Instead, email the maintainers with:

- A description of the issue
- Steps to reproduce
- The version of `sockscope` and the Linux distribution you are running

We will acknowledge the report within 72 hours and aim to provide a fix or mitigation within 14 days for confirmed issues.

## Scope

`sockscope` is a read-only tool. It reads `/proc` and does not open network sockets, modify state, or require elevated privileges beyond what `/proc` access already implies. Any behaviour that contradicts this is in scope for a security report.

## Out of scope

- Issues that require an attacker to already have the same privileges as the target process
- Information disclosed by `/proc` that is already visible to the invoking user through standard tools (`ps`, `ss`, `lsof`)
