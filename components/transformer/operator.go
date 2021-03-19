package transformer

import (
	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/selectors"

	"github.com/pavlo67/data_exchange/components/structures"
)

type Operator interface {
	Name() string

	In(pack structures.Pack, params common.Map) error                         // import from external source
	Out(selector *selectors.Term, params common.Map) (structures.Pack, error) // export to external source
	Stat(selector *selectors.Term, params common.Map) (interface{}, error)    // internal storage statistics
	Copy(selector *selectors.Term, params common.Map) (interface{}, error)    // internal storage snapshot
}
