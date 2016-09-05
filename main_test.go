package main

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuya-takeyama/posixexec"
)

func TestExit0(t *testing.T) {
	cmd := exec.Command("./heroku-deploy-wrap", "--", "sh", "-c", "echo \"remote: Verifying deploy... done.\"")
	exitStatus, err := posixexec.Run(cmd)

	assert.Equal(t, 0, exitStatus)
	assert.Nil(t, err)
}

func TestDeployCommandFailed(t *testing.T) {
	cmd := exec.Command("./heroku-deploy-wrap", "--", "sh", "-c", "echo \"remote: Verifying deploy... failed!\"")
	exitStatus, err := posixexec.Run(cmd)

	assert.Equal(t, 1, exitStatus)
	assert.Nil(t, err)
}
func TestDeployComandExitWithNonZero(t *testing.T) {
	cmd := exec.Command("./heroku-deploy-wrap", "--", "sh", "-c", "echo \"remote: Verifying deploy... done.\"; exit 128")
	exitStatus, err := posixexec.Run(cmd)

	assert.Equal(t, 128, exitStatus)
	assert.Nil(t, err)
}
