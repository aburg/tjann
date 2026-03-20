package models

import "encoding/json"

type Annotation map[string]string

// here comes AI code

func (s *Annotation) UnmarshalJSON(data []byte) error {
	// 1. Unmarshal the data into a temporary string
	var innerString string
	if err := json.Unmarshal(data, &innerString); err != nil {
		return err
	}

	// 2. Unmarshal that string into the actual map
	// We use a temporary map to avoid infinite recursion
	var tempMap map[string]string
	if err := json.Unmarshal([]byte(innerString), &tempMap); err != nil {
		return err
	}

	*s = tempMap
	return nil
}
