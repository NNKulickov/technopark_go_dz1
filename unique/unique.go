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

	if len(input) == 0 {
		output = input
		return
	}

	if !options.Validate() {
		output = []string{}
		return
	}
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

	const (
		duplicateInitialIndex = iota
		duplicateStartIndex
		duplicateThreshold
	)
	currentDuplicateString := input[duplicateInitialIndex]
	duplicateNum := duplicateInitialIndex

	for i, str := range input {
		isEqual := isEqualStrings(str, currentDuplicateString)
		isNotLast := i != len(input)-1
		if isEqual {
			duplicateNum++
			if isNotLast {
				continue
			}
		}
		switch {
		case options.OnlyDuplicates && duplicateNum >= duplicateThreshold:
			output = append(output, currentDuplicateString)

		case options.OnlyUnique && !isNotLast && !isEqual && duplicateNum < duplicateThreshold:
			output = append(output, currentDuplicateString)
			output = append(output, str)

		case options.OnlyUnique && duplicateNum < duplicateThreshold:
			output = append(output, currentDuplicateString)

		case options.ShowCount && (isNotLast || isEqual):
			output = append(output, fmt.Sprintf("%d %s", duplicateNum, currentDuplicateString))

		case options.ShowCount:
			output = append(output, fmt.Sprintf("%d %s", duplicateNum, currentDuplicateString))
			output = append(output, fmt.Sprintf("%d %s", duplicateStartIndex, str))

		case !isNotLast && !isEqual && !options.OnlyDuplicates:
			output = append(output, currentDuplicateString)
			output = append(output, str)

		case !options.OnlyDuplicates && !options.OnlyUnique:
			output = append(output, currentDuplicateString)
		}

		duplicateNum = duplicateStartIndex
		currentDuplicateString = str
	}
	return
}
