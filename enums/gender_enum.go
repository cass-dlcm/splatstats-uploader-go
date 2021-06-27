package enums

import (
	"encoding/json"
	"errors"
	"fmt"
)

type GenderEnum string

const (
	boy        GenderEnum = "boy"
	girl       GenderEnum = "girl"
	genderNone GenderEnum = ""
)

func (ge *GenderEnum) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type GE GenderEnum
	r := (*GE)(ge)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *ge {
	case boy, girl, genderNone:
		return nil
	}
	return errors.New("Invalid GenderEnum. Got: " + fmt.Sprint(*ge))
}
