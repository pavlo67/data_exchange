package transformer

import (
	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/selectors"
)

type Operator interface {
	Name() string

	Reset() error                                                          // internal storage reset
	Stat(selector *selectors.Term, params common.Map) (interface{}, error) // internal storage statistics
	Copy(selector *selectors.Term, params common.Map) (interface{}, error) // internal storage snapshot

	In(selector *selectors.Term, params common.Map, data interface{}) error        // import from external source
	Out(selector *selectors.Term, params common.Map) (data interface{}, err error) // export to external source
}
