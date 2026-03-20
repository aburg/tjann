package util

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/aburg/tjann/models"
)

func ReadTimes(from string) (models.Times, error) {
	cmd := exec.Command("timew", "export", from)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return models.Times{}, fmt.Errorf("error reading from %s: %s", from, err)
	}

	var times models.Times
	err = json.Unmarshal(out, &times)
	if err != nil {
		return models.Times{}, err
	}

	return times, nil
}
