package word

import "unicode"

func IsPalindrome(word string) bool {
    var letters = make([]rune, 0, len(word))
    for _, r := range word {
        if unicode.IsLetter(r) {
            letters = append(letters, unicode.ToLower(r))
        }
    }

    n := len(letters)/2 
    for i := 0; i < n; i++ {
        if (letters[i] != letters[len(letters) - i - 1]) {
            return false
        }    
    }
    return true
}
