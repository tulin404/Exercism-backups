package anagram

import (
    "strings"
    "sort"
    "unicode/utf8"
)

func sortString(str string) string {
	runes := []rune(str)
    
    sort.Slice(runes, func(i, j int) bool {
        return runes[i] < runes[j]
    })

    return string(runes)
}

func Detect(subject string, candidates []string) []string {
	loweredSub := strings.ToLower(subject)
    trimmedSub := strings.TrimSpace(loweredSub)
    sortedSub := sortString(trimmedSub)
    
	anagrams := []string{}

    for _, candidate := range candidates {
		candidate = strings.TrimSpace(candidate)
        
        if utf8.RuneCountInString(candidate) != utf8.RuneCountInString(trimmedSub) {
            continue
        }
        
        loweredCand := strings.ToLower(candidate)

        if loweredCand == loweredSub {
            continue
        }

        sortedCand := sortString(loweredCand)

        if sortedCand == sortedSub {
            anagrams = append(anagrams, candidate)
        }
    }
    
    return anagrams
}
