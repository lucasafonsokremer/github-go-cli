/*
Copyright © 2022 Lucas Afonso Kremer

*/
package describe

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/lucasafonsokremer/github-go-cli/cmd/util/service"
	"github.com/lucasafonsokremer/github-go-cli/cmd/util/types"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "describe info about some user",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("please, specify a username")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		client := service.CreateClient(fmt.Sprintf("/users/%s", args[0]))
		user, err := client.GetUser()

		if err != nil {
			_ = fmt.Errorf("%v", err)
			os.Exit(1)
		}

		displayInfo(user)
	},
}

func displayInfo(user types.User) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
	_, _ = fmt.Fprintln(w, "Name\tLocation")
	for _, u := range user {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("%s\t%s", u.Name, u.Location))
	}
	_ = w.Flush()
}
