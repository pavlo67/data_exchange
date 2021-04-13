package structures

var _ Pack = &PackAny{}

type PackAny struct {
	PackDescription
	PackData DataAny
}

func (pack PackAny) Description() PackDescription {
	return pack.PackDescription
}

func (pack PackAny) Data() Data {
	return &pack.PackData
}
