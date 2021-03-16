package transformer

import (
	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/selectors"
)

type Operator interface {
	Name() string

	In(params common.Map, data interface{}) error                                  // import from external source
	Out(selector *selectors.Term, params common.Map) (data interface{}, err error) // export to external source

	Stat(selector *selectors.Term, params common.Map) (interface{}, error) // internal storage statistics
	Copy(selector *selectors.Term, params common.Map) (interface{}, error) // internal storage snapshot
}
