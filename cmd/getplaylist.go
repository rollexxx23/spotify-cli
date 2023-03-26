package cmd

import (
	"fmt"
	"strings"

	"github.com/brianstrauch/spotify"
	"github.com/rollexxx23/spotify-cli/utils"
	"github.com/spf13/cobra"
)

// getplaylistCmd represents the getplaylist command
var getplaylistCmd = &cobra.Command{
	Use:   "getplaylist",
	Short: "A brief description of your command",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.Join(args, " ")
		api, err := utils.GetSpotifyAPI()
		if err != nil {
			fmt.Println(err)
			return
		}

		err = getplaylist(api, query)

		if err != nil {
			fmt.Printf(err.Error())
		}

	},
}

func init() {
	rootCmd.AddCommand(getplaylistCmd)

}

func getplaylist(api *spotify.API, name string) error {
	playlist, err := utils.SearchPlaylist(api, name)
	if err != nil {
		return err
	}

	if err := playlist.HREF.Get(api, playlist); err != nil {
		return err
	}

	output, err := utils.DisplayPlaylist(playlist)
	if err != nil {
		return err
	}

	fmt.Print(output)
	return nil
}
