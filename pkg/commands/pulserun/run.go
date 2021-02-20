package pulserun

import (
	"os"
	"os/exec"
	"strings"

	"github.com/rootly-io/cli/pkg/log"
)

// Run a given program
func RunProgram(program string, args []string) (int, log.CtxErr) {
	// Making sure that the program exists
	progPath, err := exec.LookPath(program)
	if err != nil {
		return 0, log.CtxErr{
			Context: "Looks like " + program + " is not installed on this machine. Can't find it.",
			Error:   err,
		}
	}

	// Creating the command and running it
	cmd := exec.Command(progPath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if ok {
			exitCode := exitErr.ExitCode()
			log.Warning(program, strings.Join(args, " "), "failed with exit code of", exitCode)
			return exitErr.ExitCode(), log.CtxErr{}
		}
	}

	return 0, log.CtxErr{}
}
