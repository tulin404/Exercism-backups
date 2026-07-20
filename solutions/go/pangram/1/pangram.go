package pangram
import "strings"

func IsPangram(input string) bool {
	lowered := strings.ToLower(input)
    replaced := strings.ReplaceAll(lowered, " ", "")

    alphabetString := "abcdefghijklmnopqrstuvwxyz"
    alphabetArray := strings.Split(alphabetString, "")

    hasLetter := true
    
    for _, value := range alphabetArray {
        hasLetter = strings.Contains(replaced, value)

        if !hasLetter {
            return false
        }
    }

    return true
}
