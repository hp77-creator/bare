package cmd

import (
	"bare/styles"
	"bare/utils/parser"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "get info about any bare on github",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide name of the template. \nUsage: bare info <user-name>/<repo-name>")
		} else {
			InfoBare(args[0])
		}
	},
}

func InfoBare(bareName string) {
	user, repo, branch := parser.ParseGithubRepo(bareName)
	parser.GetRecipe(user, repo, branch)
	// TODO Add variant description
	fmt.Println(styles.StatusPrompt.Render("Author   :"), parser.BareObj.Author)
	fmt.Println(styles.StatusPrompt.Render("Template :"), parser.BareObj.BareName, parser.BareObj.Version)
	fmt.Println(styles.StatusPrompt.Render("Variants :"), strings.Trim(fmt.Sprint(parser.BareObj.Variants), "[]"))
}
