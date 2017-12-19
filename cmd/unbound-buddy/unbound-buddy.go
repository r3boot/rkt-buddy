package main

import (
	"flag"

	"bytes"
	"regexp"

	"os/exec"

	"strings"

	"github.com/r3boot/rkt-buddy/lib/buddy"
	"github.com/r3boot/rkt-buddy/lib/logger"
)

const (
	D_DEBUG = false
)

var (
	useDebug = flag.Bool("d", D_DEBUG, "Enable debugging mode")

	reUnboundIsRunning = regexp.MustCompile("^unbound \\(pid [0-9]{1,5}\\) is running...")

	Logger *logger.Logger
	Buddy  *buddy.Buddy
)

func UnboundHealthCheck() bool {
	var stdoutBuf, stderrBuf bytes.Buffer

	cmd := exec.Command("unbound-control", "status")
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err := cmd.Start()
	if err != nil {
		Logger.Warningf("UnboundHealthCheck cmd.Start: %v", err)
		return false
	}

	err = cmd.Wait()
	if err != nil {
		Logger.Warningf("UnboundHealthCheck cmd.Wait: %v", err)
		if stderrBuf.Len() > 0 {
			Logger.Warningf("UnboundHealthCheck stderr: %s", stderrBuf.String())
		}
		return false
	}

	stdout := ""
	if stdoutBuf.Len() > 0 {
		stdout = stdoutBuf.String()
	}

	for _, line := range strings.Split(stdout, "\n") {
		response := reUnboundIsRunning.FindAllStringSubmatch(line, -1)
		if len(response) > 0 {
			return true
		}
	}

	return false
}

func init() {
	var err error

	flag.Parse()

	Logger = logger.NewLogger(false, *useDebug)

	Buddy, err = buddy.NewBuddy(Logger, &buddy.BuddyConfig{
		HealthCheck: UnboundHealthCheck,
	})

	if err != nil {
		Logger.Fatalf("init: %v", err)
	}
}

func main() {
	Buddy.Run()
}
