/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/brianstrauch/spotify"
	"github.com/rollexxx23/spotify-cli/utils"
	"github.com/spf13/cobra"
)

// nextCmd represents the next command
var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		api, err := utils.GetSpotifyAPI()

		if err != nil {
			fmt.Println(err)
			return
		}

		res, err := skipToNext(api)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(nextCmd)

}

func skipToNext(api *spotify.API) (string, error) {
	playback, err := api.GetPlayback()
	if err != nil {
		return "", err
	}

	if playback == nil {
		return "", errors.New("No Active Device Found")
	}

	id := playback.Item.ID
	progressMs := playback.ProgressMs

	if err := api.SkipToNextTrack(); err != nil {
		return "", err
	}

	playback, err = utils.WaitForUpdatedPlayback(api, func(playback *spotify.Playback) bool {
		return playback.Item.ID != id || playback.ProgressMs < progressMs
	})
	if err != nil {
		return "", err
	}

	return utils.Show(playback), nil
}
