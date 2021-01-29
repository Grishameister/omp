package interfaceParser

import (
	"github.com/Grishameister/omp/domain"
	"io"
)

type IParser interface {
	Parse(reader io.Reader, r *domain.Result) error
}
