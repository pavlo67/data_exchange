package transfer_persons_types01

import (
	"fmt"
	"time"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/auth"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/selectors"

	"github.com/pavlo67/data/components/structures"
	"github.com/pavlo67/data/components/transfer"

	"github.com/pavlo67/data/entities/persons"

	"github.com/pavlo67/data/exchange/types01"
)

var _ transfer.Operator = &transferPersonsTypes01{}

type transferPersonsTypes01 struct {
	personsOp persons.Operator
	identity  *auth.Identity
}

const onNew = "on transferPersonsTypes01.New(): "

func New(personsOp persons.Operator, identity *auth.Identity) (transfer.Operator, error) {
	if personsOp == nil {
		return nil, errors.New(onNew + ": no persons.Operator")
	}
	//if personsOpCleaner == nil {
	//	return nil, errors.New(onNew + ": no persons.Cleaner")
	//}

	return &transferPersonsTypes01{
		personsOp: personsOp,
		identity:  identity,
	}, nil
}

func (transferOp *transferPersonsTypes01) Name() string {
	return string(InterfaceKey)
}

const onIn = "on transferPersonsTypes01.In(): "

func (transferOp *transferPersonsTypes01) In(pack structures.Pack, params common.Map) error {
	if pack == nil {
		return errors.New(onIn + "nil pack to import")
	}

	data := pack.Data().Value()

	var persons01 []types01.Person

	switch v := data.(type) {
	case []types01.Person:
		persons01 = v
	case *[]types01.Person:
		if v != nil {
			persons01 = *v
		}
	default:
		return fmt.Errorf("wrong pack.Data() to import: %#v", data)
	}

	// TODO!!! clear old pack records

	for i, person01 := range persons01 {
		personItem := persons.Item{
			Identity: auth.Identity{
				Nickname: person01.Nickname,
				Roles:    person01.Roles,
			},
			ItemDescription: person01.ItemDescription,
			InPackURN:       pack.Description().URN,
		}
		personItem.SetCreds(person01.Creds)

		if _, err := transferOp.personsOp.Save(personItem, transferOp.identity); err != nil {
			return fmt.Errorf(onIn+": can't save item (%d / %#v), got %s", i, personItem, err)
		}
	}

	return nil
}

const onOut = "on transferPersonsTypes01.Out(): "

func (transferOp *transferPersonsTypes01) Out(selector *selectors.Term, params common.Map) (structures.Pack, error) {

	// TODO!!! create .Description
	persons01Pack := structures.PackAny{}

	var persons01 []types01.Person

	personsItems, err := transferOp.personsOp.List(selector, transferOp.identity)
	if err != nil {
		return nil, fmt.Errorf(onOut+": can't list items (%#v), got %s", selector, err)
	}

	for _, personsItem := range personsItems {

		// .CreatedAt & .UpdatedAt are local properties and can't be exported in .History only
		personsItem.ItemDescription.CreatedAt = time.Time{}
		personsItem.ItemDescription.UpdatedAt = nil

		// TODO!!! set URN (and save using transferOp) if absent
		persons01 = append(persons01, types01.Person{
			Nickname:        personsItem.Nickname,
			Roles:           personsItem.Roles,
			Creds:           personsItem.Creds(),
			ItemDescription: personsItem.ItemDescription,
		})
	}

	persons01Pack.PackData = structures.NewDataAny(persons01)

	return &persons01Pack, nil
}

const onStat = "on transferPersonsTypes01.Stat(): "

func (transferOp *transferPersonsTypes01) Stat(selector *selectors.Term, params common.Map) (interface{}, error) {
	return transferOp.personsOp.Stat(selector, transferOp.identity)
}

const onCopy = "on transferPersonsTypes01.Copy(): "

func (transferOp *transferPersonsTypes01) Copy(selector *selectors.Term, params common.Map) (data interface{}, err error) {
	personsItems, err := transferOp.personsOp.List(selector, transferOp.identity)
	if err != nil {
		return nil, fmt.Errorf(onCopy+": can't list items (%#v), got %s", selector, err)
	}

	return personsItems, nil
}
