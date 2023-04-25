package pipefilter

import (
	"errors"
	"strings"
)

var ErrSplitFilterWrongFormatError = errors.New("input data should be string")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string)
	if !ok {
		return nil, ErrSplitFilterWrongFormatError
	}

	parts := strings.Split(str, sf.delimiter)

	return parts, nil
}