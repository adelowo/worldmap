package alien

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAlienMustContainAName(t *testing.T) {
	alien := NewAlien(&City{Name: "oops"})

	require.NotEmpty(t, alien.Name)
}
