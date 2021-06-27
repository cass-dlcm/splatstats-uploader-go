package enums

import (
	"encoding/json"
	"errors"
)

type SpeciesEnum string

const (
	inklings  SpeciesEnum = "inklings"
	octolings SpeciesEnum = "octolings"
)

func (se *SpeciesEnum) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type SE SpeciesEnum
	r := (*SE)(se)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *se {
	case inklings, octolings:
		return nil
	}
	return errors.New("invalid type")
}
