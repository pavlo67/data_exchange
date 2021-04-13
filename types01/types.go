package types01

import (
	"github.com/pavlo67/common/common/auth"
	"github.com/pavlo67/data_exchange/components/ns"
	"github.com/pavlo67/data_exchange/components/structures"

	"github.com/pavlo67/common/common/rbac"
)

// persons -------------------------------------------------------------

type Person struct {
	Nickname                   string `json:",omitempty" bson:",omitempty"`
	rbac.Roles                 `       json:",omitempty" bson:",omitempty"`
	auth.Creds                 `       json:",omitempty" bson:",omitempty"`
	structures.ItemDescription `       json:",inline"    bson:",inline"`
}

// records -------------------------------------------------------------

type Content struct {
	NSS      ns.NSS    `json:",omitempty" bson:",omitempty"`
	Summary  string    `json:",omitempty" bson:",omitempty"`
	Type     string    `json:",omitempty" bson:",omitempty"`
	Data     string    `json:",omitempty" bson:",omitempty"`
	Embedded []Content `json:",omitempty" bson:",omitempty"`
}

type Record struct {
	Content                    `json:",inline" bson:",inline"`
	structures.ItemDescription `json:",inline" bson:",inline"`
}
