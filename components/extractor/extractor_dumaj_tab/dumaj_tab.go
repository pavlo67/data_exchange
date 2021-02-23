package extractor_dumaj_tab

import (
	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/config"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/filelib"
	"github.com/pavlo67/data_exchange/components/extractor/extractor_helpers"

	"github.com/pavlo67/data_exchange/components/extractor"
)

var _ extractor.Operator = &extractorDumajTab{}

type extractorDumajTab struct{}

const onNew = "on extractorDumajTab.New(): "

func New() (extractor.Operator, error) {
	extractorOp := extractorDumajTab{}
	return &extractorOp, nil
}

const onDraft = "on extractorDumajTab.Draft(): "

func (exchangeOp *extractorDumajTab) Draft(access config.Access, pathTo string) (fileTo string, err error) {

	correctedPathTo, err := filelib.Dir(pathTo)
	if err != nil {
		return "", errors.CommonError(err, onDraft)
	}

	data, tab, err := extractor_helpers.Tab(access.Path)

	l.Infof("%s [%d bytes / %d lines] --> %s", access.Path, len(data), len(tab), correctedPathTo)

	return "", common.ErrNotImplemented
}

const onConvert = "on extractorDumajTab.Convert()"

func (exchangeOp *extractorDumajTab) Convert(access config.Access, to interface{}) (result interface{}, err error) {

	return nil, common.ErrNotImplemented
}
