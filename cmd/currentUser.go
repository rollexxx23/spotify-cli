package cmd

import (
	"fmt"

	"github.com/rollexxx23/spotify-cli/utils"
	"github.com/spf13/cobra"
)

// currentUserCmd represents the currentUser command
var currentUserCmd = &cobra.Command{
	Use:   "currentUser",
	Short: "Get Name of the current user",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		api, err := utils.GetSpotifyAPI()
		if err != nil {
			fmt.Println(err)
		} else {
			user, err := api.GetUserProfile()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Logged in as %s.\n", user.DisplayName)
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(currentUserCmd)

}
