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

var daysCmd = &cobra.Command{
	Use:   "days",
	Short: "Get desired day in given days.",
	Long:  `Get desired day in given days.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, days := range args {

				d, err := strconv.Atoi(days)
				// make sure d, m, y are not any other type than a positive int
				// note: cannot pass uint into whatDay() function.
				whatDay(d)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		} else {
			fmt.Println("Please provide a number of days")
		}
	},
}

func init() {
	rootCmd.AddCommand(daysCmd)
}

func whatDay(dd int) {
	desired_dd := GetCurrDate().AddDate(0, 0, dd)
	fmt.Printf("The date in %d days: %v\n", dd, desired_dd.Format("01-02-2006"))
}
