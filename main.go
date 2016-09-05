package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	flags "github.com/jessevdk/go-flags"
	"github.com/yuya-takeyama/posixexec"
)

// AppName is displayed in help command
const AppName = "heroku-deploy-wrap"

type options struct {
}

var opts options

func main() {
	parser := flags.NewParser(&opts, flags.Default^flags.PrintErrors)
	parser.Name = AppName
	parser.Usage = "-- git push heroku master"

	args, err := parser.Parse()

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	cmdName := args[0]
	cmdArgs := args[1:]

	exitStatus, execErr := herokuDeployWrap(cmdName, cmdArgs, os.Stdin, os.Stdout, os.Stderr)
	if execErr != nil {
		panic(execErr)
	}

	os.Exit(exitStatus)
}

func herokuDeployWrap(cmdName string, cmdArgs []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) (int, error) {
	cmd := exec.Command(cmdName, cmdArgs...)

	resultBuffer := new(bytes.Buffer)

	cmd.Stdin = stdin
	cmd.Stdout = io.MultiWriter(stdout, resultBuffer)
	cmd.Stderr = io.MultiWriter(stderr, resultBuffer)

	exitStatus, err := posixexec.Run(cmd)

	if err != nil {
		return -1, err
	}

	if exitStatus == 0 && !strings.Contains(resultBuffer.String(), "remote: Verifying deploy... done.") {
		exitStatus = 1
	}

	return exitStatus, nil
}
