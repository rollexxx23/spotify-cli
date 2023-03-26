/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/rollexxx23/spotify-cli/utils"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		api, err := utils.GetSpotifyAPI()
		track := strings.Join(args, " ")

		if err != nil {
			fmt.Println(err)
			return
		} else {
			utils.SearchTrack(api, track, false)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

}
