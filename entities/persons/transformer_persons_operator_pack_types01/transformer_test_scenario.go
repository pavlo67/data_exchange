package transformer_persons_operator_pack_types01

import (
	"testing"

	"github.com/pavlo67/data/components/transformer"

	"github.com/stretchr/testify/require"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/selectors"

	"github.com/pavlo67/data/components/structures"
)

func TestOperator(t *testing.T, transformOp transformer.Operator, params common.Map, packInitial structures.Pack, checkFirstCopy, checkPackDescription bool) (copyFinal,
	statFinal interface{},
	outFinal structures.Pack) {

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

	// l.Fatalf("%#v\n\n--> %#v / %#v", packInitial.Data(), statInitial, selector)

	// data export/import and its comparison with initial one ----------------------------

	packRepeat, err := transformOp.Out(&selector, params)
	require.NoError(t, err)
	require.NotNil(t, packRepeat)

	if checkFirstCopy {
		if checkPackDescription {
			require.Equal(t, packInitial.Description(), packRepeat.Description())
		}
		require.Equal(t, packInitial.Data(), packRepeat.Data())
	}

	err = transformOp.In(packRepeat, params)
	require.NoError(t, err)

	statRepeat, err := transformOp.Stat(&selector, params)
	require.NoError(t, err)
	require.NotNil(t, statRepeat)

	if checkFirstCopy {
		require.Equal(t, statInitial, statRepeat)
	}

	// data export/import repeat and its comparison with previous one --------------------

	packFinal, err := transformOp.Out(&selector, params)
	require.NoError(t, err)
	require.NotNil(t, packFinal)

	if checkPackDescription {
		require.Equal(t, packRepeat.Description(), packFinal.Description())
	}
	require.Equal(t, packRepeat.Data(), packFinal.Data())
	// require.Equal(t, packRepeat, packFinal)

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
