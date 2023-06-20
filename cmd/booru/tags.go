package booru

import (
    "fmt"
    "github.com/spf13/cobra"
    _ "github.com/ikigai-gh/booru-dl/pkg/booru"
)

var tagsCmd = &cobra.Command {
    Use: "tags",
    Short: "Search tags",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println(42)
    },
}

func init() {
    rootCmd.AddCommand(tagsCmd)
}
