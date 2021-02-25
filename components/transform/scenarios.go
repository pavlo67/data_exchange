package transform

import (
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/selectors"
)

const KeyExportFinished errors.Key = "export_finished"

var ErrExportFinished = errors.KeyableError(KeyExportFinished, nil)

const onRun = "on transform.Run()"

func Run(transformOpFrom, transformOpTo Operator, selectorFrom, selectorTo *selectors.Term) error {
	if err := transformOpFrom.Clean(nil); err != nil {
		return errors.CommonError(err, onRun)
	}

	if err := transformOpTo.Clean(nil); err != nil {
		return errors.CommonError(err, onRun)
	}

	for {
		if err := transformOpFrom.Read(selectorFrom); err != nil {
			commonErr := errors.CommonError(err)
			if commonErr.Key() != KeyExportFinished {
				return commonErr.Append(onRun)
			}
			break
		}

		structure, data, err := transformOpFrom.Out(nil)
		if err != nil {
			return errors.CommonError(err, onRun)
		}

		if err := transformOpTo.In(selectorTo, structure, data); err != nil {
			return errors.CommonError(err, onRun)
		}

		if err := transformOpFrom.Save(nil); err != nil {
			return errors.CommonError(err, onRun)
		}
	}

	return nil
}

const onRunByOnePiece = "on transform.RunByOnePiece()"

func RunByOnePiece(transformOpFrom, transformOpTo Operator, selectorFrom, selectorTo *selectors.Term) error {
	if err := transformOpFrom.Clean(nil); err != nil {
		return errors.CommonError(err, onRunByOnePiece)
	}

	if err := transformOpTo.Clean(nil); err != nil {
		return errors.CommonError(err, onRunByOnePiece)
	}

	for {
		if err := transformOpFrom.Read(selectorFrom); err != nil {
			commonErr := errors.CommonError(err)
			if commonErr.Key() != KeyExportFinished {
				return commonErr.Append(onRunByOnePiece)
			}
			break
		}
	}

	structure, data, err := transformOpFrom.Out(nil)
	if err != nil {
		return errors.CommonError(err, onRunByOnePiece)
	}

	if err := transformOpTo.In(selectorTo, structure, data); err != nil {
		return errors.CommonError(err, onRunByOnePiece)
	}

	if err := transformOpFrom.Save(nil); err != nil {
		return errors.CommonError(err, onRunByOnePiece)
	}

	return nil
}
