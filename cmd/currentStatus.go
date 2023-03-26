package cmd

import (
	"fmt"

	"github.com/rollexxx23/spotify-cli/utils"
	"github.com/spf13/cobra"
)

// currentStatusCmd represents the currentStatus command
var currentStatusCmd = &cobra.Command{
	Use:   "currentStatus",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		api, err := utils.GetSpotifyAPI()
		if err != nil {
			fmt.Println(err)
			return
		}

		playback, err := api.GetPlayback()

		if err != nil {
			fmt.Println(err)
			return
		}

		res := utils.Show(playback)

		fmt.Println(res)

	},
}

func init() {
	rootCmd.AddCommand(currentStatusCmd)
}
