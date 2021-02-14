package exchange

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOperator(t *testing.T, exchangeOp Operator, filename string) {
	data, err := ioutil.ReadFile(filename)
	require.NoError(t, err)
	require.True(t, len(data) > 0)

	path := filepath.Dir(filename)

	_, err = exchangeOp.Import(data, path)
	require.NoError(t, err)

	dataExchanged, _, err := exchangeOp.Export(path)
	require.NoError(t, err)
	require.Equal(t, data, dataExchanged)

}
