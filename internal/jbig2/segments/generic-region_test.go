/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package segments

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/zituocn/updf/common"

	"github.com/zituocn/updf/internal/jbig2/bitmap"
	"github.com/zituocn/updf/internal/jbig2/reader"
)

// TestDecodeGenericRegion tests the decode process of the jbig2 Generic Region.
func TestDecodeGenericRegion(t *testing.T) {
	if testing.Verbose() {
		common.SetLogger(common.NewConsoleLogger(common.LogLevelDebug))
	}

	t.Run("AnnexH", func(t *testing.T) {
		t.Run("S-12th", func(t *testing.T) {
			data := []byte{
				// header
				0x00, 0x00, 0x00, 0x0B, 0x27, 0x00, 0x02, 0x00, 0x00, 0x00, 0x23,

				// data
				0x00, 0x00, 0x00, 0x36, 0x00, 0x00, 0x00, 0x2C, 0x00, 0x00, 0x00, 0x04,
				0x00, 0x00, 0x00, 0x0B, 0x00, 0x08, 0x03, 0xFF, 0xFD, 0xFF, 0x02, 0xFE,
				0xFE, 0xFE, 0x04, 0xEE, 0xED, 0x87, 0xFB, 0xCB, 0x2B, 0xFF, 0xAC,
			}

			r := reader.New(data)
			d := &document{}
			h, err := NewHeader(d, r, 0, OSequential)
			require.NoError(t, err)

			assert.Equal(t, uint32(11), h.SegmentNumber)
			assert.Equal(t, uint64(35), h.SegmentDataLength)
			assert.Equal(t, 2, h.PageAssociation)

			sg, err := h.GetSegmentData()
			require.NoError(t, err)

			s, ok := sg.(*GenericRegion)
			require.True(t, ok)

			assert.Equal(t, 44, s.RegionSegment.BitmapHeight)
			assert.Equal(t, 54, s.RegionSegment.BitmapWidth)
			assert.Equal(t, bitmap.CmbOpOr, s.RegionSegment.CombinaionOperator)
			assert.Equal(t, true, s.IsTPGDon)
			assert.Equal(t, byte(0), s.GBTemplate)

			bm, err := s.GetRegionBitmap()
			require.NoError(t, err)

			t.Log(bm.String())
		})

		t.Run("S-5th", func(t *testing.T) {
			data := []byte{
				// Header
				0x00, 0x00, 0x00, 0x04, 0x27, 0x00, 0x01, 0x00, 0x00, 0x00, 0x2C,

				// Data part
				0x00, 0x00, 0x00, 0x36, 0x00, 0x00, 0x00, 0x2C, 0x00, 0x00, 0x00,
				0x04, 0x00, 0x00, 0x00, 0x0B, 0x00, 0x01, 0x26, 0xA0, 0x71, 0xCE,
				0xA7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF8, 0xF0,
			}
			r := reader.New(data)
			d := &document{}
			h, err := NewHeader(d, r, 0, OSequential)
			require.NoError(t, err)

			assert.Equal(t, uint32(4), h.SegmentNumber)
			assert.Equal(t, uint64(44), h.SegmentDataLength)
			assert.Equal(t, 1, h.PageAssociation)

			gs, err := h.GetSegmentData()
			require.NoError(t, err)

			s, ok := gs.(*GenericRegion)
			require.True(t, ok)

			assert.Equal(t, 44, s.RegionSegment.BitmapHeight)
			assert.Equal(t, 54, s.RegionSegment.BitmapWidth)
			assert.Equal(t, bitmap.CmbOpOr, s.RegionSegment.CombinaionOperator)
			assert.Equal(t, true, s.IsMMREncoded)

			b, err := s.GetRegionBitmap()
			require.NoError(t, err)

			t.Log(b.String())
		})
	})
}
