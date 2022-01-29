package main

import (
	"fmt"
	"strings"

	"github.com/pkg/browser"
)

// import "github.com/spf13/cobra"
// To-Do:
// Build a CLI tool to open:
// twitch popout chat
// twitch channel
// and to open urls in general
func PopOut(username string) {

	var b strings.Builder
	b.WriteString("https://www.twitch.tv/popout/")
	b.WriteString(username)
	b.WriteString("/chat?popout=")
	fmt.Println("Opening: ", b.String())
	browser.OpenURL(b.String())
	fmt.Println("Opened")
}

func main() {
	PopOut("Davz")
}
