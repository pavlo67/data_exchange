package transfer_records_types01

import (
	"fmt"
	"time"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/auth"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/selectors"

	"github.com/pavlo67/data/components/structures"
	"github.com/pavlo67/data/components/transfer"

	"github.com/pavlo67/data/entities/records"

	"github.com/pavlo67/data/exchange/types01"
)

var _ transfer.Operator = &transferrecordsTypes01{}

type transferrecordsTypes01 struct {
	recordsOp records.Operator
	identity  *auth.Identity
}

const onNew = "on transferrecordsTypes01.New(): "

func New(recordsOp records.Operator, identity *auth.Identity) (transfer.Operator, error) {
	if recordsOp == nil {
		return nil, errors.New(onNew + ": no records.Operator")
	}
	//if recordsOpCleaner == nil {
	//	return nil, errors.New(onNew + ": no records.Cleaner")
	//}

	return &transferrecordsTypes01{
		recordsOp: recordsOp,
		identity:  identity,
	}, nil
}

func (transferOp *transferrecordsTypes01) Name() string {
	return string(InterfaceKey)
}

const onIn = "on transferrecordsTypes01.In(): "

func (transferOp *transferrecordsTypes01) In(pack structures.Pack, params common.Map) error {
	if pack == nil {
		return errors.New(onIn + "nil pack to import")
	}

	data := pack.Data().Value()

	var records01 []types01.Record

	switch v := data.(type) {
	case []types01.Record:
		records01 = v
	case *[]types01.Record:
		if v != nil {
			records01 = *v
		}
	default:
		return fmt.Errorf("wrong pack.Data() to import: %#v", data)
	}

	// TODO!!! clear old pack records

	for i, record01 := range records01 {
		embedded := make([]records.Content, len(record01.Embedded))
		for i, c := range record01.Embedded {
			embedded[i] = records.Content{
				Title:   c.Title,
				Summary: c.Summary,
				Type:    c.Type,
				Data:    c.Data,
			}
		}

		recordsItem := records.Item{
			Content: records.Content{
				Title:   record01.Title,
				Summary: record01.Summary,
				Type:    record01.Type,
				Data:    record01.Data,
			},
			Embedded:        embedded,
			ItemDescription: record01.ItemDescription,
		}

		if _, err := transferOp.recordsOp.Save(recordsItem, transferOp.identity); err != nil {
			return fmt.Errorf(onIn+": can't save item (%d / %#v), got %s", i, recordsItem, err)
		}
	}

	return nil
}

const onOut = "on transferrecordsTypes01.Out(): "

func (transferOp *transferrecordsTypes01) Out(selector *selectors.Term, params common.Map) (structures.Pack, error) {

	// TODO!!! create .Description
	records01Pack := structures.PackAny{}

	var records01 []types01.Record

	recordsItems, err := transferOp.recordsOp.List(selector, transferOp.identity)
	if err != nil {
		return nil, fmt.Errorf(onOut+": can't list items (%#v), got %s", selector, err)
	}

	for _, recordsItem := range recordsItems {

		// .CreatedAt & .UpdatedAt are local properties and can't be exported in .History only
		recordsItem.ItemDescription.CreatedAt = time.Time{}
		recordsItem.ItemDescription.UpdatedAt = nil

		embedded := make([]types01.Content, len(recordsItem.Embedded))
		for i, c := range recordsItem.Embedded {
			embedded[i] = types01.Content{
				Title:   c.Title,
				Summary: c.Summary,
				Type:    c.Type,
				Data:    c.Data,
			}
		}

		// TODO!!! set URN (and save using transferOp) if absent
		records01 = append(records01, types01.Record{
			Content: types01.Content{
				Title:   recordsItem.Title,
				Summary: recordsItem.Summary,
				Type:    recordsItem.Type,
				Data:    recordsItem.Data,
			},
			Embedded:        embedded,
			ItemDescription: recordsItem.ItemDescription,
		})

	}

	records01Pack.PackData = structures.NewDataAny(records01)

	return &records01Pack, nil
}

const onStat = "on transferrecordsTypes01.Stat(): "

func (transferOp *transferrecordsTypes01) Stat(selector *selectors.Term, params common.Map) (interface{}, error) {
	return transferOp.recordsOp.Stat(selector, transferOp.identity)
}

const onCopy = "on transferrecordsTypes01.Copy(): "

func (transferOp *transferrecordsTypes01) Copy(selector *selectors.Term, params common.Map) (data interface{}, err error) {
	recordsItems, err := transferOp.recordsOp.List(selector, transferOp.identity)
	if err != nil {
		return nil, fmt.Errorf(onCopy+": can't list items (%#v), got %s", selector, err)
	}

	return recordsItems, nil
}
