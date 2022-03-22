package uenv

import (
	"bufio"
	"os"
	"strings"

	"github.com/dunv/ulog"
)

func init() {
	ulog.AddReplaceFunction("github.com/dunv/uenv.SetDotEnv", "uenv")
}

func SetDotEnv() {
	f, err := os.Open(".env")
	if err != nil {
		ulog.Infof("not parsing .env (%s)", err)
		return
	}

	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()

		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			ulog.Infof("skipping %s (no two parts)", parts[0])
			continue
		}
		if err := os.Setenv(parts[0], parts[1]); err != nil {
			ulog.Infof("could not set %s (%s)", parts[0], err)
			continue
		}

		ulog.Infof("setenv %s=%s", parts[0], parts[1])
	}
}
