package unique

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const ()

func TestUniqueSuccess(t *testing.T) {
	options := Options{
		ShowCount:      false,
		OnlyDuplicates: false,
		OnlyUnique:     false,
		IgnoreCase:     false,
		FieldSkip:      0,
		CharSkip:       0,
	}

	testInput := []string{
		"test string 1",
		"test string 1",
		"teSt sTring 1",
		"teSt STRing 3",
		"test String 3",
		"assv stRing 3",
		"none",
		"test string 1",
		"string 1",
		"test string 1",
		"string 1",
		"test string 5",
	}
	testOutputBasic := []string{
		"test string 1",
		"teSt sTring 1",
		"teSt STRing 3",
		"test String 3",
		"assv stRing 3",
		"none",
		"test string 1",
		"string 1",
		"test string 1",
		"string 1",
		"test string 5",
	}
	testOutputCount := []string{
		"2 test string 1",
		"1 teSt sTring 1",
		"1 teSt STRing 3",
		"1 test String 3",
		"1 assv stRing 3",
		"1 none",
		"1 test string 1",
		"1 string 1",
		"1 test string 1",
		"1 string 1",
		"1 test string 5",
	}
	testOutputIgnore := []string{
		"test string 1",
		"teSt STRing 3",
		"assv stRing 3",
		"none",
		"test string 1",
		"string 1",
		"test string 1",
		"string 1",
		"test string 5",
	}
	testOutputDuplicates := []string{
		"test string 1",
	}
	testOutputUnique := []string{
		"teSt sTring 1",
		"teSt STRing 3",
		"test String 3",
		"assv stRing 3",
		"none",
		"test string 1",
		"string 1",
		"test string 1",
		"string 1",
		"test string 5",
	}
	testOutputField := []string{
		"test string 1",
		"teSt sTring 1",
		"teSt STRing 3",
		"test String 3",
		"assv stRing 3",
		"none",
		"test string 1",
		"test string 5",
	}
	testOutputChar := []string{
		"test string 1",
		"teSt STRing 3",
		"none",
		"test string 1",
		"string 1",
		"test string 1",
		"string 1",
		"test string 5",
	}
	require.Equal(t, testOutputBasic, CheckUniq(testInput, options), "Comparing basic usage")

	options.ShowCount = true
	require.Equal(t, testOutputCount, CheckUniq(testInput, options), "Comparing count usage")

	options.ShowCount = false
	options.IgnoreCase = true
	require.Equal(t, testOutputIgnore, CheckUniq(testInput, options), "Comparing ignore usage")

	options.IgnoreCase = false
	options.OnlyDuplicates = true
	require.Equal(t, testOutputDuplicates, CheckUniq(testInput, options), "Comparing duplicates usage")

	options.OnlyDuplicates = false
	options.OnlyUnique = true
	require.Equal(t, testOutputUnique, CheckUniq(testInput, options), "Comparing unique usage")

	options.OnlyUnique = false
	options.FieldSkip = 1
	require.Equal(t, testOutputField, CheckUniq(testInput, options), "Comparing field usage")

	options.FieldSkip = 0
	options.CharSkip = 8
	require.Equal(t, testOutputChar, CheckUniq(testInput, options), "Comparing char usage")

}

func TestUniqueFail(t *testing.T) {
	options := Options{
		ShowCount:      true,
		OnlyDuplicates: true,
		OnlyUnique:     true,
		IgnoreCase:     false,
		FieldSkip:      0,
		CharSkip:       0,
	}
	testInput1 := []string{
		"test string 1",
		"test string 1",
		"test string 3",
		"test string 3",
		"test string 1",
		"test string 1",
		"test string 5",
	}
	testOutput := []string{}
	testInput2 := []string{}

	require.Equal(t, testOutput, CheckUniq(testInput1, options), "Check invalid options")
	options.OnlyUnique = false
	options.OnlyDuplicates = false
	require.Equal(t, testOutput, CheckUniq(testInput2, options), "Check invalid input")

}
