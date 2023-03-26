package cmd

import (
	"fmt"
	"strings"

	"github.com/rollexxx23/spotify-cli/utils"
	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		track := strings.Join(args, " ")
		album, err := cmd.Flags().GetString("album")
		if err != nil {
			fmt.Println("An Error Occured")
		} else {
			api, err := utils.GetSpotifyAPI()
			if err != nil {
				fmt.Println(err)
			} else {
				if (track != "" && album != "") || (track == "" && album == "") {
					fmt.Println("Invalid Command")
				} else {
					if track != "" {
						fmt.Println("playing track...")
						str, err := utils.Play(api, track, 1)
						if err != nil {
							fmt.Println("error->", err)

						} else {
							fmt.Println(str)
						}
					} else {
						str, err := utils.Play(api, album, 2)
						if err != nil {
							fmt.Println(err)

						} else {
							fmt.Println(str)
						}
					}
				}
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(playCmd)

	playCmd.Flags().String("playlist", "", "playlist name")
	playCmd.Flags().String("album", "", "album name")
}
