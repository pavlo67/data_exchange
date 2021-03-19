package transformer

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/selectors"

	"github.com/pavlo67/data_exchange/components/structures"
)

func TestOperator(t *testing.T, transformOp Operator, params common.Map, packInitial structures.Pack, firstCheck bool) (copyFinal, statFinal interface{}, outFinal structures.Pack) {

	var err error

	// get selector from initial data ----------------------------------------------------

	require.NotNil(t, packInitial)

	packURN := packInitial.Description().URN
	require.NotEmpty(t, packURN)

	selector := selectors.Term{
		Key:    structures.InPack,
		Values: packURN,
	}

	// import/stat initial data ----------------------------------------------------------

	err = transformOp.In(packInitial, params)
	require.NoError(t, err)

	statInitial, err := transformOp.Stat(&selector, params)
	require.NoError(t, err)
	require.NotNil(t, statInitial)

	// data export/import and its comparison with initial one ----------------------------

	packRepeat, err := transformOp.Out(&selector, params)
	require.NoError(t, err)
	require.NotNil(t, packRepeat)

	if firstCheck {
		require.Equal(t, packInitial, packRepeat)
	}

	err = transformOp.In(packRepeat, params)
	require.NoError(t, err)

	statRepeat, err := transformOp.Stat(&selector, params)
	require.NoError(t, err)
	require.NotNil(t, statRepeat)

	if firstCheck {
		require.Equal(t, statInitial, statRepeat)
	}

	// data export/import repeat and its comparison with previous one --------------------

	packFinal, err := transformOp.Out(&selector, params)
	require.NoError(t, err)
	require.NotNil(t, packFinal)

	require.Equal(t, packRepeat, packFinal)

	err = transformOp.In(packFinal, params)
	require.NoError(t, err)

	statFinal, err = transformOp.Stat(&selector, params)
	require.NoError(t, err)
	require.NotNil(t, statFinal)

	require.Equal(t, statRepeat, statFinal)

	copyFinal, err = transformOp.Copy(&selector, params)
	require.NoError(t, err)
	require.NotNil(t, copyFinal)

	return copyFinal, statFinal, packFinal
}
