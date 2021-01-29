package jsonParser

import (
	"encoding/json"
	"errors"
	"github.com/Grishameister/omp/domain"
	"io"
)

type JsonParser struct {}

func (j JsonParser) Parse(reader io.Reader, r *domain.Result) error {
	decoder := json.NewDecoder(reader)
	stat := domain.Stat{}

	if _, err := decoder.Token(); err != nil {
		return errors.New("Token is invalid at start of file")
	}
	for decoder.More() {
		err := decoder.Decode(&stat)
		if err != nil {
			return errors.New("Can't decode into struct")
		}
		r.Compare(&stat)
	}

	if _, err := decoder.Token(); err != nil {
		return errors.New("Token is invalid at start of file")
	}
	return nil
}
