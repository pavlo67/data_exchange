package structures

import "github.com/pavlo67/common/common/errors"

var _ Pack = &PackAny{}

type PackAny struct {
	*PackDescription
	PackData DataAny
}

func (pack *PackAny) SetDescription(packDescription PackDescription) error {
	if pack == nil {
		return errors.New("no pack to set description")
	}
	pack.PackDescription = &packDescription
	return nil
}

func (pack *PackAny) Description() *PackDescription {
	if pack == nil {
		return nil
	}
	return pack.PackDescription
}

func (pack *PackAny) Data() Data {
	if pack == nil {
		return nil
	}
	return &pack.PackData
}
