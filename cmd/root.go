package cmd

import (
	"bare/styles"
	"bare/utils/osutil"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var BarePath string = osutil.GetBarePath()

var rootCmd = &cobra.Command{
	Use:   "bare",
	Short: "manager for your Bare bones",
	Long:  "manager for your Bare bones",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(styles.InitStyle.Render("\n\n\t\t\tWelcome to Bare.\n\t\t\tUsage: bare <command>\n\t\t\tFor all the commands, please type `bare help`\n\n"))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
