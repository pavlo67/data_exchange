package exchange

import (
	"github.com/pavlo67/common/common/crud"
)

type Operator interface {
	Import(data []byte, path string) (filenames []string, err error)
	Export(path string) (data []byte, filenames []string, err error)

	Save(*crud.Options) error
	Read(*crud.Options) error
}
