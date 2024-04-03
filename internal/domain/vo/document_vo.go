package vo

import (
	"errors"
	"regexp"
)

type DocumentVo struct {
	value string
}

func NewDocument(document string) (*DocumentVo, error) {
	if isInvalidDocument(document) {
		return nil, errors.New("invalid document")
	}

	return &DocumentVo{
		value: document,
	}, nil
}

func (d *DocumentVo) Value() string {
	return d.value
}

func isInvalidDocument(document string) bool {
	return !regexp.MustCompile(`[0-9]{3}\.[0-9]{3}\.[0-9]{3}-[0-9]{2}`).MatchString(document)
}
