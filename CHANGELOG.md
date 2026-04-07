# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - TBD

### Added
- Initial release.
- `sockscope inspect --pid <pid>` — inspect a specific process.
- `sockscope inspect --name <comm>` — inspect processes by `/proc/<pid>/comm` name.
- `sockscope inspect --container-id <id>` — inspect processes inside a given container.
- TCP, UDP, and UNIX socket resolution via `/proc/net/{tcp,tcp6,udp,udp6,unix}`.
- Process context: PID, PPID, command, UID, cgroup, container ID extraction.
- Table, JSON, and Markdown output formats (`-o table|json|markdown`).
- Triage notes layer with human-readable hints (no risk scores).
- Apache-2.0 license, `SECURITY.md`, `CONTRIBUTING.md`, `CODE_OF_CONDUCT.md`.
- GitHub Actions CI (build, vet, test, golangci-lint) and release workflow via GoReleaser.

[Unreleased]: https://github.com/topcug/sockscope/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/topcug/sockscope/releases/tag/v0.1.0
