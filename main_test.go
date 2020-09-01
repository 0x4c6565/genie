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
}
