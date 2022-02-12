package unique

import (
	"fmt"
	"strings"
)

type Options struct {
	InputFile      string
	OutputFile     string
	ShowCount      bool
	OnlyDuplicates bool
	OnlyUnique     bool
	IgnoreCase     bool
	FieldSkip      int
	CharSkip       int
}

func (options *Options) Validate() bool {
	switch {
	case options.ShowCount && options.OnlyDuplicates:
		return false
	case options.ShowCount && options.OnlyUnique:
		return false
	case options.OnlyUnique && options.OnlyDuplicates:
		return false
	}
	return true
}

func CheckUniq(input []string, options Options) (output []string) {

	const (
		initIndex = iota
		duplicateStartIndex
		duplicateThreshold
	)
	currentDuplicateString := input[initIndex]
	duplicateNum := initIndex

	// define equal func by case
	var isEqualCase func(str1, str2 string) bool
	if options.IgnoreCase {
		isEqualCase = strings.EqualFold
	} else {
		isEqualCase = func(str1, str2 string) bool { return str1 == str2 }
	}

	// define skip func
	skipField := func(strPtr *string) {
		// at least one word must exist
		if strSlice := strings.Fields(*strPtr); options.FieldSkip < len(strSlice)-1 {
			*strPtr = strings.Join(strSlice[options.FieldSkip:], " ")
			*strPtr = strings.TrimSpace(*strPtr)
		}
	}
	// define skip chars
	skipChars := func(strPtr *string) {
		str := *strPtr
		if options.CharSkip < len(str)-1 {
			*strPtr = str[options.CharSkip:]
		}
	}
	// define equal func by skipped chars
	isEqualStrings := func(str1, str2 string) bool {
		if options.FieldSkip > 0 {
			skipField(&str1)
			skipField(&str2)
		}
		if options.CharSkip > 0 {
			skipChars(&str1)
			skipChars(&str2)
		}
		return isEqualCase(str1, str2)
	}
	//main handler
	for i, str := range input {
		if isEqualStrings(str, currentDuplicateString) {
			duplicateNum++
			if i != len(input)-1 {
				continue
			}
		}
		switch {
		case options.OnlyDuplicates:
			if duplicateNum >= duplicateThreshold {
				output = append(output, currentDuplicateString)
			}
		case options.OnlyUnique:
			if duplicateNum < duplicateThreshold {
				output = append(output, currentDuplicateString)
			}
		case options.ShowCount:
			output = append(output, fmt.Sprintf("%d %s", duplicateNum, currentDuplicateString))
		default:
			output = append(output, currentDuplicateString)
		}

		duplicateNum = duplicateStartIndex
		currentDuplicateString = str

		if i == len(input)-1 && duplicateNum == duplicateStartIndex && !options.OnlyDuplicates {
			output = append(output, str)
		}
	}
	return
}
