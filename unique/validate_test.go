package unique

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const ()

func TestValidateSuccess(t *testing.T) {
	options1 := Options{
		ShowCount:      false,
		OnlyDuplicates: false,
		OnlyUnique:     false,
	}
	options2 := Options{
		ShowCount:      true,
		OnlyDuplicates: false,
		OnlyUnique:     false,
	}
	options3 := Options{
		ShowCount:      false,
		OnlyDuplicates: true,
		OnlyUnique:     false,
	}
	options4 := Options{
		ShowCount:      false,
		OnlyDuplicates: false,
		OnlyUnique:     true,
	}

	require.True(t, options1.Validate(), "Check concurrent flags false")
	require.True(t, options2.Validate(), "Check concurrent flags false showCount true")
	require.True(t, options3.Validate(), "Check concurrent flags false onlyDuplicates true")
	require.True(t, options4.Validate(), "Check concurrent flags false onlyUnique true")
}

func TestValidateFail(t *testing.T) {
	options1 := Options{
		ShowCount:      true,
		OnlyDuplicates: true,
		OnlyUnique:     true,
	}
	options2 := Options{
		ShowCount:      true,
		OnlyDuplicates: true,
		OnlyUnique:     false,
	}
	options3 := Options{
		ShowCount:      false,
		OnlyDuplicates: true,
		OnlyUnique:     true,
	}
	options4 := Options{
		ShowCount:      true,
		OnlyDuplicates: false,
		OnlyUnique:     true,
	}

	require.False(t, options1.Validate(), "Check concurrent flags true")
	require.False(t, options2.Validate(), "Check concurrent flags true onlyUnique false")
	require.False(t, options3.Validate(), "Check concurrent flags true showCount false")
	require.False(t, options4.Validate(), "Check concurrent flags true onlyDuplicates false")
}
