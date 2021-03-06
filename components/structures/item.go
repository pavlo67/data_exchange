package structures

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"

	"github.com/pavlo67/data/components/ns"
	"github.com/pavlo67/data/components/vcs"
)

type ItemDescription struct {
	URN       ns.URN      `json:",omitempty" bson:",omitempty"`
	Tags      []string    `json:",omitempty" bson:",omitempty"`
	OwnerNSS  ns.NSS      `json:",omitempty" bson:",omitempty"`
	ViewerNSS ns.NSS      `json:",omitempty" bson:",omitempty"`
	History   vcs.History `json:",omitempty" bson:",omitempty"`
	CreatedAt time.Time   `json:",omitempty" bson:",omitempty"`
	UpdatedAt *time.Time  `json:",omitempty" bson:",omitempty"`
}

func (item *ItemDescription) UnfoldFromJSON(tagsBytes, urnBytes, historyBytes []byte) error {
	if item == nil {
		return errors.New("nil ItemDescription to be unfolded")
	}

	if len(tagsBytes) > 0 {
		if err := json.Unmarshal(tagsBytes, &item.Tags); err != nil {
			return errors.Wrapf(err, "can't unmarshal .Tags (%s)", tagsBytes)
		}
	}

	item.URN = ns.URN(urnBytes)

	// TODO!!! append to item.History

	if len(historyBytes) > 0 {
		if err := json.Unmarshal(historyBytes, &item.History); err != nil {
			return errors.Wrapf(err, "can't unmarshal .History (%s)", historyBytes)
		}
	}

	return nil
}

func (item *ItemDescription) FoldIntoJSON() (tagsBytes, urnBytes, historyBytes []byte, err error) {
	if item == nil {
		return nil, nil, nil, errors.New("nil persons.Item to be folded")
	}

	tagsBytes = []byte{} // to satisfy NOT NULL constraint
	if len(item.Tags) > 0 {
		if tagsBytes, err = json.Marshal(item.Tags); err != nil {
			return nil, nil, nil, errors.Wrapf(err, "can't marshal .Tags (%#v)", item.Tags)
		}
	}

	if len(item.URN) > 0 {
		urnBytes = []byte(item.URN)
	}

	// TODO!!! append to item.History

	historyBytes = []byte{} // to to satisfy NOT NULL constraint
	if len(item.History) > 0 {
		historyBytes, err = json.Marshal(item.History)
		if err != nil {
			return nil, nil, nil, errors.Wrapf(err, "can't marshal .History(%#v)", item.History)
		}
	}

	return tagsBytes, urnBytes, historyBytes, nil
}
