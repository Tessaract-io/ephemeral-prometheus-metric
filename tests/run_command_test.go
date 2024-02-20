package tests

import (
	"ephemeral-prometheus-metric/commands"
	"strings"
	"testing"
)

func TestRunCommandEcho(t *testing.T) {
	errMsg, result := commands.RunCommand("echo", []string{"$PATH"})
	if errMsg != "" {
		t.Fatalf("Error message is not empty, value is %s", errMsg)
	}
	if len(result) == 0 {
		t.Fatalf("Result is empty, value is %s", result)
	}
}

func TestRunCommandPS(t *testing.T) {
	errMsg, result := commands.RunCommand("ps", []string{"-a"})
	if errMsg != "" {
		t.Fatalf("Error message is not empty, value is %s", errMsg)
	}
	if len(result) == 0 {
		t.Fatalf("Result is empty, value is %s", result)
	}
	if !strings.Contains(result[0], "PID TTY") && !strings.Contains(result[0], "PID   USER") {
		t.Fatalf("Result is not correct, value is %s", result)
	}
}
