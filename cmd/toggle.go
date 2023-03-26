package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/brianstrauch/spotify"
	"github.com/rollexxx23/spotify-cli/utils"
	"github.com/spf13/cobra"
)

var pauseCmd = &cobra.Command{
	Use:   "toggle",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		api, err := utils.GetSpotifyAPI()

		if err != nil {
			fmt.Println(err)
			return
		}

		res, err := toggle(api)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)

	},
}

func init() {
	rootCmd.AddCommand(pauseCmd)
}

func toggle(api *spotify.API) (string, error) {
	playback, err := api.GetPlayback()
	if err != nil {
		return "", err
	}

	if playback == nil {
		return "", errors.New("No Active Device Found")
	}
	currState := playback.IsPlaying
	if currState {
		err := api.Pause()
		if err != nil {
			return "", err
		}
	} else {
		query := playback.Item.Track.Name
		_, err := utils.Play(api, query, 1)
		if err != nil {
			return "", err
		}
	}

	for playback.IsPlaying == currState {
		time.Sleep(10 * time.Millisecond)
		playback, err = api.GetPlayback()
		if err != nil {
			return "", err
		}
	}

	return utils.Show(playback), nil
}
