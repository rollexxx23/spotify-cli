package cmd

import (
	"fmt"

	"github.com/rollexxx23/spotify-cli/utils"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "For Logging in",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called...")
		token, err := utils.Login()
		if err != nil {
			fmt.Println("Failed to login")

		} else {
			err = utils.SaveTokens(token)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Login Successful")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

}
