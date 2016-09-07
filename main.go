package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
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

	info(fmt.Sprintf("Version: %s , Revision: %s", Version, GitCommit))
	info(fmt.Sprintf("Running command: %s %s", cmdName, strings.Join(cmdArgs, " ")))

	exitStatus, execErr := herokuDeployWrap(cmdName, cmdArgs, os.Stdin, os.Stdout, os.Stderr)
	if execErr != nil {
		panic(execErr)
	}

	if exitStatus == 0 {
		info("Successfully deployed!")
	} else {
		info("Failed to deploy...")
	}

	os.Exit(exitStatus)
}

var deploySuccessPattern = regexp.MustCompile(`^(?:remote: Verifying deploy\.{1,10} done\.|Everything up-to-date)\n?$`)

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

	info(fmt.Sprintf("Command exitted with %d", exitStatus))

	if exitStatus == 0 && !deploySuccessPattern.MatchReader(resultBuffer) {
		exitStatus = 1
	}

	return exitStatus, nil
}

func info(s string) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", AppName, s)
}
