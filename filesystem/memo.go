package filesystem

import (
	"encoding/json"
)

type Memo struct {
	Id           int    `json:"id,omitempty"`
	LastEditedBy string `json:"lastEditedBy"`
	Content      string `json:"content"`
}

func (memo *Memo) String() string {
	jsonString, _ := json.Marshal(memo)

	return string(jsonString)
}

func ParseMemo(raw []byte) (memo Memo, err error) {
	err = json.Unmarshal(raw, &memo)
	return memo, err
}
