package util

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/aburg/tjann/models"
)

func WriteTime(time models.Time) error {
	data, err := json.Marshal(time.Annotation)
	if err != nil {
		panic(err)
	}
	return setAnnotation(fmt.Sprintf("@%d", time.ID), string(data))
}

func setAnnotation(from string, annotation string) error {
	cmd := exec.Command("timew", "annotate", from, annotation)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error writing to timewarrior id %s: %s", from, err)
	}
	fmt.Println(string(out))
	return nil
}
