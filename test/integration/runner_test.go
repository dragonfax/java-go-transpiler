package integration

import (
	"fmt"
	"os/exec"
	"testing"

	"gotest.tools/assert"
)

func execute(cmdArgs ...string) error {
	err, output := capture(cmdArgs...)
	fmt.Println(output)
	return err
}

func capture(cmdArgs ...string) (error, string) {

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return err, ""
	}
	if cmd.ProcessState.ExitCode() != 0 {
		return fmt.Errorf("none zero exit status"), ""
	}

	return nil, string(stdoutStderr)
}

func TestHelloWorld(t *testing.T) {

	dir := "test/integration/hello_world/src/"
	err := execute("make", "target=", "source="+dir, "package=")
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = execute("javac", "test/integration/hello_world/src/Main.java")
	if err != nil {
		t.Fatalf(err.Error())
	}

	err, output := capture("java", "test/integration/hello_world/src/Main.class")
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, "hello world", output)
}
