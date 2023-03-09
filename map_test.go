package alien

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRightNumberOfAliensCreated(t *testing.T) {

	worldMap, err := NewMap("testdata/map.txt")
	require.NoError(t, err)

	require.Len(t, worldMap.Aliens, 0)

	worldMap.PopulateMapWithAliens(1000)

	require.Len(t, worldMap.Aliens, 1000)
}

func TestCityMustHaveAtLeastOneConnection(t *testing.T) {

	_, err := NewMap("testdata/invalidmap.txt")
	require.Error(t, err)

	require.Equal(t, "a valid city must have at least one connection", err.Error())
}

func TestCityCountMustBeCorrect(t *testing.T) {

	worldMap, err := NewMap("testdata/map.txt")
	require.NoError(t, err)

	require.Len(t, worldMap.Cities, 4)
}

func BenchmarkNewWorldFileFromReader1000(b *testing.B) {

	s := `Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
Oops north=Bar east=Foo
Bee north=Foo east=Baz
`

	for i := 0; i < b.N; i++ {

		r := strings.NewReader(s)

		_, err := NewWorldWideMapFromReader(r)
		if err != nil {
			b.FailNow()
		}
	}
}
