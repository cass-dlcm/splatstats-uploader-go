package types

import (
	"github.com/cass-dlcm/splatstatsuploader/enums"
	"time"
)

type StatInkKeyName struct {
	Key  string      `json:"key"`
	Name StatInkName `json:"name"`
}

type StatInkName struct {
	DeDE string `json:"de_DE,omitempty"`
	EnGB string `json:"en_GB,omitempty"`
	EnUS string `json:"en_US,omitempty"`
	EsES string `json:"es_ES,omitempty"`
	EsMX string `json:"es_MX,omitempty"`
	FrCA string `json:"fr_CA,omitempty"`
	FrFR string `json:"fr_FR,omitempty"`
	ItIT string `json:"it_IT,omitempty"`
	JaJP string `json:"ja_JP,omitempty"`
	NlNL string `json:"nl_NL,omitempty"`
	RuRU string `json:"ru_RU,omitempty"`
	ZhCN string `json:"zh_CN,omitempty"`
	ZhTW string `json:"zh_TW,omitempty"`
}

type StatInkGender struct {
	Key     enums.GenderEnum `json:"key,omitempty"`
	Name    StatInkName      `json:"name,omitempty"`
	Iso5218 int              `json:"iso5218,omitempty"`
}

type StatInkTime struct {
	Time    int       `json:"time"`
	Iso8601 time.Time `json:"iso8601"`
}

type StatInkProfile struct {
	Nnid        *interface{} `json:"nnid,omitempty"`
	FriendCode  *string      `json:"friend_code,omitempty"`
	Twitter     *string      `json:"twitter,omitempty"`
	Ikanakama   *interface{} `json:"ikanakama,omitempty"`
	Ikanakama2  *interface{} `json:"ikanakama2,omitempty"`
	Environment *interface{} `json:"environment,omitempty"`
}
