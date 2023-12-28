package main

import (
	"github.com/spf13/cobra"
	"okiww/go-palindrome/palindrome"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-palindrome",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var palindromeCmd = &cobra.Command{
	Use:   "run-palindrome",
	Short: "A brief description of your application",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		input := "saippuakivikauppias"
		palindrome.RunWithLoop(input)
	},
}

var palindromeRecursionCmd = &cobra.Command{
	Use:   "run-palindrome-recursion",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		input := "saippuakivikauppias"
		palindrome.RunWithRecursion(input)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(palindromeCmd)
	rootCmd.AddCommand(palindromeRecursionCmd)
}

func main() {
	Execute()
}
