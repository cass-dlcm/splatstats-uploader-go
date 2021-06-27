package types

import "github.com/cass-dlcm/splatstatsuploader/enums"

type SplatnetTriple struct {
	Id    string `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
}

type SplatnetDouble struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type SplatnetPlayerType struct {
	Gender  enums.GenderEnum  `json:"style,omitempty"`
	Species enums.SpeciesEnum `json:"species,omitempty"`
}

type SplatnetQuad struct {
	Id     string `json:"id"`
	ImageA string `json:"image_a"`
	ImageB string `json:"image_b"`
	Name   string `json:"name"`
}
