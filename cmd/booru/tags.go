package booru

import (
    "github.com/spf13/cobra"
    booru "github.com/ikigai-gh/booru-dl/pkg/booru"
)

var tagsCmd = &cobra.Command {
    Use: "tags",
    Short: "Search tags",
    Run: func(cmd *cobra.Command, args []string) {
        booru.GetTags()
    },
}

func init() {
    rootCmd.AddCommand(tagsCmd)
}
