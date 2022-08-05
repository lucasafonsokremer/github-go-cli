/*
Copyright © 2022 Lucas Afonso Kremer

*/
package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

// followersCmd represents the followers command
var followersCmd = &cobra.Command{
	Use:   "followers",
	Short: "A followers of",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("followers called")
	},
}
