package types01

import (
	"github.com/pavlo67/common/common/auth"
	"github.com/pavlo67/data_exchange/components/structures"

	"github.com/pavlo67/common/common/rbac"
)

// persons -------------------------------------------------------------

type Person struct {
	Nickname                   string `json:",omitempty" bson:",omitempty"`
	rbac.Roles                 `       json:",omitempty" bson:",omitempty"`
	structures.ItemDescription `       json:",inline"    bson:",inline"`
	auth.Creds                 `       json:",omitempty" bson:",omitempty"`
}

// records -------------------------------------------------------------

type Content struct {
	Summary  string    `json:",omitempty" bson:",omitempty"`
	DataType string    `json:",omitempty" bson:",omitempty"`
	Data     string    `json:",omitempty" bson:",omitempty"`
	Embedded []Content `json:",omitempty" bson:",omitempty"`
}

type Record struct {
	structures.ItemDescription `json:",inline" bson:",inline"`
	Content                    `json:",inline" bson:",inline"`
}
