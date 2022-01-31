package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// To-Do:
// 1. Fix the printing of the current date.
// 2. Create a print function where it can be called in each
// of the commands instead of inside the getFunctions.

var dayCmd = &cobra.Command{
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

var monthCmd = &cobra.Command{
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

var yearCmd = &cobra.Command{
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
	rootCmd.AddCommand(dayCmd)
	rootCmd.AddCommand(monthCmd)
	rootCmd.AddCommand(yearCmd)
}

func getCurrDate() time.Time {
	current_date := time.Now().Local()
	fmt.Println("The current date is: ", current_date.Format("01-02-2006"))
	return current_date
}

func whatDay(dd int) {
	desired_dd := getCurrDate().AddDate(0, 0, dd)
	fmt.Printf("The date in %d days: %v", dd, desired_dd.Format("01-02-2006"))
}

func whatMonth(mm int) {
	desired_mm := getCurrDate().AddDate(0, mm, 0)
	fmt.Printf("The date in %d days: %v", mm, desired_mm.Format("01-02-2006"))
}

func whatYear(yyyy int) {
	desired_yyyy := getCurrDate().AddDate(yyyy, 0, 0)
	fmt.Printf("The date in %d days: %v", yyyy, desired_yyyy.Format("01-02-2006"))
}
