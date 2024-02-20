package commands

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// RunCommand function is general function to run a command with parameters.
func RunCommand(command string, params []string) (errMessage string, cmdOutput []string) {
	cmd := exec.Command(command, params...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	completeCommand := []string{command}
	completeCommand = append(completeCommand, params...)
	logger.Printf("executing command -> \"%s\"\n", strings.Join(completeCommand, " "))
	if err != nil {
		errMsg := fmt.Sprint(err) + ": " + stderr.String()
		fmt.Println(params)
		if !strings.Contains(errMsg, "No such file or directory") {
			// Long time sleep, so we can see the error message better
			// This shouldn't be happening
			time.Sleep(10 * time.Second)
		}
		return errMsg, []string{}
	}
	return "", strings.SplitAfter(out.String(), "\n")
}
