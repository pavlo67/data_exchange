package structures

import "reflect"

// DataAny -------------------------------------------------

var _ Data = &DataAny{}

func NewDataAny(data interface{}) DataAny {
	return DataAny{data}
}

type DataAny struct {
	data interface{}
}

func (dataAny *DataAny) IsEqualTo(dataAnother interface{}) (bool, error) {
	if dataAny == nil {
		// TODO???
		return dataAnother == nil, nil
	}
	return reflect.DeepEqual(dataAny.data, dataAnother), nil
}

func (dataAny *DataAny) Value() interface{} {
	if dataAny == nil {
		// TODO???
		return nil
	}
	return dataAny.data
}
