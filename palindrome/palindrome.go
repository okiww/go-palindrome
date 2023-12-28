package palindrome

import "fmt"

func isPalindrome(s string) bool {
	// check
	last := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[last] {
			return false
		}
		last--
	}

	return true
}

func findPalindromes(s string) []string {
	var palindromes []string

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			substr := s[i:j]
			if len(substr) >= 3 {
				if isPalindrome(substr) {
					palindromes = append(palindromes, substr)
				}
			}
		}
	}

	return palindromes
}
func RunWithLoop(s string) {
	fmt.Println("=============== Run Palindrome Using Loop ============")
	palindromes := findPalindromes(s)
	if len(palindromes) > 0 {
		fmt.Println("Palindromes found:")
		for _, palindrome := range palindromes {
			fmt.Println(palindrome)
		}
	} else {
		fmt.Println("Palindromes not found")
	}
	fmt.Println("=============== Finish Palindrome Using Loop ============\n")
}
