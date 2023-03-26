package cmd

import (
	"fmt"
	"strings"

	"github.com/brianstrauch/spotify"
	"github.com/rollexxx23/spotify-cli/utils"
	"github.com/spf13/cobra"
)

// queueAddCmd represents the queueAdd command
var queueAddCmd = &cobra.Command{
	Use:   "queueAdd",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		api, err := utils.GetSpotifyAPI()
		if err != nil {
			fmt.Println(err)
			return
		}

		query := strings.Join(args, " ")

		err = addToQueue(api, query)
		if err != nil {
			fmt.Println(err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(queueAddCmd)

}

func addToQueue(api *spotify.API, query string) error {
	track, err := utils.SearchTrack(api, query, true)
	if err != nil {
		return err
	}

	if err := api.Queue(track.URI); err != nil {
		return err
	}
	fmt.Println(track.Name + " -> " + utils.JoinArtists(track.Artists) + ", Added To Queue")
	return nil
}
