/*
Copyright © 2022 Lucas Afonso Kremer

*/
package describe

import (
	"fmt"
	"os"

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
		userInfo, err := getUserInfo(client)

		if err != nil {
			_ = fmt.Errorf("%v", err)
			os.Exit(1)
		}

		displayInfo(userInfo)
	},
}

func getUserInfo(client service.Client) (types.UserInfo, error) {
	userInfo, err := client.GetUser()

	if err != nil {
		return types.UserInfo{}, err
	}

	return userInfo, err
}

func displayInfo(user types.UserInfo) {
	fmt.Println(user.Name)
	fmt.Println(user.Location)
	fmt.Println(user.PublicRepos)
}
