package transform

import (
	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/selectors"
)

type Operator interface {
	Reset() error                 // for internal storage
	Stat(params common.Map) error // for internal storage

	In(selector *selectors.Term, data interface{}) error        // from external source
	Out(selector *selectors.Term) (data interface{}, err error) // to external source

}

// type Version string
// Name() string
// Version() Version
