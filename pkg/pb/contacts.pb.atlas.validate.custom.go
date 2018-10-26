package pb

import (
	"encoding/json"
	"fmt"
)

func (o *Contact) AtlasJSONValidate(r json.RawMessage, p string, allowUnknown bool) (json.RawMessage, error) {
	var rr map[string]json.RawMessage
	if err := json.Unmarshal(r, &rr); err != nil {
		return r, fmt.Errorf("Invalid value for %q: expected object", p)
	}

	if _, ok := rr["wah"]; ok {
		if _, ok := rr["blah"]; ok {
			return r, fmt.Errorf("Cannot specify 'blah' and 'wah' in the same message")
		}
	}

	return r, nil
}
