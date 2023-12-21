package GoFFprobe

import (
	"encoding/json"
	"os/exec"
)

// Options définit les options pour exécuter ffprobe.
type Options struct {
	ShowFormat  bool
	ShowStreams bool
}

var execCommand = exec.Command

// Execute exécute ffprobe avec les options fournies et retourne les données parsées.
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

// buildArgs construit les arguments de la ligne de commande pour ffprobe.
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
