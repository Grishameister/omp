package csvParser

import (
	"encoding/csv"
	"errors"
	"github.com/Grishameister/omp/domain"
	"io"
	"strconv"
)

type CsvParser struct{}

func (CsvParser) Parse(reader io.Reader, r *domain.Result) error {
	parser := csv.NewReader(reader)
	parser.Comma = ','

	st := domain.Stat{}
	first := true
	for {
		record, err := parser.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.New("can't parse row")
		}

		if len(record) != 3 {
			return errors.New("size invalid")
		}

		if first {
			first = false
			continue
		}

		st.Product = record[0]
		price, err := strconv.Atoi(record[1])
		if err != nil {
			return errors.New("price isn't int")
		}
		st.Price = price
		rating, err := strconv.Atoi(record[2])
		if err != nil {
			return errors.New("rating isn't int")
		}
		st.Rating = rating
		r.Compare(&st)
	}
	return nil
}
