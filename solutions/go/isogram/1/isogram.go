package isogram

import "strings"

func IsIsogram(word string) bool {
	replacer := strings.NewReplacer(" ", "", "-", "")
    
    lowered := strings.ToLower(word)
    replaced := replacer.Replace(lowered)

	wordMap := make(map[rune]struct{})

    for _, char := range replaced {
        _, exists := wordMap[char]

        if exists {
            return false
        }

        wordMap[char] = struct{}{}
    }

    return true
}
