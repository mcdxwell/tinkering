package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// To-Do:
// 1. Fix the printing of the current date.
// 2. Create a print function where it can be called in each
// of the commands instead of inside the getFunctions.

var yearsCmd = &cobra.Command{
	Use:   "years",
	Short: "Get desired year in given years.",
	Long:  `Get desired year in given years.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, years := range args {

				y, err := strconv.Atoi(years)
				// make sure d, m, y are not any other type than a positive int
				// note: cannot pass uint into whatDay() function.
				whatYear(y)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		} else {
			fmt.Println("Please provide a number of years")
		}
	},
}

func init() {
	rootCmd.AddCommand(yearsCmd)
}

func whatYear(yyyy int) {
	desired_yyyy := GetCurrDate().AddDate(yyyy, 0, 0)
	fmt.Printf("The date in %d days: %v\n", yyyy, desired_yyyy.Format("01-02-2006"))
}
