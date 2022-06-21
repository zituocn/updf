/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package huffman

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/zituocn/updf/common"
)

// TestGetStandardTable tests GetStandardTable function.
func TestGetStandardTable(t *testing.T) {
	setLogger()

	t.Run("Valid", func(t *testing.T) {
		table, err := GetStandardTable(1)
		require.NoError(t, err)

		if assert.NotNil(t, table.RootNode()) {
			t.Logf(table.String())
		}
	})

	t.Run("OutOfRange", func(t *testing.T) {
		t.Run("Negative", func(t *testing.T) {
			_, err := GetStandardTable(-1)
			require.Error(t, err)
		})
		t.Run("Zero", func(t *testing.T) {
			_, err := GetStandardTable(0)
			require.Error(t, err)
		})
		t.Run("GreaterThanLength", func(t *testing.T) {
			_, err := GetStandardTable(len(tables) + 1)
			require.Error(t, err)
		})
	})
}

var alreadySet bool

func setLogger() {
	if testing.Verbose() {
		if !alreadySet {
			common.SetLogger(common.NewConsoleLogger(common.LogLevelDebug))
		}
	}
}
