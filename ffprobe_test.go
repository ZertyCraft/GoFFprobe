package GoFFprobe

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func mockExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func TestExecute(t *testing.T) {
	execCommand = mockExecCommand
	defer func() { execCommand = exec.Command }()

	_, err := Execute("fakepath", Options{ShowFormat: true, ShowStreams: true})
	if err != nil {
		t.Errorf("Execute returned an error: %v", err)
	}
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	fmt.Fprint(os.Stdout, `{"format": {"filename": "fakepath", "nb_streams": 2}, "streams": [{"index": 0}, {"index": 1}]}`)
	os.Exit(0)
}
