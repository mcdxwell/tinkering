package cmd

import (
	"fmt"
	"strings"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Get desired twitch stream",
	Long:  `Get desired twitch stream`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, username := range args {
				getStream(username)
			}
		} else {
			fmt.Println("Please provide a twitch username")
		}
	},
}

func init() {
	rootCmd.AddCommand(streamCmd)
}

func getStream(username string) {

	var b strings.Builder
	b.WriteString("https://www.twitch.tv/")
	b.WriteString(username)
	fmt.Println("Opening: ", b.String())
	browser.OpenURL(b.String())
	fmt.Println("Opened")
}
