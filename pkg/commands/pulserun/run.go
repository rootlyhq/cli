package pulserun

import (
	"os"
	"os/exec"
	"strings"

	"github.com/rootlyhq/cli/pkg/log"
)

// Run a given program
func RunProgram(program string, args []string) (int, log.CtxErr) {
	fmtCommand := program + " " + strings.Join(args, " ")
	log.Info("Running", fmtCommand)

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
			log.Warning(fmtCommand, "failed with exit code of", exitCode)
			return exitErr.ExitCode(), log.CtxErr{}
		}
	}

	log.Success(false, "Ran", fmtCommand, "with an exit code of 0")
	return 0, log.CtxErr{}
}
