package render

import (
	"fmt"
	"io"
	"strings"

	"github.com/topcug/sockscope/internal/model"
)

// Table writes a human-readable, terminal-friendly report. We
// deliberately do not pull in a table library: the layout is stable
// enough that plain Fprintf keeps the binary small and the output
// copy-pasteable into incident notes.
func Table(w io.Writer, r model.Report) error {
	fmt.Fprintln(w, "Process")
	fmt.Fprintf(w, "  PID:              %d\n", r.Process.PID)
	fmt.Fprintf(w, "  Name:             %s\n", firstWord(r.Process.Command))
	fmt.Fprintf(w, "  Command:          %s\n", r.Process.Command)
	fmt.Fprintf(w, "  PPID:             %d\n", r.Process.PPID)
	fmt.Fprintf(w, "  UID:              %d\n", r.Process.UID)
	if r.Process.ContainerID != "" {
		fmt.Fprintf(w, "  Container:        %s\n", shortID(r.Process.ContainerID))
	}
	if r.Process.CgroupPath != "" {
		fmt.Fprintf(w, "  Cgroup:           %s\n", firstLine(r.Process.CgroupPath))
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Sockets")
	if len(r.Sockets) == 0 {
		fmt.Fprintln(w, "  (none)")
	}
	for _, s := range r.Sockets {
		switch s.Kind {
		case model.KindTCP:
			fmt.Fprintf(w, "  TCP   %-22s -> %-22s %s\n", s.LocalAddress, s.RemoteAddress, s.State)
		case model.KindUDP:
			fmt.Fprintf(w, "  UDP   %-22s -> %s\n", s.LocalAddress, s.RemoteAddress)
		case model.KindUnix:
			path := s.Path
			if path == "" {
				path = "(abstract or unnamed)"
			}
			fmt.Fprintf(w, "  UNIX  %s\n", path)
		}
	}

	if len(r.Hints) > 0 {
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Triage notes")
		for _, h := range r.Hints {
			fmt.Fprintf(w, "  - %s\n", h)
		}
	}
	return nil
}

func firstWord(s string) string {
	if i := strings.IndexByte(s, ' '); i >= 0 {
		return s[:i]
	}
	return s
}

func firstLine(s string) string {
	if i := strings.IndexByte(s, '\n'); i >= 0 {
		return s[:i]
	}
	return s
}

func shortID(id string) string {
	if len(id) > 12 {
		return id[:12]
	}
	return id
}
