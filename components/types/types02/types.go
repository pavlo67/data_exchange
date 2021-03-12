package types02

import (
	"time"

	"github.com/pavlo67/common/common/auth"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/rbac"

	"github.com/pavlo67/data_exchange/components/ns"
	"github.com/pavlo67/data_exchange/components/vcs"
)

// persons -------------------------------------------------------------

type Person struct {
	URN       ns.URN      // TODO: ba careful, URN can't be empty
	Nickname  string      `json:",omitempty" bson:",omitempty"`
	Roles     rbac.Roles  `json:",omitempty" bson:",omitempty"`
	Creds     auth.Creds  `json:",omitempty" bson:",omitempty"`
	Data      common.Map  `json:",omitempty" bson:",omitempty"`
	History   vcs.History `json:",omitempty" bson:",omitempty"`
	CreatedAt time.Time   `json:",omitempty" bson:",omitempty"`
	UpdatedAt *time.Time  `json:",omitempty" bson:",omitempty"`
}

// records -------------------------------------------------------------

type Content struct {
	Title    string    `json:",omitempty" bson:",omitempty"`
	Summary  string    `json:",omitempty" bson:",omitempty"`
	DataType string    `json:",omitempty" bson:",omitempty"`
	Data     string    `json:",omitempty" bson:",omitempty"`
	Embedded []Content `json:",omitempty" bson:",omitempty"`
	Tags     []string  `json:",omitempty" bson:",omitempty"`
}

type Record struct {
	URN       ns.URN      // TODO: ba careful, URN can't be empty
	OwnerURN  ns.URN      `json:",omitempty" bson:",omitempty"`
	ViewerURN ns.URN      `json:",omitempty" bson:",omitempty"`
	Content   Content     `json:",inline"    bson:",inline"`
	History   vcs.History `json:",omitempty" bson:",omitempty"`
	CreatedAt time.Time   `json:",omitempty" bson:",omitempty"`
	UpdatedAt *time.Time  `json:",omitempty" bson:",omitempty"`
}
