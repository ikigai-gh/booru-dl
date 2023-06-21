package booru

import (
    "github.com/spf13/cobra"
    "github.com/ikigai-gh/booru-dl/pkg/booru"
)

var tags, urlsFile string
var large bool
var maxPages int
var postsCmd = &cobra.Command {
    Use: "posts",
    Short: "Search posts",
    Run: func(cmd *cobra.Command, args []string) {
        booru.GetPosts(tags, large, urlsFile, maxPages)
    },
}

func init() {
    postsCmd.Flags().StringVarP(&tags, "tags", "t", "", "List of space separated tags")
    postsCmd.Flags().StringVarP(&urlsFile, "file", "f", "", "A path to file that contains list of urls to download")
    postsCmd.Flags().BoolVarP(&large, "large", "l", false, "Whether to download large images")
    postsCmd.Flags().IntVarP(&maxPages, "pages", "p", 50, "Maximum number of pages to crawl")
    rootCmd.AddCommand(postsCmd)
}
