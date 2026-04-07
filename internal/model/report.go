package model

import "time"

// Report is the single object produced by every `sockscope inspect`
// run. All three output formats (table, json, markdown) are built
// from this struct so that outputs stay consistent.
type Report struct {
	Process     ProcessSummary  `json:"process"`
	Sockets     []SocketSummary `json:"sockets"`
	Hints       []string        `json:"hints"`
	GeneratedAt time.Time       `json:"generated_at"`
}
