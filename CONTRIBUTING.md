# Contributing to sockscope

Thanks for your interest. `sockscope` aims to stay small and focused, so contributions that preserve that focus are the easiest to merge.

## Scope rule

The v1 scope is deliberately narrow: **given a process, show its sockets and the first context needed for triage.** Before opening a PR that adds new features, please open an issue first so we can agree on whether it fits.

Things we are unlikely to accept in v1:

- Packet capture
- Web UIs or TUIs
- Kubernetes API discovery
- DNS reverse lookup enrichment
- Live watch (planned for v2 via eBPF)

## Development

```bash
git clone https://github.com/topcug/sockscope
cd sockscope
go build ./...
go test ./...
```

`sockscope` targets Linux only. Parsers under `internal/proc` should be testable against fixtures in `test/fixtures`, not against the live `/proc` of the developer's machine.

## Style

- Keep public surface small. If a function does not need to be exported, do not export it.
- Comments explain _why_, not _what_.
- No risk scores. Triage output uses "review" and "worth checking", never "suspicious" or "malicious".

## Commit messages

Conventional Commits are appreciated but not required. A clear one-line summary is enough.
