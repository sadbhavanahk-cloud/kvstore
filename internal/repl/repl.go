package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"kvstore/internal/store"
)

func Run(in io.Reader, out io.Writer, s *store.Store) {
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		line := scanner.Text()

		resp, quit := Execute(s, line)
		if resp != "" {
			fmt.Fprintln(out, resp)
		}
		if quit {
			return
		}
	}
}

func Execute(s *store.Store, line string) (string, bool) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return "", false
	}

	cmd := strings.ToUpper(parts[0])

	switch cmd {

	case "WRITE":
		if len(parts) < 3 {
			return "usage: write <key> <value>", false
		}
		s.Write(parts[1], strings.Join(parts[2:], " "))
		return "", false

	case "READ":
		if len(parts) < 2 {
			return "usage: read <key>", false
		}
		v, err := s.Read(parts[1])
		if err != nil {
			return err.Error(), false
		}
		return v, false

	case "DELETE":
		if len(parts) < 2 {
			return "usage: delete <key>", false
		}
		err := s.Delete(parts[1])
		if err != nil {
			return err.Error(), false
		}
		return "", false

	case "START":
		s.Start()
		return "", false

	case "ABORT":
		err := s.Abort()
		if err != nil {
			return err.Error(), false
		}
		return "", false

	case "COMMIT":
		err := s.Commit()
		if err != nil {
			return err.Error(), false
		}
		return "", false

	case "QUIT":
		return "bye", true
	}

	return "unknown command", false
}

