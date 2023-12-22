package GoFFprobe

import (
	"encoding/json"
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

// TestExecute is a unit test function that tests the Execute function.
// It verifies that the Execute function returns the expected result and handles errors correctly.
func TestExecute(t *testing.T) {
	execCommand = mockExecCommand
	defer func() { execCommand = exec.Command }()

	// Mock the output of the helper process
	expectedOutput := `{"format": {"filename": "fakepath", "nb_streams": 2}, "streams": [{"index": 0}, {"index": 1}]}`

	// Execute the function under test
	result, err := Execute("fakepath", Options{ShowFormat: true, ShowStreams: true})
	if err != nil {
		t.Errorf("Execute returned an error: %v", err)
	}

	// Verify the result
	var parsedResult map[string]interface{}
	err = json.Unmarshal([]byte(expectedOutput), &parsedResult)
	if err != nil {
		t.Errorf("Failed to parse expected output: %v", err)
	}

	if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", parsedResult) {
		t.Errorf("Execute returned incorrect result. Expected: %v, got: %v", parsedResult, result)
	}
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	fmt.Fprint(os.Stdout, `{"format": {"filename": "fakepath", "nb_streams": 2}, "streams": [{"index": 0}, {"index": 1}]}`)
	os.Exit(0)
}
