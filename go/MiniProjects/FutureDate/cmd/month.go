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

var monthsCmd = &cobra.Command{
	Use:   "months",
	Short: "Get desired month in given months.",
	Long:  `Get desired month in given months.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, months := range args {

				m, err := strconv.Atoi(months)
				// make sure d, m, y are not any other type than a positive int
				// note: cannot pass uint into whatDay() function.
				whatMonth(m)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		} else {
			fmt.Println("Please provide a number of months")
		}
	},
}

func init() {
	rootCmd.AddCommand(monthsCmd)
}

func whatMonth(mm int) {
	desired_mm := GetCurrDate().AddDate(0, mm, 0)
	fmt.Printf("The date in %d days: %v\n", mm, desired_mm.Format("01-02-2006"))
}
