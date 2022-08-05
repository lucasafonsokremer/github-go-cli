/*
Copyright © 2022 Lucas Afonso Kremer

*/
package describe

import (
	"fmt"

	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "describe info about some user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("user called")
	},
}
