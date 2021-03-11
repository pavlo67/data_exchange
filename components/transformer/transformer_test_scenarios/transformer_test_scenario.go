package transformer_test_scenarios

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pavlo67/common/common"

	"github.com/pavlo67/data_exchange/components/transformer"
)

func TestOperator(t *testing.T, transformOp transformer.Operator, params common.Map, dataInitial interface{}, firstCheck bool) {

	// import/stat initial data ----------------------------------------------------------

	err := transformOp.Reset()
	require.NoError(t, err)

	err = transformOp.In(nil, params, dataInitial)
	require.NoError(t, err)

	statInitial, err := transformOp.Stat(nil, params)
	require.NoError(t, err)
	require.NotNil(t, statInitial)

	// data export/import and its comparison with initial one ----------------------------

	dataCopy, err := transformOp.Out(nil, params)
	require.NoError(t, err)
	require.NotNil(t, dataCopy)

	if firstCheck {
		require.Equal(t, dataInitial, dataCopy)
	}

	err = transformOp.Reset()
	require.NoError(t, err)

	err = transformOp.In(nil, params, dataCopy)
	require.NoError(t, err)

	statCopy, err := transformOp.Stat(nil, params)
	require.NoError(t, err)
	require.NotNil(t, statCopy)

	if firstCheck {
		require.Equal(t, statInitial, statCopy)
	}

	// data export/import repeat and its comparison with previous one --------------------

	dataCopyRepeat, err := transformOp.Out(nil, params)
	require.NoError(t, err)
	require.NotNil(t, dataCopyRepeat)

	require.Equal(t, dataCopy, dataCopyRepeat)

	err = transformOp.Reset()
	require.NoError(t, err)

	err = transformOp.In(nil, params, dataCopyRepeat)
	require.NoError(t, err)

	statCopyRepeat, err := transformOp.Stat(nil, params)
	require.NoError(t, err)
	require.NotNil(t, statCopyRepeat)

	require.Equal(t, statCopy, statCopyRepeat)

}
