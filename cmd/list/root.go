/*
Copyright © 2022 Lucas Afonso Kremer

*/
package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve user profile information such as name, location, public repos, followers, following",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You must specify the type of resource to list. See 'github-go-cli list -h' for help and examples")
	},
}

func init() {
	ListCmd.AddCommand(followersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
