package transform

import (
	"hackathon-work/model"
	"strings"
	"unicode"
)

func DetectCandidateType(candidate model.Candidate, words []string, wantedType string) (model.Candidate, bool) {
	var isDetected bool
	for _, word := range words {
		if strings.Contains(strings.ToLower(candidate.Position), word) {
			candidate.Type = wantedType
			isDetected = true
			break
		}
	}
	return candidate, isDetected
}

func DetectVacationType(vacation model.NewVacation, words []string, wantedType string) (model.NewVacation, bool) {
	var isDetected bool
	for _, word := range words {
		if strings.Contains(strings.ToLower(vacation.Name), word) {
			vacation.Type = wantedType
			isDetected = true
			break
		}
	}
	return vacation, isDetected
}

func DeleteStartSpaces(raw string) (formatted string) {
	c := 0
	for i := 0; !unicode.IsLetter(rune(raw[i])); i++ {
		c++
	}
	formatted = raw[c:]
	return formatted
}


