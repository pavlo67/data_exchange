package transformer_test_scenarios

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pavlo67/common/common"

	"github.com/pavlo67/data_exchange/components/transformer"
)

func TestOperator(t *testing.T, transformOp transformer.Operator, params common.Map, dataInitial interface{}, firstCheck bool) (copyFinal, finalStat, finalOut interface{}) {

	var err error

	// import/stat initial data ----------------------------------------------------------

	err = transformOp.In(params, dataInitial)
	require.NoError(t, err)

	statInitial, err := transformOp.Stat(nil, params)
	require.NoError(t, err)
	require.NotNil(t, statInitial)

	// data export/import and its comparison with initial one ----------------------------

	dataRepeat, err := transformOp.Out(nil, params)
	require.NoError(t, err)
	require.NotNil(t, dataRepeat)

	if firstCheck {
		require.Equal(t, dataInitial, dataRepeat)
	}

	err = transformOp.In(params, dataRepeat)
	require.NoError(t, err)

	statRepeat, err := transformOp.Stat(nil, params)
	require.NoError(t, err)
	require.NotNil(t, statRepeat)

	if firstCheck {
		require.Equal(t, statInitial, statRepeat)
	}

	// data export/import repeat and its comparison with previous one --------------------

	dataFinal, err := transformOp.Out(nil, params)
	require.NoError(t, err)
	require.NotNil(t, dataFinal)

	require.Equal(t, dataRepeat, dataFinal)

	err = transformOp.In(params, dataFinal)
	require.NoError(t, err)

	statFinal, err := transformOp.Stat(nil, params)
	require.NoError(t, err)
	require.NotNil(t, statFinal)

	require.Equal(t, statRepeat, statFinal)

	copyFinal, err = transformOp.Copy(nil, params)
	require.NoError(t, err)
	require.NotNil(t, copyFinal)

	return copyFinal, statFinal, dataFinal
}
