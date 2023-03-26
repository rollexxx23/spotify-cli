package cmd

import (
	"fmt"

	"github.com/brianstrauch/spotify"
	"github.com/rollexxx23/spotify-cli/utils"
	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		api, err := utils.GetSpotifyAPI()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = save(api)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Saved!")

	},
}

func init() {
	rootCmd.AddCommand(saveCmd)

}

func save(api *spotify.API) error {
	playback, err := api.GetPlayback()
	if err != nil {
		return err
	}

	return api.SaveTracks(playback.Item.ID)
}
