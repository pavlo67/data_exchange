package exchange_0_1_records

import (
	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/crud"

	"github.com/pavlo67/data_exchange/components/exchange"
	"github.com/pavlo67/data_exchange/components/exchange/exchange_0_1"
)

var _ exchange.Operator = &exchange01Records{}

type exchange01Records struct {
	exchange_0_1.RecordItems
}

func (exchange01Records) Save(*crud.Options) error {
	return common.ErrNotImplemented
}

func (exchange01Records) Read(*crud.Options) error {
	return common.ErrNotImplemented
}
