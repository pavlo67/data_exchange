package records

import (
	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/auth"
	"github.com/pavlo67/common/common/db"
	"github.com/pavlo67/common/common/selectors"
	"github.com/pavlo67/data/components/structures"
	"github.com/pavlo67/data/components/tags"
)

type ID common.IDStr

type Content struct {
	Title   string `json:",omitempty" bson:",omitempty"`
	Summary string `json:",omitempty" bson:",omitempty"`
	Type    string `json:",omitempty" bson:",omitempty"`
	Data    string `json:",omitempty" bson:",omitempty"`
}

type Item struct {
	ID       ID `       json:",omitempty" bson:"_id,omitempty"`
	Content  `          json:",inline"    bson:",inline"`
	Embedded []Content `json:",omitempty" bson:"_id,omitempty"`

	structures.ItemDescription `json:",inline" bson:",inline"`
}

type Operator interface {
	Save(Item, *auth.Identity) (ID, error)
	Read(ID, *auth.Identity) (*Item, error)
	Remove(ID, *auth.Identity) error

	List(*selectors.Term, *auth.Identity) ([]Item, error)
	Stat(*selectors.Term, *auth.Identity) (db.StatMap, error)
	Tags(*selectors.Term, *auth.Identity) (tags.StatMap, error)

	AddParent(ts []tags.Item, id ID) ([]tags.Item, error)
}

func ReadWithChildren(recordsOp Operator, id ID, identity *auth.Identity) (*Item, []Item, error) {
	r, err := recordsOp.Read(id, identity)
	if err != nil {
		return r, nil, err
	}

	selectorParent := selectors.Term{
		Key:    HasParent,
		Values: []string{string(id)},
	}

	children, err := recordsOp.List(&selectorParent, identity)
	return r, children, err
}
