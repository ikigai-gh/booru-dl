package booru

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

// Version is overrided via ldflags. See Makefile in root directory
var version = ""

var rootCmd = &cobra.Command {
    Use: "booru-dl",
    Version: version,
    Short: "booru-dl - a simple CLI to crawl imageboards",
    Long: "booru-dl rocks!",
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
        os.Exit(1)
    }
}
