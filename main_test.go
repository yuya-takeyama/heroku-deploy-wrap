package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExit0(t *testing.T) {
	stdin := new(bytes.Buffer)
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	exitStatus, err := herokuDeployWrap("sh", []string{"-c", "echo \"remote: Verifying deploy... done.\""}, stdin, stdout, stderr)

	assert.Equal(t, 0, exitStatus)
	assert.Nil(t, err)
}

func TestExit0WithMorePatterns(t *testing.T) {
	for i := 1; i <= 10; i++ {
		stdin := new(bytes.Buffer)
		stdout := new(bytes.Buffer)
		stderr := new(bytes.Buffer)
		dots := strings.Repeat(".", i)
		exitStatus, err := herokuDeployWrap("sh", []string{"-c", "echo \"remote: Verifying deploy" + dots + " done.\""}, stdin, stdout, stderr)

		assert.Equal(t, 0, exitStatus)
		assert.Nil(t, err)
	}
}

func TestDeployCommandFailed(t *testing.T) {
	stdin := new(bytes.Buffer)
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	exitStatus, err := herokuDeployWrap("sh", []string{"-c", "echo \"remote: Verifying deploy... failed!\""}, stdin, stdout, stderr)

	assert.Equal(t, 1, exitStatus)
	assert.Nil(t, err)
}

func TestDeployComandExitWithNonZero(t *testing.T) {
	stdin := new(bytes.Buffer)
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	exitStatus, err := herokuDeployWrap("sh", []string{"-c", "echo \"remote: Verifying deploy... done.\"; exit 128"}, stdin, stdout, stderr)

	assert.Equal(t, 128, exitStatus)
	assert.Nil(t, err)
}

func TestDeployCommandFailedAndExitWithNonZero(t *testing.T) {
	stdin := new(bytes.Buffer)
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	exitStatus, err := herokuDeployWrap("sh", []string{"-c", "echo \"remote: Verifying deploy... failed!\"; exit 128"}, stdin, stdout, stderr)

	assert.Equal(t, 128, exitStatus)
	assert.Nil(t, err)
}
