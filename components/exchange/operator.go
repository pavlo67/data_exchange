package exchange

import (
	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/selectors"
)

type Version string

type Operator interface {
	//Name() string
	//Version() Version

	// Reset() is required for .Read(), while Save(), Import() and Export() should clear internal storage automatically
	Reset() error
	Read(selector *selectors.Term) error // from internal database
	Stat(params common.Map) error        // from internal database
	Save(selector *selectors.Term) error // into internal database

	Import(selector *selectors.Term, structure, data interface{}) error       // from external source
	Export(selector *selectors.Term) (structure, data interface{}, err error) // to external source
}
