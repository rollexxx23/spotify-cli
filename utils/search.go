package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/brianstrauch/spotify"
	"github.com/olekukonko/tablewriter"
)

func SearchTrack(api *spotify.API, query string, forPlaying bool) (*spotify.Track, error) {
	paging, err := api.Search(query, "track", 10)
	if err != nil {
		return nil, err
	}

	tracks := paging.Tracks.Items
	if len(tracks) == 0 {
		return nil, errors.New("Not Found")
	}
	if !forPlaying {
		output := new(strings.Builder)

		table := tablewriter.NewWriter(output)

		table.SetHeader([]string{"Track", "Artists"})

		for _, track := range tracks {

			artists := make([]string, len(track.Artists))
			for j, artist := range track.Artists {
				artists[j] = artist.Name
			}

			table.Append([]string{track.Name, strings.Join(artists, ", ")})
		}
		table.Render()
		fmt.Print(output)
		return nil, nil
	}
	return paging.Tracks.Items[0], nil
}

func SearchAlbum(api *spotify.API, query string) (*spotify.Album, error) {
	paging, err := api.Search(query, "album", 1)
	if err != nil {
		return nil, err
	}

	albums := paging.Albums.Items
	if len(albums) == 0 {
		return nil, errors.New("Not Found")
	}

	return albums[0], nil
}

func SearchPlaylist(api *spotify.API, query string) (*spotify.Playlist, error) {
	playlists, err := api.GetPlaylists()
	if err != nil {
		return nil, err
	}

	for _, playlist := range playlists {

		return playlist, nil

	}

	return nil, errors.New("Not Found")
}
