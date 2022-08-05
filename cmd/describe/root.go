/*
Copyright © 2022 Lucas Afonso Kremer

*/
package describe

import (
	"fmt"

	"github.com/spf13/cobra"
)

// describeCmd represents the describe command
var DescribeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Retrieve detailed information about any object",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You must specify the type of resource to describe. See 'github-go-cli describe -h' for help and examples")
	},
}

func init() {
	DescribeCmd.AddCommand(userCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// describeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// describeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
