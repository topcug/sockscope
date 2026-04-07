<p align="center">
  <img src="sockscope.png" alt="sockscope" width="280" />
</p>

<h1 align="center">sockscope</h1>

<p align="center">
  <strong>One command to see what a Linux process is talking to — and what to check first.</strong>
</p>

<p align="center">
  <a href="https://github.com/topcug/sockscope/actions/workflows/ci.yml"><img src="https://github.com/topcug/sockscope/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
  <a href="https://github.com/topcug/sockscope/releases/latest"><img src="https://img.shields.io/github/v/release/topcug/sockscope?sort=semver" alt="Release"></a>
  <a href="https://goreportcard.com/report/github.com/topcug/sockscope"><img src="https://goreportcard.com/badge/github.com/topcug/sockscope" alt="Go Report Card"></a>
  <a href="LICENSE"><img src="https://img.shields.io/github/license/topcug/sockscope" alt="License: Apache-2.0"></a>
  <img src="https://img.shields.io/badge/platform-linux-blue" alt="Linux only">
</p>

---

`sockscope` is a small CLI that helps engineers understand which sockets a process owns, where those connections go, and what to check first during runtime triage.

<p align="center">
  <img src="docs/demo.gif" alt="sockscope inspect demo" width="720" />
</p>

## First 30 seconds

```bash
# Install
go install github.com/topcug/sockscope@latest

# Point it at any running process
sockscope inspect --pid 1234
```

```text
Process
  PID:              1234
  Name:             python
  Command:          python app.py
  PPID:             1200
  UID:              1000
  Container:        payments-api7

Sockets
  TCP   10.42.1.15:48122       -> 34.120.55.2:443         ESTABLISHED
  TCP   127.0.0.1:8080         -> 127.0.0.1:37014         ESTABLISHED
  UNIX  /tmp/agent.sock

Triage notes
  - External connection present: review whether 34.120.55.2:443 is an expected outbound destination
  - IPC via UNIX socket present
```

That's the whole product. One command, one screen, first context.

## Why this exists

During runtime triage, engineers hop between `ss`, `lsof`, `/proc`, container metadata, and process details just to answer one question: _what is this process talking to, and should I be worried?_

`sockscope` collects that context into a single view. It reads `/proc` directly (no `ss` or `lsof` shell-outs), resolves the process, its sockets, its cgroup, and its container ID, and prints a small triage report you can paste into an incident note.

## What it is not

- Not a packet capture tool
- Not a port scanner
- Not a SIEM
- Not a full eBPF observability platform

It answers one narrow question well: **given a process, show its sockets and the first context needed for triage.**

## Install

Requires Linux. `/proc` is the data source, so macOS and Windows are not supported.

**With Go:**

```bash
go install github.com/topcug/sockscope@latest
```

**From a release archive:**

Download the latest `sockscope_<version>_linux_<arch>.tar.gz` from the [releases page](https://github.com/topcug/sockscope/releases), then:

```bash
tar -xzf sockscope_*.tar.gz
sudo mv sockscope /usr/local/bin/
sockscope --version
```

**From source:**

```bash
git clone https://github.com/topcug/sockscope
cd sockscope
go build -o sockscope .
```

## Usage

v1 exposes three ways to select a target:

```bash
sockscope inspect --pid 1234
sockscope inspect --name nginx
sockscope inspect --container-id 7bc4d5e9f1a2
```

`--name` matches against `/proc/<pid>/comm` exactly (the kernel truncates comm at 15 characters, so long process names should be matched by PID).

`--container-id` matches any PID whose `/proc/<pid>/cgroup` path contains the given ID, which covers Docker, containerd, CRI-O and most Kubernetes runtimes.

## Output formats

```bash
sockscope inspect --pid 1234                 # table (default)
sockscope inspect --pid 1234 -o json         # for jq and SIEM pipelines
sockscope inspect --pid 1234 -o markdown     # for issues, incident notes, Slack
```

The JSON shape is stable within a major version: `{process, sockets, hints, generated_at}`.

## Triage notes

`sockscope` adds a small, deliberately low-confidence notes section at the end of every report. It uses words like _review_, _worth checking_, and _confirm_, never _suspicious_ or _malicious_. There are no risk scores.

The goal is to highlight things that usually matter in a first look:

- External outbound connections
- Loopback-only communication
- Listening sockets
- UNIX socket IPC
- Root-owned processes
- Interactive shells

## Roadmap

- **v1** — `/proc`-based inspect by PID / name / container ID, with table, JSON and Markdown output
- **v1.1** — Better container awareness, namespace detection, pod metadata
- **v2** — Optional live watch (`sockscope watch`) built on eBPF, tracing `connect`, `accept`, `bind` and `close`

`watch` and `graph` are intentionally _not_ in v1. The whole point of v1 is the one-command first-look experience.

## Project

- [Changelog](CHANGELOG.md)
- [Contributing](CONTRIBUTING.md)
- [Security policy](SECURITY.md)
- [Code of conduct](CODE_OF_CONDUCT.md)
- [Release process](docs/RELEASING.md)

## License

Apache-2.0. See [LICENSE](LICENSE).
