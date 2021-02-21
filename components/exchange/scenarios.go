package exchange

import (
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/selectors"
)

const KeyExportFinished errors.Key = "export_finished"

var ErrExportFinished = errors.KeyableError(KeyExportFinished, nil)

const onRun = "on exchange.Run()"

func Run(exchangeOpFrom, exchangeOpTo Operator, selectorFrom, selectorTo *selectors.Term) error {
	if err := exchangeOpFrom.Clean(nil); err != nil {
		return errors.CommonError(err, onRun)
	}

	if err := exchangeOpTo.Clean(nil); err != nil {
		return errors.CommonError(err, onRun)
	}

	for {
		if err := exchangeOpFrom.Read(selectorFrom); err != nil {
			commonErr := errors.CommonError(err)
			if commonErr.Key() != KeyExportFinished {
				return commonErr.Append(onRun)
			}
			break
		}

		structure, data, err := exchangeOpFrom.Export(nil)
		if err != nil {
			return errors.CommonError(err, onRun)
		}

		if err := exchangeOpTo.Import(selectorTo, structure, data); err != nil {
			return errors.CommonError(err, onRun)
		}

		if err := exchangeOpFrom.Save(nil); err != nil {
			return errors.CommonError(err, onRun)
		}
	}

	return nil
}

const onRunByOnePiece = "on exchange.RunByOnePiece()"

func RunByOnePiece(exchangeOpFrom, exchangeOpTo Operator, selectorFrom, selectorTo *selectors.Term) error {
	if err := exchangeOpFrom.Clean(nil); err != nil {
		return errors.CommonError(err, onRunByOnePiece)
	}

	if err := exchangeOpTo.Clean(nil); err != nil {
		return errors.CommonError(err, onRunByOnePiece)
	}

	for {
		if err := exchangeOpFrom.Read(selectorFrom); err != nil {
			commonErr := errors.CommonError(err)
			if commonErr.Key() != KeyExportFinished {
				return commonErr.Append(onRunByOnePiece)
			}
			break
		}
	}

	structure, data, err := exchangeOpFrom.Export(nil)
	if err != nil {
		return errors.CommonError(err, onRunByOnePiece)
	}

	if err := exchangeOpTo.Import(selectorTo, structure, data); err != nil {
		return errors.CommonError(err, onRunByOnePiece)
	}

	if err := exchangeOpFrom.Save(nil); err != nil {
		return errors.CommonError(err, onRunByOnePiece)
	}

	return nil
}
