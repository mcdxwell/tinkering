package cmd

import (
	"fmt"
	"strings"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var popCmd = &cobra.Command{
	Use:   "pop",
	Short: "Get popout chat of desired twitch stream",
	Long:  `Get popout chat of desired twitch stream`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, username := range args {
				popOut(username)
			}
		} else {
			fmt.Println("Please provide a twitch username")
		}
	},
}

func init() {
	rootCmd.AddCommand(popCmd)
}

func popOut(username string) {

	var b strings.Builder
	b.WriteString("https://www.twitch.tv/popout/")
	b.WriteString(username)
	b.WriteString("/chat?popout=")
	fmt.Println("Opening: ", b.String())
	browser.OpenURL(b.String())
	fmt.Println("Opened")
}
