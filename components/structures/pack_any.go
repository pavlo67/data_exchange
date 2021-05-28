package structures

import "github.com/pavlo67/common/common/errors"

var _ Pack = &PackAny{}

type PackAny struct {
	*ItemDescription
	PackData DataAny
}

func (pack *PackAny) SetDescription(packDescription ItemDescription) error {
	if pack == nil {
		return errors.New("no pack to set description")
	}
	pack.ItemDescription = &packDescription
	return nil
}

func (pack *PackAny) Description() *ItemDescription {
	if pack == nil {
		return nil
	}
	return pack.ItemDescription
}

func (pack *PackAny) Data() Data {
	if pack == nil {
		return nil
	}
	return &pack.PackData
}
