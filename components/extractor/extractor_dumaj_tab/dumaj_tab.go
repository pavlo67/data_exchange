package extractor_dumaj_tab

import (
	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/config"

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

// from internal database
func (exchangeOp *extractorDumajTab) Draft(access config.Access, pathTo string) (fileTo string, err error) {
	//correctedPathTo, err := filelib.Dir(pathTo)
	//if err != nil {
	//	return "", errors.CommonError(err, onNew)
	//}

	//var filename string
	//// TODO read filename from selector
	//
	//data, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	return fmt.Errorf(onRead+": reading %s got %s", filename, err)
	//}
	//
	//var recordsExchangePack exchange_0_1.RecordsExchangePack
	//if err = json.Unmarshal(data, &recordsExchangePack); err != nil {
	//	return fmt.Errorf(onRead+": reading %s got %s", filename, err)
	//}
	//exchangeOp.recordsExchangePack = recordsExchangePack

	return "", common.ErrNotImplemented
}

const onConvert = "on extractorDumajTab.Convert()"

// into internal database
func (exchangeOp *extractorDumajTab) Convert(access config.Access, to interface{}) (result interface{}, err error) {
	//
	//data, err := json.Marshal(exchangeOp.recordsExchangePack)
	//if err != nil {
	//	return fmt.Errorf(onSave+": marshalling data got %s", err)
	//}
	//
	//var filename string
	//// TODO read filename from selector
	//
	//if err = ioutil.WriteFile(filename, data, 0644); err != nil {
	//	return fmt.Errorf(onSave+": writing into %s got %s", filename, err)
	//}

	return nil, common.ErrNotImplemented
}
