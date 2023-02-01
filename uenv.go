package uenv

import (
	"bufio"
	"os"
	"strings"
)

func SetDotEnv() {
	f, err := os.Open(".env")
	if err != nil {
		return
	}

	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		if err := os.Setenv(parts[0], parts[1]); err != nil {
			continue
		}
	}
}
