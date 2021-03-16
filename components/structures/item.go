package structures

import (
	"time"

	"github.com/pavlo67/common/common"

	"github.com/pavlo67/data_exchange/components/ns"
	"github.com/pavlo67/data_exchange/components/vcs"
)

type ItemDescription struct {
	Label     string      `json:",omitempty" bson:",omitempty"`
	Info      common.Map  `json:",omitempty" bson:",omitempty"`
	Tags      []string    `json:",omitempty" bson:",omitempty"`
	URN       ns.URN      `json:",omitempty" bson:",omitempty"`
	PackURN   ns.URN      `json:",omitempty" bson:",omitempty"`
	OwnerNSS  ns.NSS      `json:",omitempty" bson:",omitempty"`
	ViewerNSS ns.NSS      `json:",omitempty" bson:",omitempty"`
	History   vcs.History `json:",omitempty" bson:",omitempty"`
	CreatedAt time.Time   `json:",omitempty" bson:",omitempty"`
	UpdatedAt *time.Time  `json:",omitempty" bson:",omitempty"`
}
