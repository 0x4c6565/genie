package genie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenie_parseMarker(t *testing.T) {
	t.Run("ValidMarker_NoArgs_Parses", func(t *testing.T) {
		g := &Genie{}

		marker, err := g.parseMarker("+genie:testmarker")

		assert.Nil(t, err)
		assert.Equal(t, "testmarker", marker.Name)
		assert.Equal(t, "", marker.Args)
	})

	t.Run("ValidMarker_WithArgs_Parses", func(t *testing.T) {
		g := &Genie{}

		marker, err := g.parseMarker("+genie:testmarker:testargs")

		assert.Nil(t, err)
		assert.Equal(t, "testmarker", marker.Name)
		assert.Equal(t, "testargs", marker.Args)
	})

	t.Run("ValidMarker_WithArgsWithSemicolon_Parses", func(t *testing.T) {
		g := &Genie{}

		marker, err := g.parseMarker("+genie:testmarker:testargs:abc")

		assert.Nil(t, err)
		assert.Equal(t, "testmarker", marker.Name)
		assert.Equal(t, "testargs:abc", marker.Args)
	})

	t.Run("InvalidMarker_MissingMarkerName_ReturnsError", func(t *testing.T) {
		g := &Genie{}

		_, err := g.parseMarker("+unknown:")

		assert.NotNil(t, err)
	})

	t.Run("InvalidMarker_EmptyMarker_ReturnsError", func(t *testing.T) {
		g := &Genie{}

		_, err := g.parseMarker("")

		assert.NotNil(t, err)
	})
}

func TestGenie_isMarker(t *testing.T) {
	t.Run("ValidMarker_ReturnsTrue", func(t *testing.T) {
		g := &Genie{markerPrefix: "genie"}

		isMarker := g.isMarker("+genie:testmarker")

		assert.True(t, isMarker)
	})

	t.Run("InvalidMarker_WrongPrefix_ReturnsFalse", func(t *testing.T) {
		g := &Genie{markerPrefix: "genie"}

		isMarker := g.isMarker("+unknown:testmarker")

		assert.False(t, isMarker)
	})

	t.Run("InvalidMarker_MissingMarkerName_ReturnsFalse", func(t *testing.T) {
		g := &Genie{markerPrefix: "genie"}

		isMarker := g.isMarker("+genie")

		assert.False(t, isMarker)
	})
}
