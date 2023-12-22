package GoFFprobe

import (
	"encoding/json"
	"os/exec"
)

// Options defines the options for running ffprobe.
type Options struct {
	ShowFormat  bool
	ShowStreams bool
}

var execCommand = exec.Command

// Execute runs ffprobe with the provided options and returns the parsed data.
func Execute(filePath string, opts Options) (map[string]interface{}, error) {
	args := buildArgs(opts, filePath)
	cmd := execCommand("ffprobe", args...)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(output, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// buildArgs builds the command line arguments for ffprobe.
func buildArgs(opts Options, filePath string) []string {
	args := []string{"-v", "quiet", "-print_format", "json"}

	if opts.ShowFormat {
		args = append(args, "-show_format")
	}
	if opts.ShowStreams {
		args = append(args, "-show_streams")
	}
	args = append(args, filePath)
	return args
}
