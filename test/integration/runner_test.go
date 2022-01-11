package integration

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func execute(cmdArgs ...string) error {
	output, err := capture(cmdArgs...)
	fmt.Println(output)
	return err
}

func capture(cmdArgs ...string) (string, error) {

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Dir = "../../"
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return string(stdoutStderr), err
	}
	if cmd.ProcessState.ExitCode() != 0 {
		return string(stdoutStderr), fmt.Errorf("none zero exit status")
	}

	return string(stdoutStderr), nil
}

func TestHelloWorld(t *testing.T) {

	dir := "test/integration/hello_world/src/"

	// test that the java code works.
	{
		err := execute("javac", dir+"/Main.java")
		if err != nil {
			t.Fatalf(err.Error())
		}

		output, err := capture("java", "test/integration/hello_world/src/Main.class")
		fmt.Println(output)
		if err != nil {
			t.Fatalf(err.Error())
		}

		assert.Equal(t, "hello world", output)
	}

	// transpile and test
	{
		err := execute("make", "run", "target=", "source="+dir, "package=")
		if err != nil {
			t.Fatalf(err.Error())
		}
	}

}
