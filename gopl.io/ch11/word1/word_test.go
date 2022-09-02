package word

import (
    "testing"
    "math/rand"
    "time"
)


func randomPalindrome(rng *rand.Rand) (text string) {
    n := rng.Intn(25)
    runes := make([]rune, n)
    for i := 0; i < (n + 1)/2; i++ {
    r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
    runes[i] = r
    // fmt.Printf("l: %v\n", runes[i])
    runes[n - 1 - i] = r
    // fmt.Printf("p: %v\n", runes[n - 1 - i])
    }
    return string(runes)
}

func TestRandomPalindrome(t *testing.T) {
    seed := time.Now().UTC().UnixNano()
    t.Logf("Seed: %d", seed)
    rng := rand.New(rand.NewSource(seed))
    for i := 0; i < 1000; i++ {
        text := randomPalindrome(rng)
        if !IsPalindrome(text) {
            t.Errorf("IsPalindrome(%q) = false", text) 
        }
    }   
}

func TestPalindrome(t *testing.T) {
    var tests = []struct{
    input string
    want bool
    }{
        { "", true },
        { "ala", true },
        { "napoli", false },
    }

    for _, test := range tests {
        if got := IsPalindrome(test.input); got != test.want {
            t.Errorf("IsPalidrome(%q) = %v", test.input, got) 
        }
    }
}

func TestNonPalindrome(t *testing.T) {
    if (IsPalindrome("football")) {
        t.Error(`IsPalindrome("football") = true`)
    }
}

func TestFrenchPalindrome(t *testing.T) {
    if !IsPalindrome("été") {
        t.Error(`IsPalindrome("été") = false`)
    }
}

func TestCanalPalindrome(t *testing.T) {
    input := "A man, a plan, a canal: Panama"
    if !IsPalindrome(input) {
        t.Errorf(`IsPalindrome(%q) = false`, input)
    }
}

func BenchmarkIsPalindrome(b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsPalindrome("A man, a plan, a canal: Panama")
    }
} 
