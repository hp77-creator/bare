package cmd

import (
	"bare/parser"
	"bare/styles"
	"bare/utils"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(touchCmd)
}

var touchCmd = &cobra.Command{
	Use:   "touch",
	Short: "add a single boilerplate file from a bare",
	Run: func(cmd *cobra.Command, args []string) {
		touchFile(args[0], args[1])
	},
}

func touchFile(touchName, newName string) {
	recipePath := "./recipe.json"
	parser.Parser(recipePath)
	touchMap := parser.BareObj.Touch
	barePath := parser.BareObj.BarePath
	currDir, _ := os.Getwd()
	if touchMap[touchName] == "" {
		fmt.Println(styles.InitError.Render("No such touch present"))
		os.Exit(1)
	} else {
		if utils.Exists(currDir + "/" + newName) {
			fmt.Println("Already exsists !!")
			inp := "n" // warning
			reader := bufio.NewReader(os.Stdin)
			fmt.Print(styles.Warning.Render("Do you want to override " + newName + " (y/N) > "))
			inp, _ = reader.ReadString('\n')
			inp = strings.Trim(inp, " ")
			if inp == "y" || inp == "Y" {
				utils.CopyFileDirectory((barePath + "/" + touchMap[touchName]), currDir+"/"+newName)
			} else {
				os.Exit(0)
			}
		} else {
			utils.CopyFileDirectory((barePath + "/" + touchMap[touchName]), currDir+"/"+newName)
		}
	}
}