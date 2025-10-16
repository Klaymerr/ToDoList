package valueobject

import "errors"

type Text struct {
	value string
}

func NewText(value string) (*Text, error) {
	if value == "" {
		return nil, errors.New("value is empty")
	}
	return &Text{value: value}, nil
}

func (text *Text) Value() string {
	return text.value
}
