package pipefilter

import (
	"errors"
	"strconv"
)

var errToIntFilterWrongFormatError = errors.New("input data should be []string")


type toIntFilter struct {

}

func NewToIntFilter() *toIntFilter {
	return &toIntFilter{}
}

func (tif *toIntFilter)Process(data Request) (Response, error) {
	parts, ok := data.([]string)

	if !ok {
		return nil, errToIntFilterWrongFormatError
	}

	ret := []int{}
	for _, part := range parts {
		s, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, s)
	}
	return ret, nil
}