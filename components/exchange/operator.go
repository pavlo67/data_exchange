package exchange

import (
	"github.com/pavlo67/common/common/crud"
	"github.com/pavlo67/common/common/selectors"
)

type Version string

type Operator interface {
	Name() string
	Version() Version

	Read(selector *selectors.Term) error // from internal database
	Save(selector *selectors.Term) error // into internal database

	Import(selector *selectors.Term, structure, data interface{}) error       // from external source
	Export(selector *selectors.Term) (structure, data interface{}, err error) // to external source

	// Clear is required for .Read()
	// while Save(), .Import() and .Export() should clear operator's internal storage automatically
	crud.Cleaner
}
