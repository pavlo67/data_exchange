package extractor

import "github.com/pavlo67/common/common/config"

type Operator interface {
	Draft(access config.Access, pathTo string) (fileTo string, err error)
	Convert(access config.Access, to interface{}) (result interface{}, err error)
}
