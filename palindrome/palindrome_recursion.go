package palindrome

import (
	"fmt"
)

func isPalindromeUsingRecursion(s string) bool {
	if len(s) <= 1 {
		return true
	}

	return s[0] == s[len(s)-1] && isPalindrome(s[1:len(s)-1])
}

func findPalindromesRecursion(s string) []string {
	var palindromes []string

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			substr := s[i:j]
			if len(substr) >= 3 {
				if isPalindromeUsingRecursion(substr) {
					palindromes = append(palindromes, substr)
				}
			}
		}
	}

	return palindromes
}

func RunWithRecursion(s string) {
	fmt.Println("=============== Run Palindrome Using Recursion ============")
	palindromes := findPalindromesRecursion(s)
	if len(palindromes) > 0 {
		fmt.Println("Palindromes found:")
		for _, palindrome := range palindromes {
			fmt.Println(palindrome)
		}
	} else {
		fmt.Println("Palindromes not found")
	}
	fmt.Println("=============== Finish Palindrome Using Recursion ============\n")
}
