package structures

var _ Pack = &PackAny{}

type PackAny struct {
	PackDescription
	PackData interface{}
}

func (pack PackAny) Description() PackDescription {
	return pack.PackDescription
}

func (pack PackAny) Data() interface{} {
	return pack.PackData
}
