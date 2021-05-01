package transfer

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/selectors"

	"github.com/pavlo67/data/components/structures"
)

func TestOperator(t *testing.T, transferOp Operator, params common.Map, packInitial structures.Pack,
	checkFirstCopy, checkPackDescription bool) (copyFinal, statFinal interface{}, outFinal structures.Pack) {

	var err error

	// get selector from initial data ----------------------------------------------------

	require.NotNil(t, packInitial)

	packURN := packInitial.Description().URN
	require.NotEmpty(t, packURN)

	selector := selectors.Term{
		Key:    structures.InPack,
		Values: packURN,
	}

	initialDescription := packInitial.Description()
	require.NotNil(t, initialDescription)

	// import/stat initial data ----------------------------------------------------------

	err = transferOp.In(packInitial, params)
	require.NoError(t, err)

	statInitial, err := transferOp.Stat(&selector, params)
	require.NoError(t, err)
	require.NotNil(t, statInitial)

	// l.Fatalf("%#v\n\n--> %#v / %#v", packInitial.Data(), statInitial, selector)

	// data export/import and its comparison with initial one ----------------------------

	packRepeat, err := transferOp.Out(&selector, params)
	require.NoError(t, err)
	require.NotNil(t, packRepeat)

	if checkPackDescription {
		require.Equal(t, initialDescription, packRepeat.Description())
	} else {
		err = packRepeat.SetDescription(*initialDescription)
		require.NoError(t, err)
	}

	if checkFirstCopy {
		require.Equal(t, packInitial.Data(), packRepeat.Data())
	}

	err = transferOp.In(packRepeat, params)
	require.NoError(t, err)

	statRepeat, err := transferOp.Stat(&selector, params)
	require.NoError(t, err)
	require.NotNil(t, statRepeat)

	if checkFirstCopy {
		require.Equal(t, statInitial, statRepeat)
	}

	// data export/import repeat and its comparison with previous one --------------------

	packFinal, err := transferOp.Out(&selector, params)
	require.NoError(t, err)
	require.NotNil(t, packFinal)

	if checkPackDescription {
		require.Equal(t, packRepeat.Description(), packFinal.Description())
	} else {
		err = packFinal.SetDescription(*packRepeat.Description())
		require.NoError(t, err)
	}
	require.Equal(t, packRepeat.Data(), packFinal.Data())

	err = transferOp.In(packFinal, params)
	require.NoError(t, err)

	statFinal, err = transferOp.Stat(&selector, params)
	require.NoError(t, err)
	require.NotNil(t, statFinal)

	require.Equal(t, statRepeat, statFinal)

	copyFinal, err = transferOp.Copy(&selector, params)
	require.NoError(t, err)
	require.NotNil(t, copyFinal)

	return copyFinal, statFinal, packFinal

}
